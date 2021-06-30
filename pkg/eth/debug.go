package eth

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/rpc"
	ipldEth "github.com/vulcanize/ipld-eth-server/pkg/eth"
	"github.com/vulcanize/tracing-api/pkg/cache"
	"github.com/vulcanize/tracing-api/pkg/eth/tracer"
)

const (
	// defaultTraceTimeout is the amount of time a single transaction can execute
	// by default before being forcefully aborted.
	defaultTraceTimeout = 5 * time.Second
)

type DebugAPI struct {
	// Local db backend
	backend *ipldEth.Backend
	cache   *cache.Service
}

func NewDebugAPI(b *ipldEth.Backend, cache *cache.Service) *DebugAPI {
	return &DebugAPI{b, cache}
}

func (api *DebugAPI) WriteTxTraceGraph(ctx context.Context, hash common.Hash) (*cache.TxTraceGraph, error) {
	data, err := api.TxTraceGraph(ctx, hash)
	if err != nil {
		return nil, err
	}
	return nil, api.cache.SaveTxTraceGraph(data)
}

func (api *DebugAPI) prepareEvm(ctx context.Context, hash common.Hash, tracer vm.Tracer) (*vm.EVM, *transaction, error) {
	tx, _, blockNum, txIndex, err := api.backend.GetTransaction(ctx, hash)
	if err != nil {
		return nil, nil, err
	}

	signer := types.MakeSigner(api.backend.Config.ChainConfig, big.NewInt(int64(blockNum)))

	block, err := api.backend.BlockByNumber(ctx, rpc.BlockNumber(blockNum))
	if err != nil {
		return nil, nil, err
	}

	statedb, _, err := api.backend.StateAndHeaderByNumber(ctx, rpc.BlockNumber(blockNum-1))
	if err != nil {
		return nil, nil, err
	}

	txs := block.Transactions()
	for i, ln := uint64(0), uint64(len(txs)); i < ln && i < txIndex; i++ {
		msg, err := txs[i].AsMessage(signer)
		if err != nil {
			return nil, nil, err
		}
		evm, _ := api.backend.GetEVM(ctx, msg, statedb, block.Header())
		_, err = core.ApplyMessage(evm, msg, new(core.GasPool).AddGas(math.MaxUint64))
		if err != nil {
			return nil, nil, err
		}
	}

	msg, err := tx.AsMessage(signer)
	if err != nil {
		return nil, nil, err
	}

	statedb.SetBalance(msg.From(), math.MaxBig256)
	vmctx := core.NewEVMBlockContext(block.Header(), api.backend, nil)
	txContext := core.NewEVMTxContext(msg)

	cfg := api.backend.Config.VmConfig
	cfg.Debug = true
	cfg.Tracer = tracer

	evm := vm.NewEVM(vmctx, txContext, statedb, api.backend.Config.ChainConfig, cfg)
	internalTx := transaction{
		Message:   msg,
		Index:     txIndex,
		BlockHash: block.Hash(),
		BlockNumb: blockNum,
	}
	return evm, &internalTx, nil
}

func (api *DebugAPI) TxTraceGraph(ctx context.Context, hash common.Hash) (*cache.TxTraceGraph, error) {
	degugger := tracer.NewCallTracer()

	evm, tx, err := api.prepareEvm(ctx, hash, degugger)
	if err != nil {
		return nil, err
	}

	_, err = core.ApplyMessage(evm, tx.Message, new(core.GasPool).AddGas(math.MaxUint64))
	if err != nil {
		return nil, err
	}

	frames := []tracer.Frame{
		{
			Op:     vm.CALL,
			From:   tx.Message.From(),
			To:     *tx.Message.To(),
			Input:  tx.Message.Data(),
			Output: degugger.Output(),
			Gas:    tx.Message.Gas(),
			Cost:   tx.Message.Gas(),
			Value:  tx.Message.Value(),
		},
	}

	return &cache.TxTraceGraph{
		TxHash:      hash,
		TxIndex:     tx.Index,
		BlockHash:   tx.BlockHash,
		BlockNumber: tx.BlockNumb,
		Frames:      append(frames, degugger.Frames()...),
	}, nil
}

func (api *DebugAPI) TraceTransaction(ctx context.Context, hash common.Hash, config *eth.TraceConfig) (interface{}, error) {
	var (
		tracer vm.Tracer
		err    error
	)
	switch {
	case config != nil && config.Tracer != nil:
		// Define a meaningful timeout of a single transaction trace
		timeout := defaultTraceTimeout
		if config.Timeout != nil {
			if timeout, err = time.ParseDuration(*config.Timeout); err != nil {
				return nil, err
			}
		}
		// Constuct the JavaScript tracer to execute with
		if tracer, err = tracers.New(*config.Tracer); err != nil {
			return nil, err
		}
		// Handle timeouts and RPC cancellations
		deadlineCtx, cancel := context.WithTimeout(ctx, timeout)
		go func() {
			<-deadlineCtx.Done()
			tracer.(*tracers.Tracer).Stop(errors.New("execution timeout"))
		}()
		defer cancel()

	case config == nil:
		tracer = vm.NewStructLogger(nil)

	default:
		tracer = vm.NewStructLogger(config.LogConfig)
	}

	evm, tx, err := api.prepareEvm(ctx, hash, tracer)
	if err != nil {
		return nil, err
	}

	result, err := core.ApplyMessage(evm, tx.Message, new(core.GasPool).AddGas(math.MaxUint64))
	if err != nil {
		return nil, fmt.Errorf("tracing failed: %v", err)
	}

	// Depending on the tracer type, format and return the output
	switch tracer := tracer.(type) {
	case *vm.StructLogger:
		// If the result contains a revert reason, return it.
		returnVal := fmt.Sprintf("%x", result.Return())
		if len(result.Revert()) > 0 {
			returnVal = fmt.Sprintf("%x", result.Revert())
		}
		return &ExecutionResult{
			Gas:         result.UsedGas,
			Failed:      result.Failed(),
			ReturnValue: returnVal,
			StructLogs:  FormatLogs(tracer.StructLogs()),
		}, nil

	case *tracers.Tracer:
		return tracer.GetResult()

	default:
		panic(fmt.Sprintf("bad tracer type %T", tracer))
	}
}

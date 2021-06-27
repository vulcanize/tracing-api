package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/vulcanize/ipld-eth-server/pkg/eth"
	"github.com/vulcanize/tracing-api/pkg/cache"
	"github.com/vulcanize/tracing-api/pkg/eth/tracer"
)

type DebugAPI struct {
	// Local db backend
	backend *eth.Backend
	cache   *cache.Service
}

func NewDebugAPI(b *eth.Backend, cache *cache.Service) *DebugAPI {
	return &DebugAPI{b, cache}
}

func (api *DebugAPI) WriteTxTraceGraph(ctx context.Context, hash common.Hash) (*cache.TxTraceGraph, error) {
	data, err := api.TxTraceGraph(ctx, hash)
	if err != nil {
		return nil, err
	}
	return nil, api.cache.SaveTxTraceGraph(data)
}

func (api *DebugAPI) prepareEvm(ctx context.Context, hash common.Hash, ttracer vm.Tracer) (*vm.EVM, *transaction, error) {
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
	cfg.Tracer = ttracer

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

func (api *DebugAPI) TraceTransaction(ctx context.Context, hash common.Hash) (*ExecutionResult, error) {
	degugger := vm.NewStructLogger(&vm.LogConfig{
		DisableMemory:     false,
		DisableStack:      false,
		DisableStorage:    false,
		DisableReturnData: false,
		Debug:             true,
		Limit:             0,
	})

	evm, tx, err := api.prepareEvm(ctx, hash, degugger)
	if err != nil {
		return nil, err
	}

	result, err := core.ApplyMessage(evm, tx.Message, new(core.GasPool).AddGas(math.MaxUint64))
	if err != nil {
		return nil, err
	}

	returnVal := fmt.Sprintf("%x", result.Return())
	if len(result.Revert()) > 0 {
		returnVal = fmt.Sprintf("%x", result.Revert())
	}

	return &ExecutionResult{
		Gas:         result.UsedGas,
		Failed:      result.Failed(),
		ReturnValue: returnVal,
		StructLogs:  FormatLogs(degugger.StructLogs()),
	}, nil
}

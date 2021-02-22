package eth

import (
	"context"
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

func (api *DebugAPI) TxTraceGraph(ctx context.Context, hash common.Hash) (*cache.TxTraceGraph, error) {
	tx, _, blockNum, txIndex, err := api.backend.GetTransaction(ctx, hash)
	if err != nil {
		return nil, err
	}

	signer := types.MakeSigner(api.backend.Config.ChainConfig, big.NewInt(int64(blockNum)))

	block, err := api.backend.BlockByNumber(ctx, rpc.BlockNumber(blockNum))
	if err != nil {
		return nil, err
	}

	statedb, _, err := api.backend.StateAndHeaderByNumber(ctx, rpc.BlockNumber(blockNum-1))
	if err != nil {
		return nil, err
	}

	txs := block.Transactions()
	for i, ln := uint64(0), uint64(len(txs)); i < ln && i < txIndex; i++ {
		msg, err := tx.AsMessage(signer)
		if err != nil {
			return nil, err
		}
		evm, _ := api.backend.GetEVM(ctx, msg, statedb, block.Header())
		_, err = core.ApplyMessage(evm, msg, new(core.GasPool).AddGas(math.MaxUint64))
		if err != nil {
			return nil, err
		}
	}

	msg, err := tx.AsMessage(signer)
	statedb.SetBalance(msg.From(), math.MaxBig256)
	vmctx := core.NewEVMBlockContext(block.Header(), api.backend, nil)
	txContext := core.NewEVMTxContext(msg)

	tracer := tracer.NewCallTracer()
	cfg := api.backend.Config.VmConfig
	cfg.Debug = true
	cfg.Tracer = tracer

	evm := vm.NewEVM(vmctx, txContext, statedb, api.backend.Config.ChainConfig, cfg)
	_, err = core.ApplyMessage(evm, msg, new(core.GasPool).AddGas(math.MaxUint64))
	if err != nil {
		return nil, err
	}

	return &cache.TxTraceGraph{
		TxHash:      hash,
		TxIndex:     txIndex,
		BlockHash:   block.Hash(),
		BlockNumber: blockNum,
		Frames:      tracer.Frames(),
	}, nil
}

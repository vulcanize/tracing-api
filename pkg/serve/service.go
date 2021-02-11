package serve

import (
	"github.com/ethereum/go-ethereum/rpc"
	srv "github.com/vulcanize/ipld-eth-server/pkg/serve"
	"github.com/vulcanize/tracing-api/pkg/cache"
	"github.com/vulcanize/tracing-api/pkg/eth"
)

type DebugService struct {
	srv.Server
	cache *cache.Service
}

// NewServer creates a new Server using an underlying Service struct
func NewServer(settings *srv.Config, cache *cache.Service) (srv.Server, error) {
	srv, err := srv.NewServer(settings)
	if err != nil {
		return nil, err
	}
	return &DebugService{srv, cache}, nil
}

// APIs returns the RPC descriptors the watcher service offers
func (sap *DebugService) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "debug",
			Version:   "0.0.1",
			Service:   eth.NewDebugAPI(sap.Backend(), sap.cache),
			Public:    true,
		},
	}
}

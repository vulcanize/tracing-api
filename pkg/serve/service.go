package serve

import (
	"github.com/ethereum/go-ethereum/rpc"
	srv "github.com/vulcanize/ipld-eth-server/pkg/serve"
	"github.com/vulcanize/tracing-api/pkg/eth"
)

type DebugService struct {
	srv.Server
}

// NewServer creates a new Server using an underlying Service struct
func NewServer(settings *srv.Config) (srv.Server, error) {
	srv, err := srv.NewServer(settings)
	if err != nil {
		return nil, err
	}
	return &DebugService{srv}, nil
}

// APIs returns the RPC descriptors the watcher service offers
func (sap *DebugService) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "debug",
			Version:   "0.0.1",
			Service:   eth.NewDebugAPI(sap.Backend()),
			Public:    true,
		},
	}
}

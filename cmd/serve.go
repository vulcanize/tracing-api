package cmd

import (
	"os"
	"os/signal"
	"sync"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vulcanize/ipld-eth-indexer/pkg/eth"
	srpc "github.com/vulcanize/ipld-eth-server/pkg/rpc"
	srv "github.com/vulcanize/ipld-eth-server/pkg/serve"
	"github.com/vulcanize/tracing-api/pkg/cache"
	"github.com/vulcanize/tracing-api/pkg/serve"
)

var (
	serveCmd = &cobra.Command{
		Use: "serve",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.WithField("command", cmd.CalledAs()).Infof("running tracing-api version: %s", version)

			serverConfig, err := srv.NewConfig()
			if err != nil {
				return err
			}

			cache, err := cache.New()
			if err != nil {
				return err
			}

			server, err := serve.NewServer(serverConfig, cache)
			if err != nil {
				return err
			}

			wg := new(sync.WaitGroup)
			forwardPayloadChan := make(chan eth.ConvertedPayload, srv.PayloadChanBufferSize)
			server.Serve(wg, forwardPayloadChan)
			if err := startServers(server, serverConfig); err != nil {
				return err
			}

			shutdown := make(chan os.Signal)
			signal.Notify(shutdown, os.Interrupt)
			<-shutdown

			server.Stop()
			wg.Wait()

			return nil
		},
	}
)

func startServers(server srv.Server, settings *srv.Config) error {
	logrus.Info("starting up IPC server")
	_, _, err := srpc.StartIPCEndpoint(settings.IPCEndpoint, server.APIs())
	if err != nil {
		return err
	}
	logrus.Info("starting up WS server")
	_, _, err = srpc.StartWSEndpoint(settings.WSEndpoint, server.APIs(), []string{"debug"}, nil, true)
	if err != nil {
		return err
	}
	logrus.Info("starting up HTTP server")
	_, err = srpc.StartHTTPEndpoint(settings.HTTPEndpoint, server.APIs(), []string{"debug"}, nil, []string{"*"}, rpc.HTTPTimeouts{})
	return err
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// flags
	serveCmd.PersistentFlags().String("http-host", "127.0.0.1", "http host")
	serveCmd.PersistentFlags().String("http-port", "8080", "http port")
	serveCmd.PersistentFlags().String("http-path", "/", "http base path")

	serveCmd.PersistentFlags().String("eth-http-path", "", "http url for ethereum node")
	serveCmd.PersistentFlags().String("eth-node-id", "", "eth node id")
	serveCmd.PersistentFlags().String("eth-client-name", "Geth", "eth client name")
	serveCmd.PersistentFlags().String("eth-genesis-block", "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3", "eth genesis block hash")
	serveCmd.PersistentFlags().String("eth-network-id", "1", "eth network id")
	serveCmd.PersistentFlags().String("eth-chain-id", "1", "eth chain id")
	serveCmd.PersistentFlags().String("eth-default-sender", "", "default sender address")
	serveCmd.PersistentFlags().String("eth-rpc-gas-cap", "", "rpc gas cap (for eth_Call execution)")

	serveCmd.PersistentFlags().String("cache-database-name", "vulcanize_public", "database name")
	serveCmd.PersistentFlags().String("cache-database-hostname", "localhost", "database hostname")
	serveCmd.PersistentFlags().Int("cache-database-port", 5432, "database port")
	serveCmd.PersistentFlags().String("cache-database-user", "postgres", "database user")
	serveCmd.PersistentFlags().String("cache-database-password", "", "database password")
	serveCmd.PersistentFlags().Int("cache-database-maxIdle", 0, "database password")
	serveCmd.PersistentFlags().Int("cache-database-maxOpen", 0, "database password")
	serveCmd.PersistentFlags().Int("cache-database-maxLifetime", 0, "database password")

	// and their .toml config bindings
	viper.BindPFlag("http.host", serveCmd.PersistentFlags().Lookup("http-host"))
	viper.BindPFlag("http.port", serveCmd.PersistentFlags().Lookup("http-port"))
	viper.BindPFlag("http.path", serveCmd.PersistentFlags().Lookup("http-path"))

	viper.BindPFlag("eth.rpc", serveCmd.PersistentFlags().Lookup("eth-rpc"))

	viper.BindPFlag("ethereum.httpPath", serveCmd.PersistentFlags().Lookup("eth-http-path"))
	viper.BindPFlag("ethereum.nodeID", serveCmd.PersistentFlags().Lookup("eth-node-id"))
	viper.BindPFlag("ethereum.clientName", serveCmd.PersistentFlags().Lookup("eth-client-name"))
	viper.BindPFlag("ethereum.genesisBlock", serveCmd.PersistentFlags().Lookup("eth-genesis-block"))
	viper.BindPFlag("ethereum.networkID", serveCmd.PersistentFlags().Lookup("eth-network-id"))
	viper.BindPFlag("ethereum.chainID", serveCmd.PersistentFlags().Lookup("eth-chain-id"))
	viper.BindPFlag("ethereum.defaultSender", serveCmd.PersistentFlags().Lookup("eth-default-sender"))
	viper.BindPFlag("ethereum.rpcGasCap", serveCmd.PersistentFlags().Lookup("eth-rpc-gas-cap"))

	viper.BindPFlag("cache.database.name", serveCmd.PersistentFlags().Lookup("cache-database-name"))
	viper.BindPFlag("cache.database.hostname", serveCmd.PersistentFlags().Lookup("cache-database-hostname"))
	viper.BindPFlag("cache.database.port", serveCmd.PersistentFlags().Lookup("cache-database-port"))
	viper.BindPFlag("cache.database.user", serveCmd.PersistentFlags().Lookup("cache-database-user"))
	viper.BindPFlag("cache.database.password", serveCmd.PersistentFlags().Lookup("cache-database-password"))
	viper.BindPFlag("cache.database.maxIdle", serveCmd.PersistentFlags().Lookup("cache-database-maxIdle"))
	viper.BindPFlag("cache.database.maxOpen", serveCmd.PersistentFlags().Lookup("cache-database-maxOpen"))
	viper.BindPFlag("cache.database.maxLifetime", serveCmd.PersistentFlags().Lookup("cache-database-maxLifetime"))
}

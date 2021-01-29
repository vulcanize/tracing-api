package cmd

import (
	"os"
	"os/signal"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vulcanize/ipld-eth-indexer/pkg/eth"
	"github.com/vulcanize/ipld-eth-server/pkg/serve"
)

var (
	serveCmd = &cobra.Command{
		Use: "serve",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.WithField("command", cmd.CalledAs()).Infof("running tracing-api version: %s", version)

			serverConfig, err := serve.NewConfig()
			if err != nil {
				return err
			}
			spew.Dump(serverConfig)

			server, err := serve.NewServer(serverConfig)
			if err != nil {
				return err
			}

			wg := new(sync.WaitGroup)
			forwardPayloadChan := make(chan eth.ConvertedPayload, serve.PayloadChanBufferSize)
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

func startServers(server serve.Server, settings *serve.Config) error {
	return nil
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
}

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version = ""

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of tracing-api",
	Long: `Use this command to fetch the version of tracing-api

Usage: ./tracing-api version`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("tracing-api version: %s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

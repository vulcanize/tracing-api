package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config  string
	rootCmd = &cobra.Command{
		Use: "tracer",
	}
)

func init() {
	cobra.OnInitialize(func() {
		loglevel, err := logrus.ParseLevel(viper.GetString("log.level"))
		if err == nil {
			logrus.SetLevel(loglevel)
		}
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: viper.GetBool("log.timestamp"),
		})
		if config == "" {
			logrus.Warn("no config file passed with --config flag")
			return
		}
		viper.SetConfigFile(config)
		if err := viper.ReadInConfig(); err == nil {
			logrus.WithField("config", viper.ConfigFileUsed()).Info("using config file")
		} else {
			logrus.WithError(err).Fatal("couldn't read config file")
		}

	})

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file location")

	// flags
	rootCmd.PersistentFlags().String("log-level", logrus.InfoLevel.String(), "log level (trace, debug, info, warn, error, fatal, panic)")
	rootCmd.PersistentFlags().String("log-file", "", "file path for logging")
	rootCmd.PersistentFlags().Bool("log-timestamp", true, "show full timestamp in logger")

	rootCmd.PersistentFlags().Bool("metrics", false, "enable prometheus")
	rootCmd.PersistentFlags().String("metrics-host", "127.0.0.1", "prometheus http host")
	rootCmd.PersistentFlags().String("metrics-port", "8080", "prometheus http port")

	rootCmd.PersistentFlags().String("database-name", "vulcanize_public", "database name")
	rootCmd.PersistentFlags().Int("database-port", 5432, "database port")
	rootCmd.PersistentFlags().String("database-hostname", "localhost", "database hostname")
	rootCmd.PersistentFlags().String("database-user", "postgres", "database user")
	rootCmd.PersistentFlags().String("database-password", "", "database password")

	// and their .toml config bindings
	viper.BindPFlag("log.file", rootCmd.PersistentFlags().Lookup("log-file"))
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("log.timestamp", rootCmd.PersistentFlags().Lookup("log-timestamp"))

	viper.BindPFlag("metrics", rootCmd.PersistentFlags().Lookup("metrics"))
	viper.BindPFlag("metrics.host", rootCmd.PersistentFlags().Lookup("metrics-host"))
	viper.BindPFlag("metrics.port", rootCmd.PersistentFlags().Lookup("metrics-port"))

	viper.BindPFlag("database.name", rootCmd.PersistentFlags().Lookup("database-name"))
	viper.BindPFlag("database.port", rootCmd.PersistentFlags().Lookup("database-port"))
	viper.BindPFlag("database.hostname", rootCmd.PersistentFlags().Lookup("database-hostname"))
	viper.BindPFlag("database.user", rootCmd.PersistentFlags().Lookup("database-user"))
	viper.BindPFlag("database.password", rootCmd.PersistentFlags().Lookup("database-password"))
}

// Execute main function
func Execute() error {
	logrus.Info("----- Starting tracing api -----")
	return rootCmd.Execute()
}

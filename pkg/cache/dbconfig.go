package cache

import (
	"github.com/spf13/viper"
	"github.com/vulcanize/ipld-eth-indexer/pkg/postgres"
)

// Env variables
const (
	DATABASE_NAME                 = "CACHE_DATABASE_NAME"
	DATABASE_HOSTNAME             = "CACHE_DATABASE_HOSTNAME"
	DATABASE_PORT                 = "CACHE_DATABASE_PORT"
	DATABASE_USER                 = "CACHE_DATABASE_USER"
	DATABASE_PASSWORD             = "CACHE_DATABASE_PASSWORD"
	DATABASE_MAX_IDLE_CONNECTIONS = "CACHE_DATABASE_MAX_IDLE_CONNECTIONS"
	DATABASE_MAX_OPEN_CONNECTIONS = "CACHE_DATABASE_MAX_OPEN_CONNECTIONS"
	DATABASE_MAX_CONN_LIFETIME    = "CACHE_DATABASE_MAX_CONN_LIFETIME"
)

func dbConfig() postgres.Config {
	viper.BindEnv("cache.database.name", DATABASE_NAME)
	viper.BindEnv("cache.database.hostname", DATABASE_HOSTNAME)
	viper.BindEnv("cache.database.port", DATABASE_PORT)
	viper.BindEnv("cache.database.user", DATABASE_USER)
	viper.BindEnv("cache.database.password", DATABASE_PASSWORD)
	viper.BindEnv("cache.database.maxIdle", DATABASE_MAX_IDLE_CONNECTIONS)
	viper.BindEnv("cache.database.maxOpen", DATABASE_MAX_OPEN_CONNECTIONS)
	viper.BindEnv("cache.database.maxLifetime", DATABASE_MAX_CONN_LIFETIME)
	return postgres.Config{
		Name:        viper.GetString("cache.database.name"),
		Hostname:    viper.GetString("cache.database.hostname"),
		Port:        viper.GetInt("cache.database.port"),
		User:        viper.GetString("cache.database.user"),
		Password:    viper.GetString("cache.database.password"),
		MaxIdle:     viper.GetInt("cache.database.maxIdle"),
		MaxOpen:     viper.GetInt("cache.database.maxOpen"),
		MaxLifetime: viper.GetInt("cache.database.maxLifetime"),
	}
}

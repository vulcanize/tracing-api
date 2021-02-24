module github.com/vulcanize/tracing-api

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.9.25
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.8.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/vulcanize/ipfs-ethdb v0.0.2-alpha
	github.com/vulcanize/ipld-eth-indexer v0.7.0-alpha
	github.com/vulcanize/ipld-eth-server v0.3.0-alpha
)

replace github.com/ethereum/go-ethereum v1.9.25 => github.com/vulcanize/go-ethereum v1.9.25-statediff-0.0.14

replace github.com/vulcanize/ipfs-ethdb v0.0.2-alpha => github.com/vulcanize/pg-ipfs-ethdb v0.0.2-alpha

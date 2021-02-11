module github.com/vulcanize/tracing-api

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.9.11
	github.com/friendsofgo/graphiql v0.2.2
	github.com/graphql-go/graphql v0.7.9
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.5.2
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	github.com/valyala/fastjson v1.6.3
	github.com/vulcanize/ipld-eth-indexer v0.6.0-alpha
	github.com/vulcanize/ipld-eth-server v0.2.0-alpha
)

replace github.com/ethereum/go-ethereum v1.9.11 => github.com/vulcanize/go-ethereum v1.9.11-statediff-0.0.8

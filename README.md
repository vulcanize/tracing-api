# tracing-api

tracing-api serves JSON-RPC endpoints which allows to get internal transactions by transaction hash.
It uses indexed ETH IPLD objects from [geth-statediff](https://github.com/vulcanize/go-ethereum/releases)

## JSON-RPC endpoints

The currently supported standard endpoints are:
* `debug_txTraceGraph` - returns traces in graph format
* `debug_writeTxTraceGraph` - writes traces to cache database

## Environment Variables

| Name                      | Default Value    | Comment                          |
|---------------------------|------------------|----------------------------------|
| DATABASE_NAME             | vulcanize_public | Source database name             |
| DATABASE_PORT             | 5432             | Source database port             |
| DATABASE_HOSTNAME         | localhost        | Source database host             |
| DATABASE_USER             | postgres         | Source database user             |
| DATABASE_PASSWORD         |                  | Source database password         |
| CACHE_DATABASE_NAME       | vulcanize_public | Cache database name              |
| CACHE_DATABASE_PORT       | 5432             | Cache database port              |
| CACHE_DATABASE_HOSTNAME   | localhost        | Cache database host              |
| CACHE_DATABASE_USER       | postgres         | Cache database user              |
| CACHE_DATABASE_PASSWORD   |                  | Cache database password          |
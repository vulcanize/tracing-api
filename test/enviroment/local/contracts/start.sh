#!/bin/bash
sleep 120
cd /usr/src/app
./node_modules/.bin/truffle migrate --network development
./node_modules/.bin/graph codegen
./node_modules/.bin/graph build
./node_modules/.bin/graph create --node ${CNT_GRAPH_URL} vulcanize/bnumstore
./node_modules/.bin/graph deploy --node ${CNT_GRAPH_URL} --ipfs ${CNT_IPFS_URL} vulcanize/bnumstore
node ./caller
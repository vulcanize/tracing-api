Smart Contract being deployed to [mainnet](https://etherscan.io/address/0x51900544dfef84a65f6dbe0c999f88ebe87ba2dd#code)

Update addresses in `subgraph.yaml` to:
* `0x51900544dfef84a65f6dbe0c999f88ebe87ba2dd` for BlockNumStorage
* `0xddfe1a03bab27687f49f1864c98ceaf8a26caf2f` for UintStorage 

Add `startBlock: 12056420` to `dataSources[*].source`

### Deploy mainnet subgraph

* `npm i`
* `npm run codegen`
* `npm run build`
* `npm run deploy-mainnet`
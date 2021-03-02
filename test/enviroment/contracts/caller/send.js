const fs = require('fs');
const yaml = require('yaml');
const Web3 = require('web3');
const JsonRPC = require('./jsonrpc');
const IBlockNumStorage = require('../build/contracts/BlockNumStorage.json');

const web3 = new Web3('http://localhost:8545');
const rpc = JsonRPC('http://localhost:8083');

main().catch(console.log);
async function main() {
  const [from] = await web3.eth.getAccounts();
  let txs = [];
  for (let network in IBlockNumStorage.networks) {
    const { address } = IBlockNumStorage.networks[network];
    const bstore = new web3.eth.Contract(IBlockNumStorage.abi, address);
    txs = await Promise.all([
      bstore.methods.sync("counter").send({ from, value: '0x5551' }),
      bstore.methods.sync("counter").send({ from, value: '0x5552' }),
      bstore.methods.sync("counter").send({ from, value: '0x5553' })
    ]);
  }
  await sleep(2000);
  for (const { transactionHash } of txs) {
    console.log(`Debug: ${transactionHash}`);
    try {
      console.log(await rpc.exec("debug_writeTxTraceGraph", transactionHash));
    } catch (e) {
      console.log(`Error: ${e}`);
    }
  }
}

async function sleep(timeout) {
  return new Promise(r => setTimeout(r, timeout));
}


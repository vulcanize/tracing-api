const fastify = require('fastify')({ logger: true })
const Web3 = require('web3');
const IBlockNumStorage = require('../build/contracts/BlockNumStorage.json');
const web3 = new Web3(`http://${process.env.CNT_ETH_HOST}:${process.env.CNT_ETH_PORT}`);

fastify.get('/', async (req, reply) => {
  const [from] = await web3.eth.getAccounts();
  let txs = [];
  for (let network in IBlockNumStorage.networks) {
    const { address } = IBlockNumStorage.networks[network];
    const bstore = new web3.eth.Contract(IBlockNumStorage.abi, address);
    txs = await Promise.all([
      bstore.methods.sync("counter").send({ from, value: '0x5551' })
    ]);
  }
  return txs.map(({ transactionHash }) => transactionHash);
})

async function main() {
  try {
    await fastify.listen(3000, '0.0.0.0');
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
}

main();



console.log(`${process.env.CNT_ETH_HOST}:${process.env.CNT_ETH_PORT}`);

module.exports = {
  networks: {
    development: {
      host: process.env.CNT_ETH_HOST,
      port: process.env.CNT_ETH_PORT,
      network_id: "*",
      gas: 8000000,
    },
  },

  mocha: {
    // timeout: 100000
  },

  // Configure your compilers
  compilers: {
    solc: {
      settings: {
        optimizer: {
          enabled: true,
          runs: 1
        }
      },
      //version: "node_modules/solc"
    }
  }
};

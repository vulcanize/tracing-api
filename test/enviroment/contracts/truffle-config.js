
module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
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

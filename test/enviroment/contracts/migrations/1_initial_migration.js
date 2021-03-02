const fs = require('fs');
const yaml = require('yaml');
const Migrations = artifacts.require("Migrations");
const BlockNumStorage = artifacts.require("BlockNumStorage");

async function main(deployer) {
  await deployer.deploy(Migrations);
  const store = await deployer.deploy(BlockNumStorage);

  let subgraph = yaml.parse(fs.readFileSync("subgraph.yaml", {encoding:"utf-8"}));
  subgraph.dataSources[0].source = store.address;
  fs.writeFileSync("subgraph.yaml", yaml.stringify(subgraph))
}

module.exports = function (deployer) {
  deployer.then(() => main(deployer));
};

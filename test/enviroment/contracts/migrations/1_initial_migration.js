const Migrations = artifacts.require("Migrations");
const BlockNumStorage = artifacts.require("BlockNumStorage");

async function main(deployer) {
  await Promise.all([
    deployer.deploy(Migrations),
    deployer.deploy(BlockNumStorage),
  ]);
}

module.exports = function (deployer) {
  deployer.then(() => main(deployer));
};

//固定总量代币
const ZombieCore = artifacts.require("ZombieCore");

module.exports = function (deployer) {
  deployer.deploy(ZombieCore);
};

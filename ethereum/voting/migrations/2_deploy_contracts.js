const Voting = artifacts.require("Voting");

module.exports = function (deployer) {
  deployer.deploy(
    Voting,
    10000,
    web3.utils.toWei("0.01", "ether"),
    ["Alice", "Bob", "Cary"].map((x) => web3.utils.asciiToHex(x))
  );
};

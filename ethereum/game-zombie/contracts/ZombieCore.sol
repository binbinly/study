// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ZombieMarket.sol";
import "./ZombieFeeding.sol";
import "./ZombieAttack.sol";

//入口
contract ZombieCore is ZombieMarket, ZombieFeeding, ZombieAttack {
    string public constant name = "MyCryptoZombie";
    string public constant symbol = "MCZ";

    receive() external payable {}

    function withdraw() external onlyOwner {
        payable(owner()).transfer(address(this).balance);
    }

    function balance() external view onlyOwner returns (uint256) {
        return address(this).balance;
    }
}

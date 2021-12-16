// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

//可销毁代币
contract ERC20WithBurnable is ERC20Burnable {
    constructor(
        string memory name, //代币名称
        string memory symbol, //代币缩写
        uint256 totalSupply //发行总量
    ) ERC20(name, symbol) {
        _mint(msg.sender, totalSupply * 10**18);
    }
}

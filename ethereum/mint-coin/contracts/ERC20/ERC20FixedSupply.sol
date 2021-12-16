// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

//固定总量代币
contract ERC20FixedSupply is ERC20 {
    constructor(
        string memory name, //代币名称
        string memory symbol, //代币缩写
        uint256 totalSupply //发行总量
    ) ERC20(name, symbol) {
        _mint(msg.sender, totalSupply * 10**18);
    }
}

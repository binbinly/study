// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Ownable.sol";

//工厂
contract Factory is Ownable {
    //治理币
    address public token;
    //接受币地址
    address public receiveAddress;
    //最小交易额
    uint256 public minPrice = 1 ether;

    constructor() {
        receiveAddress = msg.sender;
    }

    //设置接收币地址
    function setReceiveAddress(address addr) public onlyOwner {
        receiveAddress = addr;
    }

    //设置代币地址
    function setToken(address _token) public onlyOwner {
        token = _token;
    }

    function setMinPrice(uint256 _price) public onlyOwner {
        minPrice = _price;
    }
}

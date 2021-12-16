// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Crowdsale.sol";
import "./Allowance.sol";

//普通的众筹
contract AllowanceCrowdsaleContract is Allowance {
    constructor(
        uint256 rate, // 兑换比例
        address payable wallet, // 接收ETH受益人地址
        IERC20 token, // 代币地址
        address tokenWallet // 代币从这个地址发送
    ) Allowance(tokenWallet) Crowdsale(rate, wallet, token) {}
}

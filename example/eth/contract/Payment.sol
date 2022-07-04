// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Factory.sol";
import "./interface/IERC20.sol";

//支付合约
contract Payment is Factory {
    struct TradeLog {
        address addr;
        uint256 amount;
        uint256 date;
    }

    //交易记录
    mapping(string => TradeLog) public trades;

    //支付成功回调事件
    event OnFinish(address addr, string order_no);

    //支付
    function pay(string memory order_no, uint256 value) public {
        require(value >= minPrice, "value min limit");
        require(trades[order_no].amount == 0, "repeat pay");
        IERC20(token).transferFrom(msg.sender, receiveAddress, value);
        trades[order_no] = TradeLog(msg.sender, value, block.timestamp);
        emit OnFinish(msg.sender, order_no);
    }

    //检查是否已支付
    function query(string memory order_no, uint256 value)
        public
        view
        returns (bool)
    {
        return trades[order_no].amount == value;
    }
}

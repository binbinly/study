// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Factory.sol";
import "./interface/IERC20.sol";

//兑换合约
contract Exchange is Factory {
    struct TradeLog {
        address addr;
        uint256 amount;
        uint256 date;
    }

    //交易记录
    mapping(string => TradeLog) public trades;

    //兑换成功回调事件
    event OnFinish(address addr, string order_no);

    //提币
    function withdraw(
        string memory order_no,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public {
        require(value >= minPrice, "value min limit");
        require(trades[order_no].amount == 0, "repeat withdraw");
        IERC20(token).permit(
            receiveAddress,
            address(this),
            value,
            deadline,
            v,
            r,
            s
        );
        IERC20(token).transferFrom(receiveAddress, msg.sender, value);
        trades[order_no] = TradeLog(msg.sender, value, block.timestamp);
        emit OnFinish(msg.sender, order_no);
    }

    //检查是否已兑换
    function query(string memory order_no, uint256 value)
        public
        view
        returns (bool)
    {
        return trades[order_no].amount == value;
    }
}

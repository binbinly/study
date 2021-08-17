// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;

contract Lottery{
    //彩票期数
    uint public period = 1;
    //管理员
    address public owner;

    //投注者
    struct Account {
        address addr;
        string nums;
    }
    Account[] public accountList;

    //开奖结果
    uint[] public retNumsList;

    mapping(uint => Account[]) public accounts;
    mapping(uint=>uint[]) public retNums;

    constructor() {
        owner = msg.sender;
    }

    modifier checkOwner(){
        require(msg.sender == owner);
        _;
    }

    modifier checkMoney(){
        require(msg.value == 1 ether);
        _;
    }

    event BetFinish(address addr, string nums);
    event DrawFinish(uint[] nums);

    //投注
    function bet(string memory _nums) public payable checkMoney {
        accountList.push(Account({
        addr: msg.sender,
        nums: _nums
        }));
        emit BetFinish(msg.sender, _nums);
    }

    function getBalance(address addr) public view returns(uint) {
        return addr.balance;
    }

    //开奖
    function draw() public checkOwner {
        // 红球从1-4里面随机选一个，选4次
        for(uint i=1;i<=4;i++) {
            retNumsList.push(random(i) % 4 + 1);
        }

        // 篮球从1-3里面随机选一个，选一次
        retNumsList.push(random(10) % 3 + 1);

        // 存储投注人和投注号码历史记录
        accounts[period] = accountList;
        // 存储彩票期数，开奖结果历史记录
        retNums[period] = retNumsList;

        //清空投注记录
        delete accountList;
        delete retNumsList;
        period += 1;
        emit DrawFinish(retNumsList);
    }

    //产生随机数
    function random(uint seed) private view returns(uint){
        return uint(keccak256(abi.encodePacked(seed, block.timestamp)));
    }
}
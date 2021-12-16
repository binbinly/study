// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Ownable.sol";

//僵尸工厂
contract ZombieFactory is Ownable {
    //生成新僵尸事件
    event NewZombie(uint256 zombieId, string name, uint256 dna);

    //基因位数
    uint256 dnaDigits = 16;

    uint256 dnaModulus = 10**dnaDigits;

    //冷却时间
    uint256 public cooldownTime = 1 days;

    //价格
    uint256 public price = 0.01 ether;

    //僵尸总数
    uint256 public total = 0;

    //僵尸结构
    struct Zombie {
        string name; //名称
        uint256 dna;
        uint16 winCount; //胜利次数
        uint16 loseCount; //失败次数
        uint32 level; //等级
        uint32 readyTime; //冷却时间
    }

    //所有僵尸
    Zombie[] public zombies;

    // 僵尸id => 所属者
    mapping(uint256 => address) public zombieToOwner;
    // 所有者 => 僵尸数量
    mapping(address => uint256) public ownerZombieCount;
    // 僵尸id => 喂食次数
    mapping(uint256 => uint256) public zombieFeedCount;

    //创建僵尸
    function _createZombie(string memory _name, uint256 _dna) internal {
        zombies.push(
            Zombie({
                name: _name,
                dna: _dna,
                winCount: 0,
                loseCount: 0,
                level: 1,
                readyTime: 0
            })
        );
        uint256 id = zombies.length - 1;
        zombieToOwner[id] = msg.sender;
        ownerZombieCount[msg.sender]++;
        total++;
        emit NewZombie(id, _name, _dna);
    }

    //生成随机dna序列
    function _generateRandomDna(string memory _str)
        private
        view
        returns (uint256)
    {
        return
            uint256(keccak256(abi.encodePacked(_str, block.timestamp))) %
            dnaModulus;
    }

    //免费获取一只僵尸，单个账号限领取一只
    function freeZombie(string memory _name) public {
        require(ownerZombieCount[msg.sender] == 0, "Restricted to receive one");
        uint256 dna = _generateRandomDna(_name);
        //免费僵尸 dna 最后一位标识为 0
        dna = dna - (dna % 10);
        _createZombie(_name, dna);
    }

    //购买僵尸
    function buyZombie(string memory _name) public payable {
        require(
            ownerZombieCount[msg.sender] > 0,
            "Please go and get one for free"
        );
        require(msg.value >= price, "Insufficient purchase amount");
        uint256 dna = _generateRandomDna(_name);
        //付费僵尸 dna 最后一位标识为 1
        dna = dna - (dna % 10) + 1;
        _createZombie(_name, dna);
    }

    //设置僵尸单价
    function setPrice(uint256 _price) external onlyOwner {
        price = _price;
    }
}

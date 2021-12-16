// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ZombieHelper.sol";

//僵尸攻击
contract ZombieAttack is ZombieHelper {
    //随机数种子
    uint256 randNonce = 0;

    //胜率
    uint256 public attackWinRate = 70;

    //生成随机数
    function _rand(uint256 _modulus) internal returns (uint256) {
        randNonce++;
        return
            uint256(
                keccak256(
                    abi.encodePacked(block.timestamp, msg.sender, randNonce)
                )
            ) % _modulus;
    }

    //设置胜率
    function setAttackWinRate(uint256 _rate) public onlyOwner {
        attackWinRate = _rate;
    }

    //僵尸对战
    function attack(uint256 _zombieId, uint256 _targetId)
        external
        onlyOwnerOf(_zombieId)
        returns (uint256)
    {
        require(
            msg.sender != zombieToOwner[_targetId],
            "The target zombie is yours!"
        );
        Zombie storage myZombie = zombies[_zombieId];
        require(_isReady(myZombie), "Your zombie is not ready!");
        Zombie storage enemyZombie = zombies[_targetId];
        uint256 rand = _rand(100);
        if (rand <= attackWinRate) {
            //胜
            myZombie.winCount++;
            myZombie.level++;
            enemyZombie.loseCount++;
            multiply(_zombieId, enemyZombie.dna);
            return _zombieId;
        }
        //输
        myZombie.loseCount++;
        enemyZombie.winCount++;
        _triggerCooldown(myZombie);
        return _targetId;
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ZombieHelper.sol";

//僵尸喂食
contract ZombieFeeding is ZombieHelper {
    //喂食
    function feed(uint256 _zombieId) public onlyOwnerOf(_zombieId) {
        Zombie storage myZombie = zombies[_zombieId];
        require(_isReady(myZombie), "Zombie is not ready");
        zombieFeedCount[_zombieId]++;
        _triggerCooldown(myZombie);
        //满10次奖励新僵尸
        if (zombieFeedCount[_zombieId] % 10 == 0) {
            uint256 dna = myZombie.dna - (myZombie.dna % 10) + 8;
            _createZombie("zombie's son", dna);
        }
    }
}

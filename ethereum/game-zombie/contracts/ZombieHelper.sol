// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ZombieFactory.sol";

//僵尸助手
contract ZombieHelper is ZombieFactory {
    //僵尸升级手续费
    uint256 public levelupFee = 0.001 ether;

    //僵尸等级判断
    modifier aboveLevel(uint256 _level, uint256 _zombieId) {
        require(zombies[_zombieId].level >= _level, "Level is not sufficient");
        _;
    }

    //僵尸拥有者判断
    modifier onlyOwnerOf(uint256 _zombieId) {
        require(msg.sender == zombieToOwner[_zombieId], "Zombie is not yours");
        _;
    }

    //设置升级手续费
    function setLevelupFee(uint256 _fee) external onlyOwner {
        levelupFee = _fee;
    }

    //僵尸升级
    function levelup(uint256 _zombieId)
        external
        payable
        onlyOwnerOf(_zombieId)
    {
        require(msg.value >= levelupFee, "No enough money");
        zombies[_zombieId].level++;
    }

    //修改僵尸名称，等级必须大于二级
    function changeName(uint256 _zombieId, string calldata _newName)
        external
        aboveLevel(2, _zombieId)
        onlyOwnerOf(_zombieId)
    {
        zombies[_zombieId].name = _newName;
    }

    //获取当前账号下所有僵尸
    function getZombiesByOwner(address _owner)
        external
        view
        returns (uint256[] memory)
    {
        uint256[] memory result = new uint256[](ownerZombieCount[_owner]);
        uint256 counter = 0;
        for (uint256 i = 0; i < zombies.length; i++) {
            if (zombieToOwner[i] == _owner) {
                result[counter] = i;
                counter++;
            }
        }
        return result;
    }

    //触发冷却
    function _triggerCooldown(Zombie storage _zombie) internal {
        _zombie.readyTime =
            uint32(block.timestamp + cooldownTime) -
            uint32((block.timestamp + cooldownTime) % 1 days);
    }

    //验证冷却
    function _isReady(Zombie storage _zombie) internal view returns (bool) {
        return (_zombie.readyTime <= block.timestamp);
    }

    //对战胜利生成新僵尸
    function multiply(uint256 _zombieId, uint256 _targetDna)
        internal
        onlyOwnerOf(_zombieId)
    {
        Zombie storage myZombie = zombies[_zombieId];
        require(_isReady(myZombie), "Zombie is not ready");
        _targetDna = _targetDna % dnaModulus;
        uint256 dna = (myZombie.dna + _targetDna) / 2;
        //对战胜利僵尸 dna 最后一位标识 9
        dna = dna - (dna % 10) + 9;
        _createZombie("NoName", dna);
        _triggerCooldown(myZombie);
    }
}

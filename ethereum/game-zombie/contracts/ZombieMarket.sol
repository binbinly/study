// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ZombieOwnership.sol";

//僵尸市场
contract ZombieMarket is ZombieOwnership {
    //卖家
    struct Sale {
        address payable seller; //卖家地址
        uint256 price; //卖出价格
    }

    //商店
    mapping(uint256 => Sale) public shop;

    //商店僵尸数量
    uint256 shopZombieCount;

    //税
    uint256 public tax = 0.001 ether;
    //最低售价
    uint256 public minPrice = 0.001 ether;

    //僵尸入市场事件
    event SaleZombie(uint256 indexed zombieId, address indexed seller);
    //僵尸被买事件
    event BuyShopZombie(
        uint256 indexed zombieId,
        address indexed buyer,
        address indexed seller
    );

    //僵尸进入市场
    function saleMyZombie(uint256 _zombieId, uint256 _price)
        public
        onlyOwnerOf(_zombieId)
    {
        require(_price >= minPrice + tax, "Your price must > minPrice+tax");
        shop[_zombieId] = Sale(payable(msg.sender), _price);
        shopZombieCount += 1;
        emit SaleZombie(_zombieId, msg.sender);
    }

    //购买市场僵尸
    function buyShopZombie(uint256 _zombieId) public payable {
        require(msg.value >= shop[_zombieId].price, "No enough money");
        _transfer(shop[_zombieId].seller, msg.sender, _zombieId);
        shop[_zombieId].seller.transfer(msg.value - tax);
        delete shop[_zombieId];
        shopZombieCount -= 1;
        emit BuyShopZombie(_zombieId, msg.sender, shop[_zombieId].seller);
    }

    //获取市场内所有僵尸
    function getShopZombies() external view returns (uint256[] memory) {
        uint256[] memory result = new uint256[](shopZombieCount);
        uint256 counter = 0;
        for (uint256 i = 0; i < zombies.length; i++) {
            if (shop[i].price > 0) {
                result[counter] = i;
                counter++;
            }
        }
        return result;
    }

    function setTax(uint256 _value) public onlyOwner {
        tax = _value;
    }

    function setMinPrice(uint256 _value) public onlyOwner {
        minPrice = _value;
    }
}

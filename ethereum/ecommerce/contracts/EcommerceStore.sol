// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0;

import './Escrow.sol';

contract EcommerceStore {
    enum ProductStatus {Open, Sold, Unsold}
    enum ProductCondition {New, Used}

    uint public productIndex;
    mapping (uint => address) productIdInStore;
    mapping (address => mapping(uint => Product)) stores;
    mapping (uint => address) productEscrow;

    struct Product {
        uint id;
        string name;
        string category;
        string imageLink;
        string descLink;
        uint auctionStartTime;
        uint auctionEndTime;
        uint startPrice;
        address highestBidder;
        uint highestBid;
        uint secondHighestBid;
        uint totalBids;
        ProductStatus status;
        ProductCondition condition;
    }

    struct Bid {
        address bidder;
        uint productId;
        uint value;
        bool revealed;
    }
    mapping(uint => mapping (address => mapping (bytes32 => Bid))) bids;

    constructor() {
        productIndex = 0;
    }

    /// 添加商品
    function addProductToStore(string memory _name, string memory _category, string memory _imageLink, string memory _descLink
        , uint _auctionStartTime, uint _auctionEndTime
        , uint _startPrice, uint _productCondition) public{
        
        productIndex += 1;
        stores[msg.sender][productIndex] = Product({
            id: productIndex,
            name: _name,
            category: _category,
            imageLink: _imageLink,
            descLink: _descLink,
            auctionStartTime: _auctionStartTime,
            auctionEndTime: _auctionEndTime,
            startPrice: _startPrice,
            highestBidder: address(0),
            highestBid: 0,
            secondHighestBid: 0,
            totalBids:0,
            status: ProductStatus.Open,
            condition: ProductCondition(_productCondition)
        });
        productIdInStore[productIndex] = msg.sender;
    }

    /// 获取商品信息
    function getProduct(uint _productId) public view returns (uint, string memory, string memory, string memory, string memory, uint, uint, uint, ProductStatus, ProductCondition){
        Product memory product = stores[productIdInStore[_productId]][_productId];
        return (product.id, product.name, product.category, product.imageLink, product.descLink
            , product.auctionStartTime, product.auctionEndTime, product.startPrice, product.status, product.condition);
    }

    /// 对商品出价
    function bid(uint _productId, bytes32 _bid) public payable returns(bool) {
        Product storage product = stores[productIdInStore[_productId]][_productId];
        require(block.timestamp >= product.auctionStartTime, "Current time should be later than auction start time");
        require(block.timestamp <= product.auctionEndTime, "Current time should be earlier than auction end time");
        require(msg.value > product.startPrice, "Value should be larger than start price");
        require(bids[_productId][msg.sender][_bid].bidder == address(0), 'Bidder should be bull');
        bids[_productId][msg.sender][_bid] = Bid({
            bidder: msg.sender,
            productId: _productId,
            value: msg.value,
            revealed: false
        });
        product.totalBids += 1;
        return true;
    }

    function revealBid(uint _productId, string memory _amount, string memory _secret) public {
        Product storage product = stores[productIdInStore[_productId]][_productId];
        require(block.timestamp >= product.auctionEndTime);
        bytes32 sealedBid = keccak256(abi.encodePacked(_amount, _secret));
        Bid memory bidInfo = bids[_productId][msg.sender][sealedBid];
        require(bidInfo.bidder != address(0), "Bidder should exist");
        require(bidInfo.revealed == false, "Bid should not be revealed");

        uint refund;
        uint amount = stringToUint(_amount);
        if (bidInfo.value < amount) {
            refund = bidInfo.value;
        } else {
            if (product.highestBidder == address(0)) {
                product.highestBidder = msg.sender;
                product.highestBid = amount;
                product.secondHighestBid = product.startPrice;
                refund = bidInfo.value - amount;
            } else {
                if (amount > product.highestBid) {
                    product.secondHighestBid = product.highestBid;
                    payable(product.highestBidder).transfer(product.highestBid);
                    product.highestBid = amount;
                    product.highestBidder = msg.sender;
                    refund = bidInfo.value - amount;
                } else if (amount > product.secondHighestBid) {
                    product.secondHighestBid = amount;
                    refund = bidInfo.value;
                } else {
                    refund = bidInfo.value;
                }
            }
        }
        bids[_productId][msg.sender][sealedBid].revealed = true;
        if (refund > 0) {
            payable(msg.sender).transfer(refund);
        }
    }

    function stringToUint(string memory s) private pure returns(uint) {
        bytes memory b = bytes(s);
        uint result = 0;
        for (uint i = 0; i < b.length; i++) {
            if (b[i] >= 0x30 && b[i] <= 0x39) {
                result = result * 10 + (uint8(b[i]) - 48);
            }
        }
        return result;
    }

    function highestBidderInfo(uint _productId) public view returns(address, uint, uint) {
        Product memory product = stores[productIdInStore[_productId]][_productId];
        return (product.highestBidder, product.highestBid, product.secondHighestBid);
    }

    function totalBids(uint _productId) public view returns(uint) {
        Product memory product = stores[productIdInStore[_productId]][_productId];
        return product.totalBids;
    }

    function finalizeAuction(uint _productId) public {
        Product storage product = stores[productIdInStore[_productId]][_productId];
        require((block.timestamp > product.auctionEndTime), "Current time should be later than auction end time");
        require(product.status == ProductStatus.Open, "Product status should be open");
        require(msg.sender != productIdInStore[_productId], "Caller should not be seller");
        require(msg.sender != product.highestBidder, "Caller should not be buyer");

        if (product.highestBidder == address(0)) {
            product.status = ProductStatus.Unsold;
        } else {
            Escrow escrow = new Escrow{value: product.secondHighestBid}(_productId, payable(productIdInStore[_productId]), payable(product.highestBidder), msg.sender);
            productEscrow[_productId] = address(escrow);
            product.status = ProductStatus.Sold;
            uint refund = product.highestBid - product.secondHighestBid;
            if (refund > 0)
                payable(product.highestBidder).transfer(refund);
        }
    }

    function escrowAddressForProduct(uint _productId) public view returns(address) {
        return productEscrow[_productId];
    }

    function escrowInfo(uint _productId) public view returns(address, address, address, bool, uint, uint) {
        return Escrow(productEscrow[_productId]).escrowInfo();
    }

    function releaseAmountToSeller(uint _productId) public {
        Escrow(productEscrow[_productId]).realseAmountToSeller(msg.sender);
    }

    function refundAmountToBuyer(uint _productId) public {
        Escrow(productEscrow[_productId]).refundAmountToBuyer(msg.sender);
    }
}
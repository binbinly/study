import React, { Component } from "react";

export default class Buy extends Component {
  buyShopZombie = () => {
    const { id, price } = this.props;
    const abi = React.$abi;
    abi.buyShopZombie(id, price).then((transactionHash) => {
      console.log("hash", transactionHash);
    });
  };

  render() {
    const { price } = this.props;
    return (
      <div>
        <div className="zombieInput">售价：{price} ether</div>
        <div>
          <button className="pay-btn pay-btn-last" onClick={this.buyShopZombie}>
            <span>买下它</span>
          </button>
        </div>
      </div>
    );
  }
}

import React, { Component } from "react";

export default class Sale extends Component {
  state = { price: "" };
  sale = () => {
    const { id, min, price } = this.props;
    const abi = React.$abi;
    let curPrice = this.state.price || price;
    if (curPrice >= min) {
      abi.saleMyZombie(id, curPrice).then((transactionHash) => {
        console.log("hash", transactionHash);
      });
    } else {
      alert("价格必须大于最小价格限制");
    }
  };

  setPrice = (event) => {
    this.setState({
      price: event.target.value,
    });
  };

  render() {
    const { price } = this.props;
    return (
      <div>
        <div className="zombieInput">
          <input
            type="text"
            id="salePrice"
            value={this.state.price}
            placeholder={price}
            onChange={this.setPrice}
          ></input>
        </div>
        <div>
          <button className="pay-btn pay-btn-last" onClick={this.sale}>
            <span>卖了它</span>
          </button>
        </div>
      </div>
    );
  }
}

import React, { Component } from "react";

export default class Feed extends Component {
  feed = () => {
    const { id } = this.props;
    const abi = React.$abi;
    abi.feed(id).then((transactionHash) => {
      console.log("hash", transactionHash);
    });
  };

  render() {
    return (
      <div>
        <button className="pay-btn" onClick={this.feed}>
          <span>喂食一次</span>
        </button>
      </div>
    );
  }
}

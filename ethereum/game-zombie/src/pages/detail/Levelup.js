import React, { Component } from "react";

export default class Levelup extends Component {
  levelUp = () => {
    const { id } = this.props;
    const abi = React.$abi;
    abi.levelUp(id).then((transactionHash) => {
      console.log("hash", transactionHash);
    });
  };

  render() {
    return (
      <div>
        <button className="pay-btn" onClick={this.levelUp}>
          <span>付费升级</span>
        </button>
      </div>
    );
  }
}

import React, { Component } from "react";

export default class Rename extends Component {
  state = { name: "" };

  onChangeName = (event) => {
    this.setState({
      name: event.target.value,
    });
  };

  changeName = () => {
    if (!this.state.name) {
      return alert("请输入新名称");
    }
    const { id } = this.props;
    const abi = React.$abi;
    abi.changeName(id, this.state.name).then((transactionHash) => {
      console.log("hash", transactionHash);
    });
  };

  render() {
    const { name } = this.props;
    return (
      <div>
        <div className="zombieInput">
          <input
            type="text"
            id="zombieName"
            value={this.state.name}
            placeholder={name}
            onChange={this.onChangeName}
          ></input>
        </div>
        <div>
          <button className="pay-btn pay-btn-last" onClick={this.changeName}>
            <span>改个名字</span>
          </button>
        </div>
      </div>
    );
  }
}

import React, { Component } from "react";
import { Link } from "react-router-dom";
import Card from "../components/Card";
import "../static/ZombiePreview.css";

export default class My extends Component {
  state = {
    zombieCount: "",
    zombies: [],
    zombieName: "",
    transactionHash: "",
    buyAreaDisp: 1,
    createAreaDisp: 1,
    txHashDisp: 0,
  };

  componentDidMount() {
    const abi = React.$abi;
    abi.init().then(() => {
      abi.getZombiesByOwner().then((zombies) => {
        if (zombies.length > 0) {
          let _zombies = this.state.zombies;
          for (let i = 0; i < zombies.length; i++) {
            abi.zombies(zombies[i]).then((result) => {
              result.zombieId = zombies[i];
              _zombies.push(result);
              this.setState({ zombies: _zombies });
            });
          }
        }
      });
    });
  }

  createZombie = () => {
    const abi = React.$abi;
    abi.createZombie(this.state.zombieName).then((transactionHash) => {
      this.setState({
        transactionHash: transactionHash,
        createAreaDisp: 0,
        txHashDisp: 1,
      });
    });
  };

  buyZombie = () => {
    const abi = React.$abi;
    abi.buyZombie(this.state.zombieName).then((transactionHash) => {
      this.setState({
        transactionHash: transactionHash,
        buyAreaDisp: 0,
        txHashDisp: 1,
      });
    });
  };

  inputChange = (event) => {
    console.log("event", event);
    this.setState({ zombieName: event.target.value });
  };

  render() {
    if (this.state.zombies.length > 0) {
      return (
        <div className="cards">
          {this.state.zombies.map((item, index) => {
            var name = item.name;
            var level = item.level;
            return (
              <Link to={`/detail/` + item.zombieId} key={index}>
                <Card
                  zombie={item}
                  name={name}
                  level={level}
                  key={index}
                ></Card>
              </Link>
            );
          })}
          <div className="buyArea" display={this.state.buyAreaDisp}>
            <div className="zombieInput">
              <input
                type="text"
                id="zombieName"
                placeholder="给僵尸起个好名字"
                value={this.state.zombieName}
                onChange={this.inputChange}
              ></input>
            </div>
            <div>
              <button className="attack-btn" onClick={this.buyZombie}>
                <span>购买僵尸</span>
              </button>
            </div>
          </div>
          <div className="transactionHash" display={this.state.txHashDisp}>
            {this.state.transactionHash}
            <br></br>等待确认中...
          </div>
        </div>
      );
    } else {
      return (
        <div>
          <div className="createArea" display={this.state.createAreaDisp}>
            <div className="zombieInput">
              <input
                type="text"
                id="zombieName"
                placeholder="给僵尸起个好名字"
                value={this.state.zombieName}
                onChange={this.inputChange}
              ></input>
            </div>
            <div>
              <button className="attack-btn" onClick={this.createZombie}>
                <span>免费领养僵尸</span>
              </button>
            </div>
          </div>
          <div className="transactionHash" display={this.state.txHashDisp}>
            {this.state.transactionHash}
            <br></br>等待确认中...
          </div>
        </div>
      );
    }
  }
}

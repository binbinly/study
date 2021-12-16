import React, { Component } from "react";
import "../static/ZombiePreview.css";
import moment from "moment";
import Preview from "../components/Preview";

export default class Attack extends Component {
  state = {
    targetZombie: {},
    myZombies: [],
    myZombie: {},
    myZombieId: "",
    active: {},
    buttonTxt: "",
    modalDisplay: "none",
    transactionHash: "",
    AttackBtn: () => {
      return (
        <button className="attack-btn">
          <span role="img" aria-label="zombie">
            选一只🧟‍♂️干它！
          </span>
        </button>
      );
    },
  };

  componentDidMount() {
    const { id } = this.props.match.params;
    const abi = React.$abi;
    abi.init().then(() => {
      this.getZombie(id);
      this.getMyZombies();
    });
  }

  getZombie = (zombieId) => {
    const abi = React.$abi;
    abi.zombies(zombieId).then((result) => {
      this.setState({ targetZombie: result });
    });
  };
  getMyZombies = () => {
    const abi = React.$abi;
    abi.getZombiesByOwner().then((zombies) => {
      if (zombies.length > 0) {
        for (let i = 0; i < zombies.length; i++) {
          abi.zombies(zombies[i]).then((result) => {
            let _zombies = this.state.myZombies;
            result.zombieId = zombies[i];
            if (
              result.readyTime === 0 ||
              moment().format("X") > result.readyTime
            ) {
              _zombies.push(result);
            }
            this.setState({ myZombies: _zombies });
          });
        }
      }
    });
  };
  selectZombie = (index) => {
    var _active = this.state.active;
    var prev_active = _active[index];
    for (var i = 0; i < this.state.myZombies.length; i++) {
      _active[i] = 0;
    }
    _active[index] = prev_active === 0 || prev_active === undefined ? 1 : 0;
    this.setState({
      active: _active,
      buttonTxt: "用" + this.state.myZombies[index].name,
      myZombie: this.state.myZombies[index],
      myZombieId: this.state.myZombies[index].zombieId,
      AttackBtn: () => {
        return (
          <button className="attack-btn" onClick={this.zombieAttack}>
            <span role="img">用{this.state.myZombies[index].name}干它！</span>
          </button>
        );
      },
    });
  };

  zombieAttack = () => {
    if (this.state.myZombie !== undefined) {
      const abi = React.$abi;
      const { id } = this.props.match.params;
      this.setState({ modalDisplay: "" });
      abi.attack(this.state.myZombieId, id).then((transactionHash) => {
        this.setState({
          transactionHash: transactionHash,
          AttackBtn: () => {
            return <div></div>;
          },
        });
      });
    }
  };
  render() {
    let AttackBtn = this.state.AttackBtn;
    if (this.state.myZombies.length > 0) {
      return (
        <div className="App zombie-attack">
          <div
            className="modal"
            style={{
              display: this.state.modalDisplay,
            }}
          >
            <div className="battelArea">
              <div className="targetZombie">
                <Preview zombie={this.state.targetZombie} />
              </div>
              <div className="vs">VS</div>
              <div className="myZombie">
                <Preview zombie={this.state.myZombie} />
              </div>
            </div>
            <div>
              <h2>{this.state.transactionHash}</h2>
            </div>
          </div>
          <div className="row zombie-parts-bin-component">
            <div className="game-card home-card target-card">
              <div className="zombie-char">
                <Preview zombie={this.state.targetZombie} />
              </div>
            </div>
            <div className="zombie-detail">
              <div className="flex">
                {this.state.myZombies.map((item, index) => {
                  var name = item.name;
                  var level = item.level;
                  return (
                    <div
                      className="game-card home-card selectable"
                      key={index}
                      active={this.state.active[index] || 0}
                      onClick={() => this.selectZombie(index)}
                    >
                      <div className="zombie-char">
                        <Preview zombie={item} />
                        <div className="zombie-card card bg-shaded">
                          <div className="card-header bg-dark hide-overflow-text">
                            <strong>{name}</strong>
                          </div>
                          <small className="hide-overflow-text">
                            CryptoZombie{level}级
                          </small>
                        </div>
                      </div>
                    </div>
                  );
                })}
              </div>
              <AttackBtn></AttackBtn>
            </div>
          </div>
        </div>
      );
    } else {
      return <div>没有能干它的僵尸</div>;
    }
  }
}

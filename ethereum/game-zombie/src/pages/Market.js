import React, { Component } from "react";
import { Link } from "react-router-dom";
import "../static/ZombiePreview.css";
import Card from "../components/Card";

export default class Market extends Component {
  state = { shopZombies: [] };

  componentDidMount() {
    const abi = React.$abi;
    abi.init().then((res) => {
      abi.getShopZombies().then((zombieIds) => {
        if (zombieIds.length > 0) {
          for (var i = 0; i < zombieIds.length; i++) {
            let zombieId = zombieIds[i];
            if (zombieId >= 0) {
              abi.zombies(zombieId).then((zombies) => {
                let _shopZombies = this.state.shopZombies;
                zombies.zombieId = zombieId;
                _shopZombies.push(zombies);
                this.setState({ shopZombies: _shopZombies });
              });
            }
          }
        }
      });
    });
  }

  render() {
    if (this.state.shopZombies.length > 0) {
      return (
        <div className="cards">
          {this.state.shopZombies.map((item, index) => {
            var name = item.name;
            var level = item.level;
            return (
              <Link to={`/detail/` + item.zombieId} key={index}>
                <Card zombie={item} name={name} level={level} key={index} />
              </Link>
            );
          })}
        </div>
      );
    } else {
      return <div></div>;
    }
  }
}

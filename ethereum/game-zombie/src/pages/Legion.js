import React, { Component } from "react";
import { Link } from "react-router-dom";
import Card from "../components/Card";
import "../static/ZombiePreview.css";

export default class Legion extends Component {
  state = { count: "", zombies: [] };

  componentDidMount() {
    const abi = React.$abi;
    abi.init().then(() => {
      abi.zombieCount().then((count) => {
        if (count > 0) {
          let zombies = [];
          for (let i = 0; i < count; i++) {
            abi.zombies(i).then((res) => {
              res.zombieId = i;
              zombies.push(res);
              this.setState({ zombies });
            });
          }
        }
      });
    });
  }

  render() {
    if (this.state.zombies.length > 0) {
      return (
        <div className="cards">
          {this.state.zombies.map((item, index) => {
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

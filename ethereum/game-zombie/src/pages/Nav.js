import React, { Component, Fragment } from "react";
import { Link } from "react-router-dom";

export default class Nav extends Component {
  state = {
    list: [
      { name: "僵尸军团", to: "/legion" },
      { name: "我的僵尸", to: "/my" },
      { name: "僵尸市场", to: "/market" },
      { name: "基因模拟", to: "/gene" },
    ],
  };

  componentDidMount() {
    const abi = React.$abi;
    abi.init().then(() => {
      abi.owner().then((owner) => {
        if (abi.isMyself(owner)) {
          this.setState({
            list: [...this.state.list, { name: "系统管理", to: "/admin" }],
          });
        }
      });
    });
  }

  render() {
    return (
      <Fragment>
        {this.state.list.map((item, index) => {
          return (
            <li key={index}>
              <button className="start-course-btn">
                <span>
                  <Link to={item.to}>{item.name}</Link>
                </span>
              </button>
            </li>
          );
        })}
      </Fragment>
    );
  }
}

import React, { Component } from "react";
import { Link } from "react-router-dom";

export default class Attack extends Component {
  render() {
    const { id } = this.props;
    return (
      <button className="attack-btn">
        <span>
          <Link to={`/attack/` + id}>发起挑战</Link>
        </span>
      </button>
    );
  }
}

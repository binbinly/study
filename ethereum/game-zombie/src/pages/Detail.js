import React, { Component } from "react";
import "../static/ZombiePreview.css";
import Preview from "../components/Preview";
import moment from "moment";
import Attack from "./detail/Attack";
import Buy from "./detail/Buy";
import Feed from "./detail/Feed";
import Levelup from "./detail/Levelup";
import Rename from "./detail/Rename";
import Sale from "./detail/Sale";

export default class Detail extends Component {
  state = {
    zombie: {},
    owner: "",
    zombieFeedCount: 0,
    myPrice: 0,
    minPrice: 0,
    onShop: false,
    shopInfo: {},
  };

  componentDidMount() {
    const { id } = this.props.match.params;
    const abi = React.$abi;
    abi.init().then(() => {
      this.getZombie(id);
      this.getZombieFeedCount(id);
      this.getMinPrice();
      this.getZombieShop(id);
    });
  }
  getZombieShop = (zombieId) => {
    const abi = React.$abi;
    abi.zombieShop(zombieId).then((shopInfo) => {
      if (shopInfo.price > 0) {
        this.setState({ onShop: true, shopInfo: shopInfo });
      }
    });
  };
  getMinPrice = () => {
    const abi = React.$abi;
    abi.minPrice().then((minPrice) => {
      if (minPrice > 0) {
        abi.tax().then((tax) => {
          if (tax > 0) {
            this.setState({
              myPrice: parseFloat(minPrice) + parseFloat(tax),
              minPrice: parseFloat(minPrice) + parseFloat(tax),
            });
          }
        });
      }
    });
  };
  getZombieFeedCount = (zombieId) => {
    const abi = React.$abi;
    abi.zombieFeedCount(zombieId).then((result) => {
      if (result > 0) {
        this.setState({ zombieFeedCount: result });
      }
    });
  };

  getZombie = (zombieId) => {
    const abi = React.$abi;
    abi.zombies(zombieId).then((result) => {
      this.setState({ zombie: result });
      this.setState({ zombieNewname: result.name });
      abi.zombieToOwner(zombieId).then((zombieOwner) => {
        this.setState({ owner: zombieOwner });
      });
    });
  };

  render() {
    const { id } = this.props.match.params;
    const abi = React.$abi;
    var readyTime = "已冷却";
    if (
      this.state.zombie.readyTime !== undefined &&
      moment().format("X") < this.state.zombie.readyTime
    ) {
      readyTime = moment(parseInt(this.state.zombie.readyTime) * 1000).format(
        "YYYY-MM-DD"
      );
    }
    return (
      <div className="App">
        <div
          className="row zombie-parts-bin-component"
          authenticated="true"
          lesson="1"
          lessonidx="1"
        >
          <div className="zombie-preview" id="zombie-preview">
            <div className="zombie-char">
              <div
                className="zombie-loading zombie-parts"
                style={{ display: "none" }}
              ></div>
              <Preview zombie={this.state.zombie} />
              <div className="hide">
                <div className="card-header bg-dark hide-overflow-text">
                  <strong></strong>
                </div>
                <small className="hide-overflow-text">CryptoZombie第一级</small>
              </div>
            </div>
          </div>
          <div className="zombie-detail">
            <dl>
              <dt>{this.state.zombie.name}</dt>
              <dt>主人</dt>
              <dd>{this.state.owner}</dd>
              <dt>等级</dt>
              <dd>{this.state.zombie.level}</dd>
              <dt>胜利次数</dt>
              <dd>{this.state.zombie.winCount}</dd>
              <dt>失败次数</dt>
              <dd>{this.state.zombie.loseCount}</dd>
              <dt>冷却时间</dt>
              <dd>{readyTime}</dd>
              <dt>喂食次数</dt>
              <dd>{this.state.zombieFeedCount}</dd>
              <dt></dt>
              <dd>
                {abi.isMyself(this.state.owner) ? (
                  <>
                    {this.state.zombie.level > 1 ? (
                      <Rename id={id} name={this.state.zombie.name} />
                    ) : null}
                    <Feed id={id} />
                    <Levelup id={id} />
                    <Sale
                      id={id}
                      min={this.state.minPrice}
                      price={this.state.myPrice}
                    />
                  </>
                ) : (
                  <>
                    <Attack id={id} />
                    {this.state.onShop ? (
                      <Buy id={id} price={this.state.shopInfo.price} />
                    ) : null}
                  </>
                )}
              </dd>
            </dl>
          </div>
        </div>
      </div>
    );
  }
}

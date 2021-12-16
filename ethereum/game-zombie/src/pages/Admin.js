import React, { Component } from "react";

export default class Admin extends Component {
  state = {
    contractAddress: "",
    contractOwner: "",
    ContractName: "",
    ContractSymbol: "",
    ContractBalance: "",
    attackVictoryProbability: 70,
    levelUpFee: 0,
    minPrice: 0,
    tax: 0,
    zombiePrice: 0,
  };

  componentDidMount() {
    const abi = React.$abi;
    abi.init().then(() => {
      abi.owner().then((contractOwner) => {
        if (abi.isMyself(contractOwner)) {
          this.setState({ contractOwner: contractOwner });
          this.getContractName();
          this.getContractSymbol();
          this.checkBalance();
          this.levelUpFee();
          this.minPrice();
          this.tax();
          this.zombiePrice();
          this.setState({
            contractAddress: abi.contract.options.address,
          });
        }
      });
    });
  }
  getContractName = () => {
    const abi = React.$abi;
    abi.name().then((name) => {
      this.setState({ ContractName: name });
    });
  };

  getContractSymbol = () => {
    const abi = React.$abi;
    abi.symbol().then((symbol) => {
      this.setState({ ContractSymbol: symbol });
    });
  };

  checkBalance = () => {
    const abi = React.$abi;
    abi.checkBalance().then((balance) => {
      this.setState({ ContractBalance: balance });
    });
  };
  withdraw = () => {
    const abi = React.$abi;
    abi.withdraw().then(() => {
      this.checkBalance();
    });
  };
  setAttackWinRate = () => {
    const abi = React.$abi;
    abi.setAttackWinRate(this.state.attackVictoryProbability).then((res) => {
      console.log(res);
      window.location.reload();
    });
  };
  inputAttackVictoryProbability = (event) => {
    this.setState({
      attackVictoryProbability: event.target.value,
    });
  };
  levelUpFee = () => {
    const abi = React.$abi;
    abi.levelUpFee().then((levelUpFee) => {
      this.setState({ levelUpFee: levelUpFee });
    });
  };
  setLevelUpFee = () => {
    const abi = React.$abi;
    abi.setLevelUpFee(this.state.levelUpFee).then((res) => {
      console.log(res);
      window.location.reload();
    });
  };
  inputLevelUpFee = (event) => {
    this.setState({
      levelUpFee: event.target.value,
    });
  };
  minPrice = () => {
    const abi = React.$abi;
    abi.minPrice().then((minPrice) => {
      this.setState({ minPrice: minPrice });
    });
  };
  setMinPrice = () => {
    const abi = React.$abi;
    abi.setMinPrice(this.state.minPrice).then((res) => {
      console.log(res);
      window.location.reload();
    });
  };
  inputMinPrice = (event) => {
    this.setState({
      minPrice: event.target.value,
    });
  };
  tax = () => {
    const abi = React.$abi;
    abi.tax().then((tax) => {
      this.setState({ tax: tax });
    });
  };
  setTax = () => {
    const abi = React.$abi;
    abi.setTax(this.state.tax).then((res) => {
      console.log(res);
      window.location.reload();
    });
  };
  inputTax = (event) => {
    this.setState({
      tax: event.target.value,
    });
  };
  zombiePrice = () => {
    const abi = React.$abi;
    abi.zombiePrice().then((zombiePrice) => {
      this.setState({ zombiePrice: zombiePrice });
    });
  };
  setZombiePrice = () => {
    const abi = React.$abi;
    abi.setZombiePrice(this.state.zombiePrice).then((res) => {
      console.log(res);
      window.location.reload();
    });
  };
  inputZombiePrice = (event) => {
    this.setState({
      zombiePrice: event.target.value,
    });
  };

  render() {
    const abi = React.$abi;
    if (abi.isMyself(this.state.contractOwner)) {
      return (
        <div className="contract-admin">
          <dl>
            <dt>合约地址</dt>
            <dd className="lowcase">{this.state.contractAddress}</dd>
            <dt>管理员</dt>
            <dd className="lowcase">{this.state.contractOwner}</dd>
            <dt>合约名称</dt>
            <dd>{this.state.ContractName}</dd>
            <dt>合约标识</dt>
            <dd>{this.state.ContractSymbol}</dd>
            <dt>合约余额</dt>
            <dd>
              {this.state.ContractBalance}
              <button className="pay-btn pay-btn-last" onClick={this.withdraw}>
                <span>提款</span>
              </button>
            </dd>
            <dt>对战胜率</dt>
            <dd>
              <input
                type="text"
                id="attackVictoryProbability"
                value={this.state.attackVictoryProbability}
                onChange={this.inputAttackVictoryProbability}
              ></input>
              <button
                className="pay-btn pay-btn-last"
                onClick={this.setAttackWinRate}
              >
                <span>设置</span>
              </button>
            </dd>
            <dt>升级费</dt>
            <dd>
              <input
                type="text"
                id="levelUpFee"
                value={this.state.levelUpFee}
                onChange={this.inputLevelUpFee}
              ></input>
              <button
                className="pay-btn pay-btn-last"
                onClick={this.setLevelUpFee}
              >
                <span>设置</span>
              </button>
            </dd>
            <dt>最低售价</dt>
            <dd>
              <input
                type="text"
                id="minPrice"
                value={this.state.minPrice}
                onChange={this.inputMinPrice}
              ></input>
              <button
                className="pay-btn pay-btn-last"
                onClick={this.setMinPrice}
              >
                <span>设置</span>
              </button>
            </dd>
            <dt>税金</dt>
            <dd>
              <input
                type="text"
                id="tax"
                value={this.state.tax}
                onChange={this.inputTax}
              ></input>
              <button className="pay-btn pay-btn-last" onClick={this.setTax}>
                <span>设置</span>
              </button>
            </dd>
            <dt>僵尸售价</dt>
            <dd>
              <input
                type="text"
                id="zombiePrice"
                value={this.state.zombiePrice}
                onChange={this.inputZombiePrice}
              ></input>
              <button
                className="pay-btn pay-btn-last"
                onClick={this.setZombiePrice}
              >
                <span>设置</span>
              </button>
            </dd>
          </dl>
        </div>
      );
    } else {
      return <div></div>;
    }
  }
}

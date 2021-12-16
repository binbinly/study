import Web3 from "web3";
import abi from "../abi/ZombieCore.json";
import contractAddress from "../config/index";

class zombie {
  constructor(arg) {
    this.web3 = null;
    //当前账号
    this.account = null;
    //合约
    this.contract = null;
    this.status = 0;
    this.eventEth();
  }

  eventEth() {
    let ethereum = window.ethereum;
    if (typeof ethereum !== "undefined" || typeof window.web3 !== "undefined") {
      ethereum.on("accountsChanged", function (accounts) {
        console.log("accountsChanged:" + accounts);
        window.location.reload();
      });
      ethereum.on("chainChanged", function (chainId) {
        console.log("chainChanged:" + chainId);
        window.location.reload();
      });
      ethereum.on("networkChanged", function (networkVersion) {
        console.log("networkChanged:" + networkVersion);
        window.location.reload();
      });
    } else {
      alert("You have to install MetaMask !");
    }
  }

  async initAsync() {
    //let currentChainId = parseInt(window.ethereum.chainId, 16)
    let ethereum = window.ethereum;
    //禁止自动刷新，metamask要求写的
    ethereum.autoRefreshOnNetworkChange = false;
    let accounts = await window.ethereum.enable();
    console.log("account", accounts);
    //获取到当前默认的以太坊地址
    this.account = accounts[0].toLowerCase();

    //初始化provider
    let provider = window["ethereum"] || this.web3.currentProvider;
    //初始化Web3
    this.web3 = new Web3(provider);
    const chainId = await this.web3.eth.net.getId();

    this.web3.currentProvider.setMaxListeners(300);
    //从json获取到当前网络id下的合约地址
    let address = contractAddress[chainId];
    if (!address) {
      console.log("Undefined Your ChainId:" + chainId);
      return;
    }
    //实例化合约
    this.contract = new this.web3.eth.Contract(abi.abi, address);
    console.log("connect success", this.web3, this.contract, this.account);
  }

  init() {
    return new Promise((resolve, reject) => {
      if (this.status === 1) {
        return resolve(true);
      } else if (this.status === 2) {
        return reject("repeat init");
      }
      //let currentChainId = parseInt(window.ethereum.chainId, 16)
      let ethereum = window.ethereum;
      //禁止自动刷新，metamask要求写的
      ethereum.autoRefreshOnNetworkChange = false;
      //开始调用metamask
      ethereum
        .enable()
        .then((accounts) => {
          console.log("accounts", accounts);
          //初始化provider
          let provider = window["ethereum"] || this.web3.currentProvider;
          console.log("provider", provider);
          //初始化Web3
          this.web3 = new Web3(provider);
          //获取到当前以太坊网络id
          this.web3.eth.net.getId().then((chainId) => {
            //设置最大监听器数量，否则出现warning
            this.web3.currentProvider.setMaxListeners(300);

            //从json获取到当前网络id下的合约地址
            let currentContractAddress = contractAddress[chainId];
            if (currentContractAddress !== undefined) {
              //实例化合约
              this.contract = new this.web3.eth.Contract(
                abi.abi,
                currentContractAddress
              );
              //获取到当前默认的以太坊地址
              this.account = accounts[0].toLowerCase();
              console.log(
                "connect success",
                this.web3,
                this.contract,
                this.account
              );
              this.status = 1;
              resolve(true);
            } else {
              this.status = 2;
              console.log("Undefined Your ChainId:" + chainId);
              reject("Undefined Your ChainId:" + chainId);
            }
          });
        })
        .catch((err) => {
          this.status = 2;
          console.log("ethereum enable err: ", err);
          alert("You have to install MetaMask !");
          reject(err);
        });
    });
  }

  //是否是当前账号
  isMyself(account) {
    return this.account === account;
  }

  //僵尸总量
  zombieCount() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .total()
        .call()
        .then((count) => resolve(count));
    });
  }
  //获得单个僵尸数据
  zombies(zombieId) {
    return new Promise((resolve, reject) => {
      if (zombieId >= 0) {
        this.contract.methods
          .zombies(zombieId)
          .call()
          .then((zombies) => resolve(zombies));
      }
    });
  }
  //获得僵尸拥有者地址
  zombieToOwner(zombieId) {
    return new Promise((resolve, reject) => {
      if (zombieId >= 0) {
        this.contract.methods
          .zombieToOwner(zombieId)
          .call()
          .then((zombies) => resolve(zombies.toLowerCase()));
      }
    });
  }

  //获得当前用户的所有僵尸id
  getZombiesByOwner() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .getZombiesByOwner(this.account)
        .call()
        .then((zombies) => resolve(zombies));
    });
  }

  //创建随机僵尸
  createZombie(_name) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .freeZombie(_name)
        .send({ from: this.account })
        .on("transactionHash", function (transactionHash) {
          resolve(transactionHash);
        })
        .on("confirmation", function (confirmationNumber, receipt) {
          console.log({
            confirmationNumber: confirmationNumber,
            receipt: receipt,
          });
        })
        .on("receipt", function (receipt) {
          console.log({ receipt: receipt });
          window.location.reload();
        })
        .on("error", function (error, receipt) {
          console.log({ error: error, receipt: receipt });
          reject({ error: error, receipt: receipt });
        });
    });
  }

  //购买僵尸
  buyZombie(_name) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .price()
        .call()
        .then((zombiePrice) => {
          this.contract.methods
            .buyZombie(_name)
            .send({ from: this.account, value: zombiePrice })
            .on("transactionHash", function (transactionHash) {
              resolve(transactionHash);
            })
            .on("confirmation", function (confirmationNumber, receipt) {
              console.log({
                confirmationNumber: confirmationNumber,
                receipt: receipt,
              });
            })
            .on("receipt", function (receipt) {
              console.log({ receipt: receipt });
              window.location.reload();
            })
            .on("error", function (error, receipt) {
              console.log({ error: error, receipt: receipt });
              reject({ error: error, receipt: receipt });
            });
        });
    });
  }

  //僵尸对战
  attack(_zombieId, _targetId) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .attack(_zombieId, _targetId)
        .send({ from: this.account })
        .on("transactionHash", function (transactionHash) {
          resolve(transactionHash);
        })
        .on("confirmation", function (confirmationNumber, receipt) {
          console.log({
            confirmationNumber: confirmationNumber,
            receipt: receipt,
          });
        })
        .on("receipt", function (receipt) {
          console.log({ receipt: receipt });
          window.location.reload();
        })
        .on("error", function (error, receipt) {
          console.log({ error: error, receipt: receipt });
          reject({ error: error, receipt: receipt });
        });
    });
  }

  //僵尸改名
  changeName(_zombieId, _name) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .changeName(_zombieId, _name)
        .send({ from: this.account })
        .on("transactionHash", function (transactionHash) {
          resolve(transactionHash);
        })
        .on("confirmation", function (confirmationNumber, receipt) {
          console.log({
            confirmationNumber: confirmationNumber,
            receipt: receipt,
          });
        })
        .on("receipt", function (receipt) {
          console.log({ receipt: receipt });
          window.location.reload();
        })
        .on("error", function (error, receipt) {
          console.log({ error: error, receipt: receipt });
          reject({ error: error, receipt: receipt });
        });
    });
  }

  //僵尸喂食
  feed(_zombieId) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .feed(_zombieId)
        .send({ from: this.account })
        .on("transactionHash", function (transactionHash) {
          resolve(transactionHash);
        })
        .on("confirmation", function (confirmationNumber, receipt) {
          console.log({
            confirmationNumber: confirmationNumber,
            receipt: receipt,
          });
        })
        .on("receipt", function (receipt) {
          console.log({ receipt: receipt });
          window.location.reload();
        })
        .on("error", function (error, receipt) {
          console.log({ error: error, receipt: receipt });
          reject({ error: error, receipt: receipt });
        });
    });
  }

  //僵尸付费升级
  levelUp(_zombieId) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .levelupFee()
        .call()
        .then((levelUpFee) => {
          this.contract.methods
            .levelup(_zombieId)
            .send({ from: this.account, value: levelUpFee })
            .on("transactionHash", function (transactionHash) {
              resolve(transactionHash);
            })
            .on("confirmation", function (confirmationNumber, receipt) {
              console.log({
                confirmationNumber: confirmationNumber,
                receipt: receipt,
              });
            })
            .on("receipt", function (receipt) {
              console.log({ receipt: receipt });
              window.location.reload();
            })
            .on("error", function (error, receipt) {
              console.log({ error: error, receipt: receipt });
              reject({ error: error, receipt: receipt });
            });
        });
    });
  }

  //获取僵尸喂食次数
  zombieFeedCount(_zombieId) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .zombieFeedCount(_zombieId)
        .call()
        .then((count) => resolve(count));
    });
  }

  //获取最低售价
  minPrice() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .minPrice()
        .call()
        .then((minPrice) =>
          resolve(this.web3.utils.fromWei(minPrice, "ether"))
        );
    });
  }

  //获取税金
  tax() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .tax()
        .call()
        .then((tax) => resolve(this.web3.utils.fromWei(tax, "ether")));
    });
  }

  //出售我的僵尸
  saleMyZombie(_zombieId, _price) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .saleMyZombie(_zombieId, this.web3.utils.toWei(_price.toString()))
        .send({ from: this.account })
        .on("transactionHash", function (transactionHash) {
          resolve(transactionHash);
        })
        .on("confirmation", function (confirmationNumber, receipt) {
          console.log({
            confirmationNumber: confirmationNumber,
            receipt: receipt,
          });
        })
        .on("receipt", function (receipt) {
          console.log({ receipt: receipt });
          window.location.reload();
        })
        .on("error", function (error, receipt) {
          console.log({ error: error, receipt: receipt });
          reject({ error: error, receipt: receipt });
        });
    });
  }

  //获得商店里僵尸数据
  zombieShop(_zombieId) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .shop(_zombieId)
        .call()
        .then((shopInfo) => {
          shopInfo.price = this.web3.utils.fromWei(shopInfo.price, "ether");
          resolve(shopInfo);
        });
    });
  }

  //获得商店所有僵尸
  getShopZombies() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .getShopZombies()
        .call()
        .then((zombieIds) => resolve(zombieIds));
    });
  }

  //购买商店里的僵尸
  buyShopZombie(_zombieId, _price) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .buyShopZombie(_zombieId)
        .send({
          from: this.account,
          value: this.web3.utils.toWei(_price.toString()),
        })
        .on("transactionHash", function (transactionHash) {
          resolve(transactionHash);
        })
        .on("confirmation", function (confirmationNumber, receipt) {
          console.log({
            confirmationNumber: confirmationNumber,
            receipt: receipt,
          });
        })
        .on("receipt", function (receipt) {
          console.log({ receipt: receipt });
          window.location.reload();
        })
        .on("error", function (error, receipt) {
          console.log({ error: error, receipt: receipt });
          reject({ error: error, receipt: receipt });
        });
    });
  }

  //获得合约拥有者地址
  owner() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .owner()
        .call()
        .then((owner) => resolve(owner.toLowerCase()));
    });
  }

  //获得合约名称
  name() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .name()
        .call()
        .then((name) => resolve(name));
    });
  }

  //获得合约标识
  symbol() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .symbol()
        .call()
        .then((symbol) => resolve(symbol));
    });
  }

  //查询余额
  checkBalance() {
    return new Promise((resolve, reject) => {
      this.owner().then((owner) => {
        if (this.account === owner) {
          this.contract.methods
            .balance()
            .call({ from: this.account })
            .then((balance) => {
              resolve(this.web3.utils.fromWei(balance, "ether"));
            });
        } else {
          reject("You are not contract owner");
        }
      });
    });
  }

  //设置攻击胜率
  setAttackWinRate(rate) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .setAttackWinRate(rate)
        .send({ from: this.account })
        .then((result) => resolve(result));
    });
  }

  //获得升级费
  levelUpFee() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .levelupFee()
        .call()
        .then((levelUpFee) =>
          resolve(this.web3.utils.fromWei(levelUpFee, "ether"))
        );
    });
  }

  //设置升级费
  setLevelUpFee(_fee) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .setLevelupFee(this.web3.utils.toWei(_fee.toString()))
        .send({ from: this.account })
        .then((result) => resolve(result));
    });
  }

  //设置最低售价
  setMinPrice(_value) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .setMinPrice(this.web3.utils.toWei(_value.toString()))
        .send({ from: this.account })
        .then((result) => resolve(result));
    });
  }

  //获得僵尸售价
  zombiePrice() {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .price()
        .call()
        .then((price) => resolve(this.web3.utils.fromWei(price, "ether")));
    });
  }

  //设置僵尸售价
  setZombiePrice(_value) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .setPrice(this.web3.utils.toWei(_value.toString()))
        .send({ from: this.account })
        .then((result) => resolve(result));
    });
  }

  //设置税金
  setTax(_value) {
    return new Promise((resolve, reject) => {
      this.contract.methods
        .setTax(this.web3.utils.toWei(_value.toString()))
        .send({ from: this.account })
        .then((result) => resolve(result));
    });
  }

  //提款
  withdraw() {
    return new Promise((resolve, reject) => {
      this.owner().then((owner) => {
        if (this.account === owner) {
          this.contract.methods
            .withdraw()
            .send({ from: this.account })
            .then((result) => resolve(result));
        } else {
          reject("You are not contract owner");
        }
      });
    });
  }

  //新僵尸事件
  EventNewZombie() {
    return this.contract.events.NewZombie(
      {},
      { fromBlock: 0, toBlock: "latest" }
    );
  }

  //出售僵尸事件
  EventSaleZombie() {
    return new Promise((resolve, reject) => {
      this.contract.events.SaleZombie(
        { fromBlock: 0, toBlock: "latest" },
        function (error, event) {
          resolve(event);
        }
      );
    });
  }

  //所有事件
  allEvents() {
    this.contract.events
      .allEvents({ fromBlock: 0 }, function (error, event) {
        console.log({ allEvents: event });
      })
      .on("connected", function (subscriptionId) {
        console.log({ connected_subscriptionId: subscriptionId });
      })
      .on("data", function (event) {
        console.log({ event_data: event });
      })
      .on("changed", function (event) {
        console.log({ event_changed: event });
      })
      .on("error", function (error, receipt) {
        console.log({ event_error: error, receipt: receipt });
      });
  }
}

export default zombie;

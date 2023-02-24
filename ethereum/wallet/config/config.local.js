const Web3 = require("web3");
const path = require("path");

module.exports = (appInfo) => {
  return {
    mysql: {
      // 单数据库信息配置
      client: {
        // host
        host: "192.168.8.76",
        // 端口号
        port: "3306",
        // 用户名
        user: "root",
        // 密码
        password: "123456",
        // 数据库名
        database: "wallet",
      },
      // 是否加载到 app 上，默认开启
      app: true,
      // 是否加载到 agent 上，默认关闭
      agent: false,
    },
    redis: {
      client: {
        port: 6379, // Redis port
        host: "192.168.8.76", // Redis host
        password: "",
        db: 0,
      },
    },
    tron: {
      api: "https://nile.trongrid.io",
      apiKey: "18c19d1b-edac-4fae-80f8-411e74d8afb8",
      usdt: "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj",
      gasAddress: "THzTkg7ydSH4sF1TRM6ZhbVovVDc9ptrtd",
    },
    bsc: {
      chainId: 97,
      api: "https://api-testnet.bscscan.com/api",
      apiKey: "GQEIXEWICMNP4ECIVS1BQFC2S4N6YHWRAB",
      web3Url: "https://data-seed-prebsc-2-s2.binance.org:8545",
      usdt: "0x7ef95a0FEE0Dd31b22626fA2e10Ee6A223F8a684",
    },
    eth: {
      web3Url: "https://rinkeby.infura.io/v3/28962f91cd514073b78bb3736bd96e1b",
      apiKey: "17Y4MCAKUPWASA5ZA7STPZ8JKB4HHEWB7T",
      api: "https://api-rinkeby.etherscan.io/api",
      usdt: "0xD9BA894E0097f8cC2BBc9D24D308b98e36dc6D02",
    },
    minAmount: 1000000,
    contractPath: path.join(appInfo.baseDir, "contracts/wallet.json"),
    gasPrice: "10", //gwei
    pancakeRouter: "0x9ac64cc6e4415144c455bd8e4837fea55603e5c3",
  };
};

const fs = require("fs");
const { ethers } = require("ethers");
const Token = require("./contracts/wallet.json");

//总号
const privateKey = "";

const data = fs.readFileSync("wallet.json");
const list = JSON.parse(data.toString());
// console.log("list", list);

const USDT = "0x55d398326f99059ff775485246999027b3197955";

const getRpcUrl = (chainId) => {
  const rpc_urls = {
    97: [
      "https://data-seed-prebsc-1-s1.binance.org:8545",
      "https://data-seed-prebsc-2-s1.binance.org:8545",
      "https://data-seed-prebsc-1-s3.binance.org:8545",
    ],
    56: [
      "https://bsc-dataseed1.ninicoin.io",
      "https://bsc-dataseed1.defibit.io",
      "https://bsc-dataseed.binance.org",
    ],
  };
  const rpc_url =
    rpc_urls[chainId][Math.floor(Math.random() * rpc_urls[chainId].length)];
  console.log("rpc_url", rpc_url);
  // return "http://127.0.0.1:8545";
  return rpc_url;
};

const provider = new ethers.providers.JsonRpcProvider(getRpcUrl(56));
const wallet = new ethers.Wallet(privateKey, provider);

const run = async () => {
  for (var i = 0; i < list.length; i++) {
    // await sleep(4000);
    console.log("bnb", i, list[i].address);
    await sendBNB(list[i].address);
    // await sleep(8000);
    // console.log("usdt", i);
    // await sendUSDT(list[i].address);
    await sleep(3000);
  }
};

const sendBNB = async (to) => {
  let transaction = {
    to,
    value: ethers.utils.parseEther("0.01"),
  };
  let tx = await wallet.sendTransaction(transaction);
  await tx.wait();
  console.log("bnb hash", tx.hash);
};

let tokenInst = new ethers.Contract(USDT, Token.token, wallet);
const sendUSDT = async (to) => {
  const amount = ethers.utils.parseEther("5");
  const tx = await tokenInst.transfer(to, amount);
  await tx.wait();
  console.log("tx", tx.hash);
};

const sleep = (ms) => {
  return new Promise((resolve) => setTimeout(resolve, ms));
};

run();

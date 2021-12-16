const { contract, accounts } = require("@openzeppelin/test-environment");
const { constants } = require("@openzeppelin/test-helpers");
const ERC20Contract = contract.fromArtifact("ERC20FixedSupply");
const ERC20 = require("../inc/ERC20");
//固定总量代币
const totalSupply = "10000"; //发行总量
[owner, sender, receiver, purchaser, beneficiary] = accounts;
EthValue = "10";

describe("固定总量代币", function () {
  it("布署代币合约", async function () {
    ERC20Param = [
      "My Golden Coin", //代币名称
      "MGC", //代币缩写
      totalSupply, //发行总量
    ];
    ERC20Instance = await ERC20Contract.new(...ERC20Param, { from: owner });
  });
});

describe("测试ERC20合约基本信息", function () {
  ERC20.detail();
});
describe("测试ERC20合约的标准方法", async function () {
  //测试余额
  ERC20.balanceOf(totalSupply, owner, "创建者账户余额");
  //测试发送
  ERC20.transfer(
    owner,
    constants.ZERO_ADDRESS,
    EthValue,
    "代币发送,0地址错误",
    true,
    /ERC20: transfer to the zero address/
  );
  ERC20.transfer(owner, receiver, EthValue, "代币发送");
  //测试超额发送
  ERC20.transfer(
    owner,
    receiver,
    totalSupply,
    "超额发送错误",
    true,
    /ERC20: transfer amount exceeds balance/
  );
  //测试余额
  ERC20.balanceOf(EthValue, receiver, "接收者账户余额"); //receiver.balance = value
  //测试批准
  ERC20.approve(
    owner,
    constants.ZERO_ADDRESS,
    EthValue,
    "批准代币,0地址错误",
    true,
    /ERC20: approve to the zero address/
  );
  ERC20.approve(receiver, purchaser, EthValue, "批准代币");
  //验证批准
  ERC20.allowance(receiver, purchaser, EthValue, "验证批准数额"); //receiver=>purchaser = value
  //测试传送批准
  ERC20.transferFrom(receiver, purchaser, beneficiary, EthValue, "批准发送"); //beneficiary.balance = value
  //测试余额
  ERC20.balanceOf(EthValue, beneficiary, "接收者账户余额"); //receiver.balance = value
  //测试超额发送批准
  ERC20.transferFrom(
    receiver,
    purchaser,
    beneficiary,
    EthValue,
    "超额批准发送",
    true,
    /ERC20: transfer amount exceeds balance/
  );
  //验证批准归零
  ERC20.allowance(receiver, purchaser, "0", "批准额归零"); //receiver=>purchaser = 0
  //增加批准
  ERC20.increaseAllowance(receiver, purchaser, EthValue, "增加批准额");
  //验证批准
  ERC20.allowance(receiver, purchaser, EthValue, "验证批准数额"); //receiver=>purchaser = value
  //减少批准
  ERC20.decreaseAllowance(receiver, purchaser, EthValue, "减少批准额");
  //验证批准
  ERC20.allowance(receiver, purchaser, "0", "批准数额归零"); //receiver=>purchaser = 0
  //超额减少批准
  ERC20.decreaseAllowance(
    receiver,
    purchaser,
    EthValue,
    "超额减少批准额",
    true,
    /ERC20: decreased allowance below zero/
  );
});

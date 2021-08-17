package utils

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

type Eth struct {
	Client *rpc.Client
}

func NewEth() *Eth {
	c, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("rpc dial err:%v", err)
	}
	return &Eth{Client: c}
}

//Accounts 获取账户地址列表
func (e *Eth) Accounts() (accounts []string, err error) {
	err = e.Client.Call(&accounts, "eth_accounts")
	if err != nil {
		return nil, err
	}
	return
}

//GetBalance 获取指定账户余额
func (e *Eth) GetBalance(account string) (balance string, err error) {
	err = e.Client.Call(&balance, "eth_getBalance", account, "latest")
	if err != nil {
		return "", err
	}
	return
}

//GasPrice 获取gas的价格
func (e *Eth) GasPrice() (gasPrice string, err error) {
	err = e.Client.Call(&gasPrice, "eth_gasPrice")
	if err != nil {
		return "", err
	}
	return
}

//CoinBase 获取挖矿账户地址
func (e *Eth) CoinBase() (coinbase string, err error) {
	err = e.Client.Call(&coinbase, "eth_coinbase")
	if err != nil {
		return "", err
	}
	return
}

//Mining 是否在挖矿中
func (e *Eth) Mining() (isMining bool, err error) {
	err = e.Client.Call(&isMining, "eth_mining")
	if err != nil {
		return false, err
	}
	return
}

//HashRate 获取每秒计算的哈希数量
func (e *Eth) HashRate() (hashRate string, err error) {
	err = e.Client.Call(&hashRate, "eth_hashrate")
	if err != nil {
		return "", err
	}
	return
}

//GetTransactionCount 返回指定地址的交易数量
func (e *Eth) GetTransactionCount(addr string) (count string, err error) {
	err = e.Client.Call(&count, "eth_getTransactionCount", addr, "latest")
	if err != nil {
		return "", err
	}
	return
}

//BlockNumber 获取当前节点的块号
func (e *Eth) BlockNumber() (blockNum string, err error) {
	err = e.Client.Call(&blockNum, "eth_blockNumber")
	if err != nil {
		return "", err
	}
	return
}
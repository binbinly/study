package utils

import (
	"testing"
)

func TestNewEth(t *testing.T) {
	eth := NewEth()
	accounts, err := eth.Accounts()
	if err != nil {
		t.Errorf("accounts err :%v", err)
	}
	t.Logf("accounts:%v", accounts)
	balance, err := eth.GetBalance(accounts[0])
	if err != nil {
		t.Errorf("getBanalce err:%v", err)
	}
	t.Logf("getBalance:%v", balance)
	gasPrice, err := eth.GasPrice()
	if err != nil {
		t.Errorf("gasPrice err:%v", err)
	}
	t.Logf("gasPrice:%v", gasPrice)
	coinBase, err := eth.CoinBase()
	if err != nil {
		t.Errorf("coinBase err:%v", err)
	}
	t.Logf("cpinBase:%v", coinBase)
	isMining, err := eth.Mining()
	if err != nil {
		t.Errorf("Mining err:%v", err)
	}
	t.Logf("Mining:%v", isMining)
	hashRate, err := eth.HashRate()
	if err != nil {
		t.Errorf("hashRate err:%v", err)
	}
	t.Logf("hashRate:%v", hashRate)
	count, err := eth.GetTransactionCount(accounts[0])
	if err != nil {
		t.Errorf("get transaction count err:%v", err)
	}
	t.Logf("get transaction count :%v", count)
	blockNum, err := eth.BlockNumber()
	if err != nil {
		t.Errorf("blockNUmber err:%v", err)
	}
	t.Logf("block number:%v", blockNum)
}

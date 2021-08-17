package utils

import (
	"io/ioutil"
	"math/big"
	"testing"
)

var (
	conn *ContractConnect
	keystorePath1 = "/Users/q/study/geth/block/keystore/UTC--2021-07-23T15-55-09.439760000Z--3d3311ac1746c8d600d3835975fdbc822fd78c13"
)

func TestMain(m *testing.M) {
	conn = NewContractConnect()
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestQuery(t *testing.T) {
	number, err := conn.BlockNumber()
	if err != nil {
		t.Errorf("blockNumber err:%v", err)
	}
	t.Logf("number:%v", number)
	owner, err := conn.GetOwner()
	if err != nil {
		t.Errorf("GetOwner err:%v", err)
	}
	t.Logf("owner addr:%v", owner)
	period, err := conn.GetPeriod()
	if err != nil {
		t.Errorf("GetPeriod err:%v", err)
	}
	t.Logf("current period:%v", period.Int64())
	balance, err := conn.GetContractBalance()
	if err != nil {
		t.Errorf("GetContractBalance err:%v", err)
	}
	t.Logf("current contract balance:%v", balance)
}

func TestBet(t *testing.T) {
	keyData, err := ioutil.ReadFile(keystorePath1)
	if err != nil {
		t.Fatalf("read file err:%v with path:%v", err, keystorePath1)
	}
	auth := AuthAccount(keyData, "123456")
	auth.Value = big.NewInt(1e18)
	t.Logf("value:%v", auth.Value.Int64())
	trans, err := conn.Bet(auth, "1,2,3,4")
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	t.Logf("success:%v", trans.Hash())
}

func TestBetAccount(t *testing.T) {
	addr, nums, err := conn.GetBetAccounts(big.NewInt(1))
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	t.Logf("account addr:%v", addr)
	t.Logf("account bet nums:%v", nums)
}

func TestDraw(t *testing.T) {
	keyData, err := ioutil.ReadFile(keystorePath1)
	if err != nil {
		t.Fatalf("read file err:%v with path:%v", err, keystorePath1)
	}
	auth := AuthAccount(keyData, "123456")
	err = conn.Draw(auth)
	if err != nil {
		t.Fatalf("draw err:%v", err)
	}
	t.Logf("success")
}

func TestNums(t *testing.T)  {
	nums, err := conn.Nums(big.NewInt(1))
	if err != nil {
		t.Fatalf("nums err:%v", err)
	}
	t.Logf("draw nums:%v", nums)
}
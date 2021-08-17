package utils

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

type Personal struct {
	Client *rpc.Client
}

func NewPersonal() *Personal {
	c, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("rpc dial err:%v", err)
	}
	return &Personal{Client: c}
}

//Accounts 获取该节点下的所有账户地址
func (p *Personal) Accounts() (accounts []string, err error) {
	err = p.Client.Call(&accounts, "personal_listAccounts")
	if err != nil {
		return nil, err
	}
	return
}

//NewAccount 创建一个账户
func (p *Personal) NewAccount(pwd string) (account string, err error) {
	err = p.Client.Call(&account, "personal_newAccount", pwd)
	if err != nil {
		return "", err
	}
	return
}

//LockAccount 锁定用户
func (p *Personal) LockAccount(addr string) (isLock bool, err error) {
	err = p.Client.Call(&isLock, "personal_lockAccount", addr)
	if err != nil {
		return false, err
	}
	return
}

//UnlockAccount 解锁账户
func (p *Personal) UnlockAccount(addr, pwd string) (isUnlock bool, err error) {
	err = p.Client.Call(&isUnlock, "personal_unlockAccount", addr, pwd)
	if err != nil {
		return false, err
	}
	return
}

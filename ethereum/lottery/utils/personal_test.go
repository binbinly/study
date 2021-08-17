package utils

import (
	"testing"
)

func TestPersonal_Accounts(t *testing.T) {
	personal := NewPersonal()
	accounts, err := personal.Accounts()
	if err != nil {
		t.Errorf("accounts err:%v", err)
	}
	t.Logf("accounts:%v", accounts)
	account, err := personal.NewAccount("123456")
	if err != nil {
		t.Errorf("create account err:%v", err)
	}
	t.Logf("create account :%v", account)
	isLock, err := personal.LockAccount(account)
	if err != nil {
		t.Errorf("lock account err:%v", err)
	}
	t.Logf("lock account:%v", isLock)
	isUnlock, err := personal.UnlockAccount(account, "123456")
	if err != nil {
		t.Errorf("unlock account err:%v", err)
	}
	t.Logf("unlock account:%v", isUnlock)
}

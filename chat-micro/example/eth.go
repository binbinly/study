package main

import (
	"chat-micro/example/eth"
	"fmt"
)

const (
	fromAddress = ""
	toAddress   = ""
)

func main() {
	contract := eth.NewContract()
	if err := contract.Balance(fromAddress); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	if err := contract.TransferERC20FromInternalAccount(toAddress); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	select {}
}

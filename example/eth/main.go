package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"lib/eth/eth"
	"log"
	"math/big"
	"time"
)

func main() {
	amount := new(big.Int)
	amount.SetString("100000000000000000000", 10)
	sign, err := signPacket(amount)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	fmt.Println("sign", sign)
}

//signPacket 签名包
func signPacket(amount *big.Int) (string, error) {
	//contract := eth.NewContract("http://127.0.0.1:8545")
	contract := eth.NewContract("https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	err := contract.InitToken("")
	if err != nil {
		return "", err
	}
	err = contract.InitPayment("")
	if err != nil {
		return "", err
	}
	val := new(big.Int)
	val.SetString("10000000000000000000", 10)
	exist, err := contract.CheckPay("12345", val)
	if err != nil {
		return "", err
	}
	fmt.Println("exist", exist)

	exchange := ""
	account := ""
	nonce, err := contract.GetNonce(account)
	if err != nil {
		return "", err
	}

	uint256Ty, _ := abi.NewType("uint256", "", nil)
	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
	}
	deadline := time.Now().Add(time.Hour).Unix()
	bytes, _ := arguments.Pack(
		amount,
		nonce,
		big.NewInt(deadline),
	)
	fmt.Println("params", amount, nonce, deadline)
	prefix := crypto.Keccak256Hash([]byte("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"))

	owner := common.HexToAddress(account)
	spender := common.HexToAddress(exchange)
	hash := crypto.Keccak256Hash(prefix.Bytes(), owner.Bytes(), spender.Bytes(), bytes)
	fmt.Println(bytes)
	fmt.Println(hash.Hex())

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return hexutil.Encode(signature), nil
}

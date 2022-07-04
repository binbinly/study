package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main()  {
	signature := hexutil.MustDecode("")
	//if err != nil {
	//	fmt.Printf("decode err: %v\n", err)
	//}
	signature[64] -= 27
	//hash := crypto.Keccak256Hash([]byte("Hello world"))
	pubKey, err := crypto.SigToPub(signHash([]byte("Hello world")), signature)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("addr: %v\n", recoveredAddr)
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func main()  {
	token := common.HexToAddress("0x4834e69dFa970B8789cBb887581CbE0C8e6C364e")
	nft := common.HexToAddress("0x37969266859eD7Cd7f9926A0222f5265209370D7")
	nftId := new(big.Int)
	nftId.SetString("80879840152102361037665888505110546932536893549875880091605893468399", 10)
	buy1(nft, nftId, token)
}

func buy1(nft common.Address, nftId *big.Int, token common.Address) {
	c, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
	if err != nil {
		log.Fatalf("ethereum connect err:%v", err)
	}

	privateKey, err := crypto.HexToECDSA("7c8d6a02ee41937a26fd55a706e3e2c3db0a921679c28e39894a57ddd43c7964")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)

	//gasPrice := new(big.Int)
	//gasPrice.SetString("5100000000", 10)
	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice", gasPrice)

	tokenAddress := common.HexToAddress("0x38898D32b65f68Ef4Af9dc88Fb198f7B05eD12Da")

	transferFnSignature := []byte("buy(address,uint256,address)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	log.Println("methodId", hexutil.Encode(methodID))

	paddedNft := common.LeftPadBytes(nft.Bytes(), 32)
	paddedToken := common.LeftPadBytes(token.Bytes(), 32)
	paddedNftId := common.LeftPadBytes(nftId.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedNft...)
	data = append(data, paddedNftId...)
	data = append(data, paddedToken...)

	var gasLimit uint64 = 2100000

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
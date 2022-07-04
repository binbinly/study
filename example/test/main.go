package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	eth "lib/test/abi"
	"log"
	"math/big"
	"time"
)

var market *eth.Market
var client *ethclient.Client
const (
	contractAddress = "0x1d9c6010Fe569786aFa0970b076280d1Ae937903"
)

func main()  {
	connect()

	log.Println("connect success")
	go watch()
	log.Println("start notify")
	select {}
}

func connect()  {
	var err error
	client, err = ethclient.Dial("wss://bsc-ws-node.nariox.org:443")
	if err != nil {
		log.Fatalf("ethereum connect err:%v", err)
	}
	//初始化market合约
	market, err = eth.NewMarket(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("connect market err:%v", err)
	}
}

func watch()  {
	betCh := make(chan *eth.MarketTrade)
	sub, err := market.WatchTrade(&bind.WatchOpts{}, betCh, nil, nil, nil)
	if err != nil {
		log.Fatalf("lottery watch bet err:%v", err)
	}

	min := new(big.Int)
	min.SetString("10000000000000000000", 10)
	for {
		select {
		case b := <-betCh:
			if b.From == common.HexToAddress("0x0000000000000000000000000000000000000000") {
				res := b.Price.Cmp(min)
				if res <= 0 {
					log.Println("执行购买", b.NftId)
					buy(b.Nft, b.NftId, b.Token)
				}
				log.Println("event", b.Id, b.Price)
			}
		case err := <-sub.Err():
			connect()
			log.Println("watch bet err", err)
			time.Sleep(time.Second)
		}
	}
}

func buy(nft common.Address, nftId *big.Int, token common.Address) {
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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	gasPrice := new(big.Int)
	gasPrice.SetString("5100000000", 10)

	tokenAddress := common.HexToAddress(contractAddress)

	transferFnSignature := []byte("buy(address,uint256,address)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	log.Println("methodId", hexutil.Encode(methodID)) // 0xa9059cbb

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

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
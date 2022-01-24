package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	url             = ""
	contractAddress = ""
	priKey          = ""
)

// Contract 合约连接结构
type Contract struct {
	ctx    context.Context
	Client *ethclient.Client
	Token  *ERC20
}

//NewContract 创建合约
func NewContract() *Contract {
	c, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("contract connect err:%v", err)
	}
	conn := &Contract{
		ctx:    context.Background(),
		Client: c,
	}
	conn.Token, err = NewERC20(common.HexToAddress(contractAddress), c)
	go conn.watch()
	return conn
}

func (c *Contract) Balance(address string) error {
	b, err := c.Token.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return err
	}
	fmt.Printf("balance: %v\n", b)
	return nil
}

func (c *Contract) getERC20Data(toAddress common.Address, value int64) []byte {

	// 将智能合约的方法hash后取4位，表示需要调用该方法。
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	// 将第一个参数左侧置0填充为32位。
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	// 将第二个参数变成最小单位后，左侧置0为32位。
	amount := new(big.Int)
	amount.SetString(fmt.Sprintf("%d000000000000000000", value), 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	// 将方法与参数合并为一个数据。
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return data
}

func (c *Contract) getGasLimit(data []byte, toAddress common.Address) uint64 {
	gasLimit, err := c.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatalf("err:%v\n", err)
	}
	return gasLimit * 10
}

func (c *Contract) TransferERC20FromInternalAccount(to string) error {
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 获取当前的矿工建议费用。
	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000)
	data := c.getERC20Data(common.HexToAddress(to), 10)
	gasLimit := c.getGasLimit(data, common.HexToAddress(to)) // in units

	token := common.HexToAddress(contractAddress)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &token, // 与以太交易不同，是发送给合约地址。
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
	fmt.Println("tx", gasLimit, gasPrice, nonce)

	chainID, err := c.Client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	err = c.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	//fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
	return nil
}

//watch 监听投注开奖完成
func (c *Contract) watch() {
	transfer := make(chan *ERC20Transfer)
	sub, err := c.Token.WatchTransfer(&bind.WatchOpts{}, transfer, nil, nil)
	if err != nil {
		log.Fatalf("lottery watch bet err:%v", err)
	}

	for {
		select {
		case b := <-transfer:
			log.Printf("Lottery watch bet finish Transfer From: %s, To:%v, Value: %v, Raw: %v", b.From, b.To, b.Value, b.Raw)
		case err := <-sub.Err():
			log.Println("watch bet err", err)
		}
	}
}
package eth

import (
	"context"
	eth "lib/eth/abi"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Contract 合约连接结构
type Contract struct {
	ctx     context.Context
	Client  *ethclient.Client
	Payment *eth.Payment
	Token   *eth.Token
}

//NewContract 创建合约
func NewContract(url string) *Contract {
	c, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("ethereum connect err:%v", err)
	}
	conn := &Contract{
		ctx:    context.Background(),
		Client: c,
	}
	return conn
}

//InitPayment 初始化支付合约
func (c *Contract) InitPayment(address string) error {
	pay, err := eth.NewPayment(common.HexToAddress(address), c.Client)
	if err != nil {
		return err
	}
	c.Payment = pay
	return nil
}

//InitToken 初始化代币合约
func (c *Contract) InitToken(address string) error {
	token, err := eth.NewToken(common.HexToAddress(address), c.Client)
	if err != nil {
		return err
	}
	c.Token = token
	return nil
}

//CheckPay 验证是否已支付
func (c *Contract) CheckPay(orderNo string, amount *big.Int) (bool, error) {
	return c.Payment.Query(nil, orderNo, amount)
}

//Balance 获取余额
func (c *Contract) Balance(owner string) (*big.Int, error) {
	return c.Token.BalanceOf(nil, common.HexToAddress(owner))
}

//GetNonce 获取签名nonce
func (c *Contract) GetNonce(owner string) (*big.Int, error) {
	return c.Token.Nonces(nil, common.HexToAddress(owner))
}

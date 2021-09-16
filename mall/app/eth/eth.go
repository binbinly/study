package eth

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"mall/app/conf"
	"mall/eth"
)

// Contract 合约连接结构
type Contract struct {
	ctx      context.Context
	Client   *ethclient.Client
	Payments map[int8]*eth.Payment
}

//NewContract 创建合约
func NewContract(conf *conf.EthConfig) *Contract {
	c, err := ethclient.Dial(conf.NetworkUrl)
	if err != nil {
		log.Fatalf("contract connect err:%v", err)
	}
	conn := &Contract{
		ctx:      context.Background(),
		Client:   c,
		Payments: make(map[int8]*eth.Payment),
	}
	return conn
}

//Connect 连接合约地址
func (c *Contract) Connect(id int8, address string) error {
	if _, ok := c.Payments[id]; ok {
		return nil
	}
	pay, err := eth.NewPayment(common.HexToAddress(address), c.Client)
	if err != nil {
		return err
	}
	c.Payments[id] = pay
	return nil
}

//CheckPay 验证是否已支付
func (c *Contract) CheckPay(id int8, orderNo string) (bool, error) {
	if pay, ok := c.Payments[id]; ok {
		return pay.CheckPay(nil, orderNo)
	}
	return false, nil
}

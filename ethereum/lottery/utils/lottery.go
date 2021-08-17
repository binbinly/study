package utils

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"lottery/eth"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//合约地址
	ContractAddr = common.HexToAddress("0x929e349826875B4700B37979b4A4e714d161C27a")
	//链ID
	ChainID = big.NewInt(15)
)

// ContractConnect 合约连接结构
type ContractConnect struct {
	ctx     context.Context
	Client  *ethclient.Client
	Lottery *eth.Lottery
}

func NewContractConnect() *ContractConnect {
	c, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatalf("contract connect err:%v", err)
	}
	lottery, err := eth.NewLottery(ContractAddr, c)
	if err != nil {
		log.Fatalf("new lottery err:%v", err)
	}
	conn := &ContractConnect{
		ctx:     context.Background(),
		Client:  c,
		Lottery: lottery,
	}
	//开启时间监听必须使用 ws 连接
	go conn.watch()
	return conn
}

// BlockNumber 当前块高度
func (c *ContractConnect) BlockNumber() (*big.Int, error) {
	blockNumber, err := c.Client.BlockByNumber(c.ctx, nil)
	if err != nil {
		return nil, err
	}
	return blockNumber.Number(), nil
}

// GetOwner 合约管理者
func (c *ContractConnect) GetOwner() (string, error) {
	owner, err := c.Lottery.Owner(nil)
	if err != nil {
		return "", err
	}
	return owner.String(), nil
}

// GetPeriod 获取当前期数
func (c *ContractConnect) GetPeriod() (*big.Int, error) {
	period, err := c.Lottery.Period(nil)
	if err != nil {
		return nil, err
	}
	return period, nil
}

// Bet 投注
func (c *ContractConnect) Bet(fromAuth *bind.TransactOpts, nums string) (*types.Transaction, error) {
	trans, err := c.Lottery.Bet(fromAuth, nums)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

// GetBetAccounts 获取投注账号
func (c *ContractConnect) GetBetAccounts(index *big.Int) (string, string, error) {
	res, err := c.Lottery.AccountList(nil, index)
	if err != nil {
		return "", "", err
	}
	return res.Addr.String(), res.Nums, nil
}

//GetBalance 获取账户余额
func (c *ContractConnect) GetBalance(addr string) (int64, error) {
	balance, err := c.Lottery.GetBalance(nil, common.HexToAddress(addr))
	if err != nil {
		return 0, err
	}
	return balance.Int64(), nil
}

//GetContractBalance 获取合约余额
func (c *ContractConnect) GetContractBalance() (int64, error) {
	balance, err := c.Lottery.GetBalance(nil, ContractAddr)
	if err != nil {
		return 0, err
	}
	return balance.Int64(), nil
}

// Draw 开奖
func (c *ContractConnect) Draw(fromAuth *bind.TransactOpts) error {
	_, err := c.Lottery.Draw(fromAuth)
	if err != nil {
		return err
	}
	return nil
}

// Nums 某期开奖号码
func (c *ContractConnect) Nums(period *big.Int) ([]int64, error) {
	var nums []int64
	for i := 0; i < 5; i++ {
		num, err := c.Lottery.RetNums(nil, period, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}
		nums = append(nums, num.Int64())
	}
	return nums, nil
}

//watch 监听投注开奖完成
func (c *ContractConnect) watch() {
	betCh := make(chan *eth.LotteryBetFinish)
	betSub, err := c.Lottery.WatchBetFinish(&bind.WatchOpts{}, betCh)
	if err != nil {
		log.Fatalf("lottery watch bet err:%v", err)
	}

	drawCh := make(chan *eth.LotteryDrawFinish)
	drawSub, err := c.Lottery.WatchDrawFinish(&bind.WatchOpts{}, drawCh)
	if err != nil {
		log.Fatalf("lottery watch draw err:%v", err)
	}

	for {
		select {
		case b := <-betCh:
			log.Printf("Lottery watch bet finish Transfer From: %s, Nums:%v", b.Addr, b.Nums)
		case err := <-betSub.Err():
			log.Println("watch bet err", err)
		case d := <-drawCh:
			log.Printf("Lottery watch draw finish Nums:%v", d.Nums)
		case err := <-drawSub.Err():
			log.Println("watch draw err", err)
		}
	}
}

// AuthAccount 解锁账户
// 正式使用时候，此处应该限制GasPrice和GasLimit
func AuthAccount(keyData []byte, passphrase string) *bind.TransactOpts {
	auth, err := bind.NewTransactorWithChainID(bytes.NewReader(keyData), passphrase, ChainID)
	if err != nil {
		log.Fatalf("new transactor err:%v", err)
	}
	auth.GasLimit = uint64(210000)
	auth.GasPrice = big.NewInt(20000)
	return auth
}

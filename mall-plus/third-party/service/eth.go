package service

import (
	"context"

	"github.com/pkg/errors"

	"common/errno"
)

// CheckPay 检查订单是否已支付
func (s *Service) CheckPay(ctx context.Context, id int64, address, orderNo string) error {
	//连接合约
	err := s.contract.Connect(id, address)
	if err != nil {
		return errors.Wrapf(err, "[third.eth] connect contract address %v", address)
	}
	//调用合约
	check, err := s.contract.CheckPay(id, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[third.eth] contract call checkPay")
	}
	if !check { //未支付
		return errno.ErrETHPayNotFound
	}

	return nil
}

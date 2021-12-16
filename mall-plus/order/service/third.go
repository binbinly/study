package service

import (
	"context"

	pb "common/proto/third"
)

//checkEthPay 检查以太币是否已支付
func (s *Service) checkEthPay(ctx context.Context, id int64, address, orderNo string) error {
	_, err := s.thirdService.CheckETHPay(ctx, &pb.ETHPayReq{
		Id:      id,
		Address: address,
		OrderNo: orderNo,
	})
	if err != nil {
		return err
	}
	return nil
}
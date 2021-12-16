package service

import (
	"context"

	pb "common/proto/member"
)

//getAddressInfo 获取收货地址详情
func (s *Service) getAddressInfo(ctx context.Context, id, memberID int64) (*pb.Address, error) {
	addr, err := s.memberService.GetAddressInfo(ctx, &pb.AddressInfoReq{
		UserId:    memberID,
		AddressId: id,
	})
	if err != nil {
		return nil, err
	}
	return addr.Info, nil
}
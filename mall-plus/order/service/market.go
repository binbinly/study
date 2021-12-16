package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "common/proto/market"
)

//getCouponInfo 获取优惠券详情
func (s *Service) getCouponInfo(ctx context.Context, id, memberID int64) (*pb.Coupon, error) {
	coupon, err := s.marketService.GetCouponInfo(ctx, &pb.CouponInfoReq{
		UserId:   memberID,
		CouponId: id,
	})
	if err != nil {
		return nil, err
	}
	return coupon.Info, nil
}

//setCouponUsed 设置优惠券已使用
func (s *Service) setCouponUsed(ctx context.Context, id, memberID, orderID int64) error {
	_, err := s.marketService.CouponUsed(ctx, &pb.CouponUsedReq{
		UserId:   memberID,
		CouponId: id,
		OrderId:  orderID,
	})
	if err != nil {
		return err
	}
	return nil
}

//getPayConfig 获取支付方式配置
func (s *Service) getPayConfig(ctx context.Context) ([]*pb.PayItem, error) {
	reply, err := s.marketService.GetPayConfig(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return reply.Data, nil
}
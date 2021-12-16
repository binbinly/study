package service

import (
	"context"

	pb "common/proto/cart"
)

//batchGetCarts 批量获取购物车信息
func (s *Service) batchGetCarts(ctx context.Context, memberID int64, skuIds []int64) ([]*pb.CartItem, error) {
	reply, err := s.cartService.BatchGetCarts(ctx, &pb.SkusReq{
		UserId: memberID,
		SkuIds: skuIds,
	})
	if err != nil {
		return nil, err
	}
	return reply.Data, nil
}

//DelCart 删除购物车
func (s *Service) delCart(ctx context.Context, memberID int64, skuIds []int64) error {
	_, err := s.cartService.BatchDelCart(ctx, &pb.SkusReq{
		UserId: memberID,
		SkuIds: skuIds,
	})
	if err != nil {
		return err
	}
	return nil
}

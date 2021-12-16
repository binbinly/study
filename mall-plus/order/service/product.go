package service

import (
	"context"

	pb "common/proto/product"
)

//产品服务调用
//getSkuByID 获取sku信息
func (s *Service) getSkuByID(ctx context.Context, skuID int64) (*pb.SkuInfo, error) {
	reply, err := s.productService.GetSkuByID(ctx, &pb.SkuReq{
		SkuId: skuID,
	})
	if err != nil {
		return nil, err
	}
	return reply.Info, nil
}

//spuComment 商品评论
func (s *Service) spuComment(ctx context.Context, skuIds []int64, userID, orderID int64, star int32, content, resources string) error {
	_, err := s.productService.SpuComment(ctx, &pb.CommentReq{
		SkuIds:    skuIds,
		UserId:    userID,
		OrderId:   orderID,
		Star:      star,
		Content:   content,
		Resources: resources,
	})
	if err != nil {
		return err
	}
	return nil
}
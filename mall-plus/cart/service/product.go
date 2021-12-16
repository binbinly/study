package service

import (
	"context"

	pb "common/proto/product"
)

//产品服务调用

//GetSkuByID 获取sku信息
func (s *Service) GetSkuByID(ctx context.Context, skuID int64) (*pb.SkuInfo, error) {
	reply, err := s.productService.GetSkuByID(ctx, &pb.SkuReq{
		SkuId: skuID,
	})
	if err != nil {
		return nil, err
	}
	return reply.Info, nil
}

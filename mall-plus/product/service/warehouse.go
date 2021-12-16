package service

import (
	"context"

	pb "common/proto/warehouse"
	"product/model"
)

//getSpuStock 获取spu下所有sku的库存
func (s *Service) getSpuStock(ctx context.Context, spuID int64, skus []*model.SkuModel) (map[int64]int32, error) {
	skuIds := make([]int64, 0, len(skus))
	for _, sku := range skus {
		skuIds = append(skuIds, sku.ID)
	}
	stocks, err := s.wareService.GetSpuStock(ctx, &pb.SpuStockReq{
		SpuId:  spuID,
		SkuIds: skuIds,
	})
	if err != nil {
		return nil, err
	}
	return stocks.SkuNum, nil
}

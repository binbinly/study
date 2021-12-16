package service

import (
	"context"
	"strings"

	pb "common/proto/warehouse"

	"order/model"
)

//batchLockSkuStock 批量锁定库存
func (s *Service) batchLockSkuStock(ctx context.Context, order *model.OrderModel, skus map[int64]int32) error {
	_, err := s.wareService.SKuStockLock(ctx, &pb.SkuStockLockReq{
		OrderId:   order.ID,
		OrderNo:   order.OrderNo,
		Consignee: order.AddressName,
		Phone:     order.AddressPhone,
		Address:   strings.Join([]string{order.AddressProvince, order.AddressCity, order.AddressCounty, order.AddressDetail}, " "),
		Note:      order.Note,
		SkuNum:    skus,
	})
	if err != nil {
		return err
	}
	return nil
}
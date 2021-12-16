package service

import (
	"common/orm"
	"common/util"
	"context"
	"warehouse/conf"
	"warehouse/repo"
)

var _ IService = (*Service)(nil)

//IService 营销服务接口
type IService interface {
	GetSkuStock(ctx context.Context, skuID int64) (int, error)
	GetSpuStock(ctx context.Context, spuID int64, skuIds []int64) (map[int64]int32, error)
	SKuStockLock(ctx context.Context, orderID int64, orderNo, consignee, phone, address, note string,
		sku map[int64]int32) error
	SkuStockUnlock(ctx context.Context, orderID int64, finish bool) error

	Close() error
}

//Service 营销服务
type Service struct {
	c    *conf.Config
	repo repo.IRepo
}

// New init service
func New(c *conf.Config) IService {
	return &Service{
		c:    c,
		repo: repo.New(orm.GetDB(), util.NewCache()),
	}
}

// Close service
func (s *Service) Close() error {
	return s.repo.Close()
}

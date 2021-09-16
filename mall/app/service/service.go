package service

import (
	"context"
	"mall/app/conf"
	"mall/app/eth"
	"mall/app/model"
	"mall/app/repo"
)

//IService 服务接口
type IService interface {
	IUser
	IGoods
	ICart
	ICoupon
	IUserAddress
	IOrder

	PayList(ctx context.Context) ([]*model.ConfigPayList, error)
	AreaList(ctx context.Context) (map[string]interface{}, error)
	CategoryTree(ctx context.Context) ([]*model.GoodsCategoryTree, error)
	HomeData(ctx context.Context) ([]*model.ConfigHomeCat, error)
	HomeCatData(ctx context.Context, cid int) ([]*model.AppSetting, error)
	NoticeList(ctx context.Context, offset, limit int) ([]*model.AppNoticeModel, error)
	SearchHotWord(ctx context.Context) (map[string]interface{}, error)

	Close() error
}

var Svc IService

// Service struct
type Service struct {
	c        *conf.Config
	repo     repo.IRepo
	contract *eth.Contract
}

// New init service
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:        c,
		repo:     repo.New(model.GetDB()),
		contract: eth.NewContract(&c.Eth),
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	return s.repo.Close()
}

package service

import (
	"common/orm"
	"common/util"
	"market/conf"
	"market/repo"
)

var _ IService = (*Service)(nil)

//IService 营销服务接口
type IService interface {
	IPage
	ICoupon

	Close() error
}

//Service 营销服务
type Service struct {
	c     *conf.Config
	repo  repo.IRepo
}

// New init service
func New(c *conf.Config) IService {
	return &Service{
		c:     c,
		repo:  repo.New(orm.GetDB(), util.NewCache()),
	}
}

// Close service
func (s *Service) Close() error {
	return s.repo.Close()
}

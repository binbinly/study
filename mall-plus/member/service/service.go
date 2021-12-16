package service

import (
	"go-micro.dev/v4/client"

	"common/constvar"
	"common/orm"
	center "common/proto/center"
	third "common/proto/third"
	"common/util"
	"member/conf"
	"member/repo"
)

var _ IService = (*Service)(nil)

//IService 营销服务接口
type IService interface {
	ICenter
	IThird
	IMember
	IMemberAddress

	Close() error
}

//Service 营销服务
type Service struct {
	c             *conf.Config
	repo          repo.IRepo
	centerService center.UserService
	thirdService  third.ThirdService
}

// New init service
func New(c *conf.Config, cl client.Client) IService {
	return &Service{
		c:             c,
		repo:          repo.New(orm.GetDB(), util.NewCache()),
		centerService: center.NewUserService(constvar.ServiceCenter, cl),
		thirdService:  third.NewThirdService(constvar.ServiceThird, cl),
	}
}

// Close service
func (s *Service) Close() error {
	return s.repo.Close()
}

package service

import (
	"go-micro.dev/v4"

	"center/conf"
	"center/repo"
	"common/constvar"
	"common/orm"
	"common/util"
)

var _ ICenter = (*Center)(nil)

//ICenter 中心服接口定义
type ICenter interface {
	IUser
	IOnline
	IPush

	Close() error
}

// Center 中心服结构
type Center struct {
	c     *conf.Config
	repo  repo.IRepo
	event micro.Event
}

// New init service
func New(c *conf.Config) ICenter {
	return &Center{
		c:     c,
		repo:  repo.New(orm.GetDB(), util.NewCache()),
		event: micro.NewEvent(constvar.TopicTask, nil),
	}
}

// Close service
func (c *Center) Close() error {
	return c.repo.Close()
}

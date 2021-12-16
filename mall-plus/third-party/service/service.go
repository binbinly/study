package service

import (
	"context"

	"third-party/conf"
	"third-party/eth"
)

var _ IService = (*Service)(nil)

//IService 第三方服务接口定义
type IService interface {
	SendSMS(ctx context.Context, phone string) (string, error)
	CheckVCode(ctx context.Context, phone int64, vCode string) error
	CheckPay(ctx context.Context, id int64, address, orderNo string) error

	Close() error
}

//Service 第三方服务
type Service struct {
	c        *conf.Config
	contract *eth.Contract
}

// New init service
func New(c *conf.Config) IService {
	return &Service{
		c:        c,
		contract: eth.NewContract(&c.Eth),
	}
}

// Close service
func (t *Service) Close() error {
	return nil
}

package service

import (
	"context"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"

	"common/constvar"
	pb "common/proto/seckill"
	"pkg/redis"
	"seckill/conf"
	"seckill/repo"
)

var _ IService = (*Service)(nil)

//IService 营销服务接口
type IService interface {
	Seckill(ctx context.Context, memberID, skuID, addressID, num int64, key string) (string, error)
	GetSessionAll(ctx context.Context) ([]*pb.Session, error)
	GetSessionSkus(ctx context.Context, sessionID int64) ([]*pb.Sku, error)
	GetSkuInfo(ctx context.Context, skuID int64) (*pb.Sku, error)

	Close() error
}

//Service 营销服务
type Service struct {
	c          *conf.Config
	repo       repo.IRepo
	orderEvent micro.Event
}

// New init service
func New(c *conf.Config, cl client.Client) IService {
	return &Service{
		c:          c,
		repo:       repo.New(redis.Client),
		orderEvent: micro.NewEvent(constvar.TopicOrderSeckill, cl),
	}
}

// Close service
func (s *Service) Close() error {
	return nil
}

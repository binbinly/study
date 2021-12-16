package service

import (
	"context"

	"go-micro.dev/v4/client"

	"cart/conf"
	"cart/repo"
	"common/constvar"
	pb "common/proto/cart"
	product "common/proto/product"
	"pkg/redis"
)

var _ IService = (*Service)(nil)

//IService 营销服务接口
type IService interface {
	AddCart(ctx context.Context, userID, skuID int64, num int) error
	EditCart(ctx context.Context, userID, oldSkuID, newSkuID int64, num int) error
	EditCartNum(ctx context.Context, userID, skuID int64, num int) error
	DelCart(ctx context.Context, userID int64, skuIds []int64) error
	ClearCart(ctx context.Context, userID int64) error
	CartList(ctx context.Context, userID int64) ([]*pb.CartItem, error)
	BatchGetCarts(ctx context.Context, userID int64, skuIds []int64) ([]*pb.CartItem, error)

	Close() error
}

//Service 营销服务
type Service struct {
	c              *conf.Config
	repo           repo.IRepo
	productService product.ProductService // 产品服务
}

// New init service
func New(c *conf.Config, cl client.Client) IService {
	return &Service{
		c:              c,
		repo:           repo.New(redis.Client),
		productService: product.NewProductService(constvar.ServiceProduct, cl),
	}
}

// Close service
func (s *Service) Close() error {
	return nil
}

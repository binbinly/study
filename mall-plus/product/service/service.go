package service

import (
	"context"
	"log"

	"go-micro.dev/v4/client"

	"common/constvar"
	"common/orm"
	pb "common/proto/product"
	kill "common/proto/seckill"
	ware "common/proto/warehouse"
	"common/util"
	"pkg/database/elasticsearch"
	"product/conf"
	"product/es"
	"product/repo"
)

var _ IService = (*Service)(nil)

//IService 营销服务接口
type IService interface {
	CategoryTree(ctx context.Context) ([]*pb.Category, error)
	SkuList(ctx context.Context, catID int64, offset, limit int) ([]*pb.SkuEs, error)
	SkuDetail(ctx context.Context, id int64) (*pb.Sku, error)
	GetSkuSaleAttrs(ctx context.Context, id int64) (*pb.SkuSaleAttr, error)
	GetSkuInfo(ctx context.Context, id int64) (*pb.SkuInfo, error)
	SpuComment(ctx context.Context, skuIds []int64, memberID, orderID int64, star int8, content, resources string) error
	Search(ctx context.Context, keyword string, catID int64, field, order, priceS, priceE int32,
		hasStock bool, brandIds []int64, attrs map[int64][]string, page int32) (*pb.SearchReply, error)

	Close() error
}

//Service 营销服务
type Service struct {
	c           *conf.Config
	repo        repo.IRepo
	productEs   *es.Product
	wareService ware.WarehouseService
	killService kill.SeckillService
}

// New init service
func New(c *conf.Config, cl client.Client) IService {
	e := es.New(elasticsearch.NewClient(&c.Elastic))
	err := e.Init(context.Background())
	if err != nil {
		log.Fatalf("es init err: %v", err)
	}
	return &Service{
		c:           c,
		repo:        repo.New(orm.GetDB(), util.NewCache()),
		productEs:   e,
		wareService: ware.NewWarehouseService(constvar.ServiceWarehouse, cl),
		killService: kill.NewSeckillService(constvar.ServiceSeckill, cl),
	}
}

// Close service
func (s *Service) Close() error {
	return s.repo.Close()
}

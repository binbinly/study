package service

import (
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"

	"common/constvar"
	"common/orm"
	"common/util"
	"order/conf"
	"order/repo"
	"pkg/rabbitmq"

	cart "common/proto/cart"
	market "common/proto/market"
	member "common/proto/member"
	product "common/proto/product"
	third "common/proto/third"
	ware "common/proto/warehouse"
)

var _ IService = (*Service)(nil)

//IService 营销服务接口
type IService interface {
	IOrder

	Close() error
}

//Service 营销服务
type Service struct {
	c              *conf.Config
	repo           repo.IRepo
	event          *rabbitmq.Producer
	wareEvent      micro.Event
	cartService    cart.CartService
	memberService  member.MemberService
	marketService  market.MarketService
	wareService    ware.WarehouseService
	productService product.ProductService
	thirdService   third.ThirdService
}

// New init service
func New(c *conf.Config, cl client.Client) IService {
	e := rabbitmq.NewProducer(c.AMQP.Addr, constvar.ExchangeOrder, constvar.KeyOrderCreate)
	if err := e.Start(); err != nil {
		log.Fatal(err)
	}
	return &Service{
		c:              c,
		repo:           repo.New(orm.GetDB(), util.NewCache()),
		event:          e,
		wareEvent:      micro.NewEvent(constvar.TopicWarehouse, cl),
		cartService:    cart.NewCartService(constvar.ServiceCart, cl),
		memberService:  member.NewMemberService(constvar.ServiceMember, cl),
		marketService:  market.NewMarketService(constvar.ServiceMarket, cl),
		wareService:    ware.NewWarehouseService(constvar.ServiceWarehouse, cl),
		productService: product.NewProductService(constvar.ServiceProduct, cl),
		thirdService:   third.NewThirdService(constvar.ServiceThird, cl),
	}
}

// Close service
func (s *Service) Close() error {
	return s.repo.Close()
}

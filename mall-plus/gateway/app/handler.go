package app

import (
	"log"

	"pkg/registry/consul"

	cart "gateway/proto/cart"
	market "gateway/proto/market"
	member "gateway/proto/member"
	order "gateway/proto/order"
	product "gateway/proto/product"
	seckill "gateway/proto/seckill"
)

//RegisterAll 注册全部处理器
func (a *App) RegisterAll() {
	//初始化 consul resolver
	consul.Init(a.opts.registry)

	a.RegisterOrder()
	a.RegisterMember()
	a.RegisterMarket()
	a.RegisterCart()
	a.RegisterProduct()
	a.RegisterSeckill()
}

//RegisterMarket 注册营销服务处理器
func (a *App) RegisterMarket() {
	conn := a.newRPCClientConn(&a.opts.services.Market)
	if err := market.RegisterMarketHandler(a.ctx, a.opts.mux, conn); err != nil {
		log.Fatalf("register market handler err: %v", err)
	}
	log.Printf("register market service success")
}

//RegisterProduct 注册产品服务处理器
func (a *App) RegisterProduct() {
	conn := a.newRPCClientConn(&a.opts.services.Product)
	if err := product.RegisterProductHandler(a.ctx, a.opts.mux, conn); err != nil {
		log.Fatalf("register product handler err: %v", err)
	}
	log.Printf("register product service success")
}

//RegisterCart 注册购物车服务处理器
func (a *App) RegisterCart() {
	conn := a.newRPCClientConn(&a.opts.services.Cart)
	if err := cart.RegisterCartHandler(a.ctx, a.opts.mux, conn); err != nil {
		log.Fatalf("register cart handler err: %v", err)
	}
	log.Printf("register cart service success")
}

//RegisterMember 注册会员服务处理器
func (a *App) RegisterMember() {
	conn := a.newRPCClientConn(&a.opts.services.Member)
	if err := member.RegisterMemberHandler(a.ctx, a.opts.mux, conn); err != nil {
		log.Fatalf("register member handler err: %v", err)
	}
	log.Printf("register member service success")
}

//RegisterOrder 注册订单服务处理器
func (a *App) RegisterOrder() {
	conn := a.newRPCClientConn(&a.opts.services.Order)
	if err := order.RegisterOrderHandler(a.ctx, a.opts.mux, conn); err != nil {
		log.Fatalf("register order handler err: %v", err)
	}
	log.Printf("register order service success")
}

//RegisterSeckill 注册秒杀服务处理器
func (a *App) RegisterSeckill() {
	conn := a.newRPCClientConn(&a.opts.services.Seckill)
	if err := seckill.RegisterSeckillHandler(a.ctx, a.opts.mux, conn); err != nil {
		log.Fatalf("register seckill handler err: %v", err)
	}
	log.Printf("register seckill service success")
}

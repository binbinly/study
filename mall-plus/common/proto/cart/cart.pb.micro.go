// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cart/cart.proto

package common

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Cart service

func NewCartEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cart service

type CartService interface {
	/// 添加购物车
	AddCart(ctx context.Context, in *AddReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 更新购物车
	EditCart(ctx context.Context, in *EditReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 更新购物车数量
	EditCartNum(ctx context.Context, in *AddReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 删除购物项
	DelCart(ctx context.Context, in *SkuReq, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 清空购物车
	ClearCart(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error)
	/// 我的购物车
	MyCart(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*CartsReply, error)
	/// 批量获取购物车信息
	BatchGetCarts(ctx context.Context, in *SkusReq, opts ...client.CallOption) (*CartsReply, error)
	/// 批量删除购物车
	BatchDelCart(ctx context.Context, in *SkusReq, opts ...client.CallOption) (*emptypb.Empty, error)
}

type cartService struct {
	c    client.Client
	name string
}

func NewCartService(name string, c client.Client) CartService {
	return &cartService{
		c:    c,
		name: name,
	}
}

func (c *cartService) AddCart(ctx context.Context, in *AddReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Cart.AddCart", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) EditCart(ctx context.Context, in *EditReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Cart.EditCart", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) EditCartNum(ctx context.Context, in *AddReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Cart.EditCartNum", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) DelCart(ctx context.Context, in *SkuReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Cart.DelCart", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) ClearCart(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Cart.ClearCart", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) MyCart(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*CartsReply, error) {
	req := c.c.NewRequest(c.name, "Cart.MyCart", in)
	out := new(CartsReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) BatchGetCarts(ctx context.Context, in *SkusReq, opts ...client.CallOption) (*CartsReply, error) {
	req := c.c.NewRequest(c.name, "Cart.BatchGetCarts", in)
	out := new(CartsReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) BatchDelCart(ctx context.Context, in *SkusReq, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Cart.BatchDelCart", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cart service

type CartHandler interface {
	/// 添加购物车
	AddCart(context.Context, *AddReq, *emptypb.Empty) error
	/// 更新购物车
	EditCart(context.Context, *EditReq, *emptypb.Empty) error
	/// 更新购物车数量
	EditCartNum(context.Context, *AddReq, *emptypb.Empty) error
	/// 删除购物项
	DelCart(context.Context, *SkuReq, *emptypb.Empty) error
	/// 清空购物车
	ClearCart(context.Context, *emptypb.Empty, *emptypb.Empty) error
	/// 我的购物车
	MyCart(context.Context, *emptypb.Empty, *CartsReply) error
	/// 批量获取购物车信息
	BatchGetCarts(context.Context, *SkusReq, *CartsReply) error
	/// 批量删除购物车
	BatchDelCart(context.Context, *SkusReq, *emptypb.Empty) error
}

func RegisterCartHandler(s server.Server, hdlr CartHandler, opts ...server.HandlerOption) error {
	type cart interface {
		AddCart(ctx context.Context, in *AddReq, out *emptypb.Empty) error
		EditCart(ctx context.Context, in *EditReq, out *emptypb.Empty) error
		EditCartNum(ctx context.Context, in *AddReq, out *emptypb.Empty) error
		DelCart(ctx context.Context, in *SkuReq, out *emptypb.Empty) error
		ClearCart(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error
		MyCart(ctx context.Context, in *emptypb.Empty, out *CartsReply) error
		BatchGetCarts(ctx context.Context, in *SkusReq, out *CartsReply) error
		BatchDelCart(ctx context.Context, in *SkusReq, out *emptypb.Empty) error
	}
	type Cart struct {
		cart
	}
	h := &cartHandler{hdlr}
	return s.Handle(s.NewHandler(&Cart{h}, opts...))
}

type cartHandler struct {
	CartHandler
}

func (h *cartHandler) AddCart(ctx context.Context, in *AddReq, out *emptypb.Empty) error {
	return h.CartHandler.AddCart(ctx, in, out)
}

func (h *cartHandler) EditCart(ctx context.Context, in *EditReq, out *emptypb.Empty) error {
	return h.CartHandler.EditCart(ctx, in, out)
}

func (h *cartHandler) EditCartNum(ctx context.Context, in *AddReq, out *emptypb.Empty) error {
	return h.CartHandler.EditCartNum(ctx, in, out)
}

func (h *cartHandler) DelCart(ctx context.Context, in *SkuReq, out *emptypb.Empty) error {
	return h.CartHandler.DelCart(ctx, in, out)
}

func (h *cartHandler) ClearCart(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	return h.CartHandler.ClearCart(ctx, in, out)
}

func (h *cartHandler) MyCart(ctx context.Context, in *emptypb.Empty, out *CartsReply) error {
	return h.CartHandler.MyCart(ctx, in, out)
}

func (h *cartHandler) BatchGetCarts(ctx context.Context, in *SkusReq, out *CartsReply) error {
	return h.CartHandler.BatchGetCarts(ctx, in, out)
}

func (h *cartHandler) BatchDelCart(ctx context.Context, in *SkusReq, out *emptypb.Empty) error {
	return h.CartHandler.BatchDelCart(ctx, in, out)
}
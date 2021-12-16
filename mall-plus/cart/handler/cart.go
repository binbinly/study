package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"cart/service"
	pb "common/proto/cart"
	"common/util"
)

//Auth 购物车服身份验证
func Auth(method string) bool {
	return true
}

//Cart 购物车处理器
type Cart struct {
	srv service.IService
}

// New 实例化购物车处理器
func New(srv service.IService) *Cart {
	return &Cart{srv: srv}
}

//AddCart 添加购物车
func (c *Cart) AddCart(ctx context.Context, req *pb.AddReq, empty *emptypb.Empty) error {
	err := c.srv.AddCart(ctx, util.GetUserID(ctx), req.SkuId, int(req.Num))
	if err != nil {
		return err
	}
	return nil
}

//EditCart 更新购物车
func (c *Cart) EditCart(ctx context.Context, req *pb.EditReq, empty *emptypb.Empty) error {
	err := c.srv.EditCart(ctx, util.GetUserID(ctx), req.OldSkuId, req.NewSkuId, int(req.Num))
	if err != nil {
		return err
	}
	return nil
}

//EditCartNum 更新购物车数量
func (c *Cart) EditCartNum(ctx context.Context, req *pb.AddReq, empty *emptypb.Empty) error {
	err := c.srv.EditCartNum(ctx, util.GetUserID(ctx), req.SkuId, int(req.Num))
	if err != nil {
		return err
	}
	return nil
}

//DelCart 删除购物车
func (c *Cart) DelCart(ctx context.Context, req *pb.SkuReq, empty *emptypb.Empty) error {
	err := c.srv.DelCart(ctx, util.GetUserID(ctx), []int64{req.SkuId})
	if err != nil {
		return err
	}
	return nil
}

//ClearCart 清空购物车
func (c *Cart) ClearCart(ctx context.Context, req *emptypb.Empty, empty *emptypb.Empty) error {
	err := c.srv.ClearCart(ctx, util.GetUserID(ctx))
	if err != nil {
		return err
	}
	return nil
}

//MyCart 我的购物车
func (c *Cart) MyCart(ctx context.Context, req *emptypb.Empty, reply *pb.CartsReply) error {
	list, err := c.srv.CartList(ctx, util.GetUserID(ctx))
	if err != nil {
		return err
	}
	reply.Data = list
	return nil
}

//BatchGetCarts 批量获取购物车信息
func (c *Cart) BatchGetCarts(ctx context.Context, req *pb.SkusReq, reply *pb.CartsReply) error {
	list, err := c.srv.BatchGetCarts(ctx, req.UserId, req.SkuIds)
	if err != nil {
		return err
	}
	reply.Data = list
	return nil
}

//BatchDelCart 批量删除购物车
func (c *Cart) BatchDelCart(ctx context.Context, req *pb.SkusReq, empty *emptypb.Empty) error {
	err := c.srv.DelCart(ctx, req.UserId, req.SkuIds)
	if err != nil {
		return err
	}
	return nil
}

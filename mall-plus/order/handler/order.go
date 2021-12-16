package handler

import (
	"context"
	"reflect"

	"google.golang.org/protobuf/types/known/emptypb"

	"common/constvar"
	"common/errno"
	pb "common/proto/order"
	"common/util"
	"order/service"
)

//Auth 购物车服身份验证
func Auth(method string) bool {
	return true
}

//Order 订单服务处理器
type Order struct {
	srv service.IService
}

//New 实例化订单服务处理器
func New(srv service.IService) *Order {
	return &Order{srv: srv}
}

//SubmitOrder 提交订单
func (o *Order) SubmitOrder(ctx context.Context, req *pb.SubmitReq, reply *pb.OrderIDReply) error {
	id, err := o.srv.OrderSubmit(ctx, util.GetUserID(ctx), req.AddressId, req.CouponId,
		req.SkuIds, req.Note)
	if err != nil {
		return errno.OrderReplyErr(err)
	}
	reply.Data = id
	return nil
}

//SubmitSkuOrder 商品直接提交订单
func (o *Order) SubmitSkuOrder(ctx context.Context, req *pb.SkuSubmitReq, reply *pb.OrderIDReply) error {
	id, err := o.srv.SubmitSkuOrder(ctx, util.GetUserID(ctx), req.SkuId, req.AddressId, req.CouponId,
		int(req.Num), req.Note)
	if err != nil {
		return errno.OrderReplyErr(err)
	}
	reply.Data = id
	return nil
}

//OrderDetail 订单详情
func (o *Order) OrderDetail(ctx context.Context, req *pb.OrderIDReq, reply *pb.OrderInfoReply) error {
	info, err := o.srv.OrderDetail(ctx, util.GetUserID(ctx), req.OrderId)
	if err != nil {
		return errno.OrderReplyErr(err)
	}
	reply.Data = info
	return nil
}

//OrderCancel 订单取消
func (o *Order) OrderCancel(ctx context.Context, req *pb.OrderIDReq, empty *emptypb.Empty) error {
	if err := o.srv.OrderCancel(ctx, util.GetUserID(ctx), req.OrderId); err != nil {
		return errno.OrderReplyErr(err)
	}
	return nil
}

//OrderList 订单列表
func (o *Order) OrderList(ctx context.Context, req *pb.ListReq, reply *pb.ListReply) error {
	list, err := o.srv.MyOrderList(ctx, util.GetUserID(ctx), int(req.Status),
		util.GetPageOffset(req.Page), constvar.DefaultLimit)
	if err != nil {
		return errno.OrderReplyErr(err)
	}
	reply.Data = list
	return nil
}

//OrderPayNotify 订单支付成功通知
func (o *Order) OrderPayNotify(ctx context.Context, req *pb.PayNotifyReq, empty *emptypb.Empty) error {
	if err := o.srv.OrderPayNotify(ctx, util.GetUserID(ctx), int(req.PayAmount), int8(req.PayType),
		req.OrderNo, req.TradeNo, req.TransHash); err != nil {
		return errno.OrderReplyErr(err)
	}
	return nil
}

//OrderRefund 订单退款
func (o *Order) OrderRefund(ctx context.Context, req *pb.RefundReq, empty *emptypb.Empty) error {
	if err := o.srv.OrderRefund(ctx, util.GetUserID(ctx), req.OrderId, req.Content); err != nil {
		return errno.OrderReplyErr(err)
	}
	return nil
}

//OrderConfirmReceipt 确认收货
func (o *Order) OrderConfirmReceipt(ctx context.Context, req *pb.OrderIDReq, empty *emptypb.Empty) error {
	if err := o.srv.OrderConfirmReceipt(ctx, util.GetUserID(ctx), req.OrderId); err != nil {
		return errno.OrderReplyErr(err)
	}
	return nil
}

//OrderComment 评价
func (o *Order) OrderComment(ctx context.Context, req *pb.CommentReq, empty *emptypb.Empty) error {
	if err := o.srv.OrderComment(ctx, util.GetUserID(ctx), req.OrderId, req.SkuIds, int8(req.Star),
		req.Content, req.Resources); err != nil {
		return errno.OrderReplyErr(err)
	}
	return nil
}

//GetOrderByID 订单信息
func (o *Order) GetOrderByID(ctx context.Context, req *pb.OrderIDReq, reply *pb.OrderInfo) error {
	info, err := o.srv.OrderInfo(ctx, req.OrderId)
	if err != nil {
		return errno.OrderReplyErr(err)
	}
	reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(info).Elem())
	return nil
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gateway

import (
	context "context"
	common "gateway/proto/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	/// 从购物车提交订单
	SubmitOrder(ctx context.Context, in *SubmitReq, opts ...grpc.CallOption) (*OrderIDReply, error)
	/// 商品直接提交订单
	SubmitSkuOrder(ctx context.Context, in *SkuSubmitReq, opts ...grpc.CallOption) (*OrderIDReply, error)
	/// 订单详情
	OrderDetail(ctx context.Context, in *OrderIDReq, opts ...grpc.CallOption) (*OrderInfoReply, error)
	/// 订单取消
	OrderCancel(ctx context.Context, in *OrderIDReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error)
	/// 订单列表
	OrderList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListReply, error)
	/// 订单支付成功回调
	OrderPayNotify(ctx context.Context, in *PayNotifyReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error)
	/// 订单退款
	OrderRefund(ctx context.Context, in *RefundReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error)
	/// 订单确认收货
	OrderConfirmReceipt(ctx context.Context, in *OrderIDReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error)
	/// 订单评价
	OrderComment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) SubmitOrder(ctx context.Context, in *SubmitReq, opts ...grpc.CallOption) (*OrderIDReply, error) {
	out := new(OrderIDReply)
	err := c.cc.Invoke(ctx, "/order.Order/SubmitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) SubmitSkuOrder(ctx context.Context, in *SkuSubmitReq, opts ...grpc.CallOption) (*OrderIDReply, error) {
	out := new(OrderIDReply)
	err := c.cc.Invoke(ctx, "/order.Order/SubmitSkuOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderDetail(ctx context.Context, in *OrderIDReq, opts ...grpc.CallOption) (*OrderInfoReply, error) {
	out := new(OrderInfoReply)
	err := c.cc.Invoke(ctx, "/order.Order/OrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderCancel(ctx context.Context, in *OrderIDReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error) {
	out := new(common.SuccessEmptyReply)
	err := c.cc.Invoke(ctx, "/order.Order/OrderCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/order.Order/OrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderPayNotify(ctx context.Context, in *PayNotifyReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error) {
	out := new(common.SuccessEmptyReply)
	err := c.cc.Invoke(ctx, "/order.Order/OrderPayNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderRefund(ctx context.Context, in *RefundReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error) {
	out := new(common.SuccessEmptyReply)
	err := c.cc.Invoke(ctx, "/order.Order/OrderRefund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderConfirmReceipt(ctx context.Context, in *OrderIDReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error) {
	out := new(common.SuccessEmptyReply)
	err := c.cc.Invoke(ctx, "/order.Order/OrderConfirmReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderComment(ctx context.Context, in *CommentReq, opts ...grpc.CallOption) (*common.SuccessEmptyReply, error) {
	out := new(common.SuccessEmptyReply)
	err := c.cc.Invoke(ctx, "/order.Order/OrderComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	/// 从购物车提交订单
	SubmitOrder(context.Context, *SubmitReq) (*OrderIDReply, error)
	/// 商品直接提交订单
	SubmitSkuOrder(context.Context, *SkuSubmitReq) (*OrderIDReply, error)
	/// 订单详情
	OrderDetail(context.Context, *OrderIDReq) (*OrderInfoReply, error)
	/// 订单取消
	OrderCancel(context.Context, *OrderIDReq) (*common.SuccessEmptyReply, error)
	/// 订单列表
	OrderList(context.Context, *ListReq) (*ListReply, error)
	/// 订单支付成功回调
	OrderPayNotify(context.Context, *PayNotifyReq) (*common.SuccessEmptyReply, error)
	/// 订单退款
	OrderRefund(context.Context, *RefundReq) (*common.SuccessEmptyReply, error)
	/// 订单确认收货
	OrderConfirmReceipt(context.Context, *OrderIDReq) (*common.SuccessEmptyReply, error)
	/// 订单评价
	OrderComment(context.Context, *CommentReq) (*common.SuccessEmptyReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) SubmitOrder(context.Context, *SubmitReq) (*OrderIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedOrderServer) SubmitSkuOrder(context.Context, *SkuSubmitReq) (*OrderIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSkuOrder not implemented")
}
func (UnimplementedOrderServer) OrderDetail(context.Context, *OrderIDReq) (*OrderInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetail not implemented")
}
func (UnimplementedOrderServer) OrderCancel(context.Context, *OrderIDReq) (*common.SuccessEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCancel not implemented")
}
func (UnimplementedOrderServer) OrderList(context.Context, *ListReq) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServer) OrderPayNotify(context.Context, *PayNotifyReq) (*common.SuccessEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPayNotify not implemented")
}
func (UnimplementedOrderServer) OrderRefund(context.Context, *RefundReq) (*common.SuccessEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderRefund not implemented")
}
func (UnimplementedOrderServer) OrderConfirmReceipt(context.Context, *OrderIDReq) (*common.SuccessEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderConfirmReceipt not implemented")
}
func (UnimplementedOrderServer) OrderComment(context.Context, *CommentReq) (*common.SuccessEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderComment not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/SubmitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).SubmitOrder(ctx, req.(*SubmitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_SubmitSkuOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuSubmitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).SubmitSkuOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/SubmitSkuOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).SubmitSkuOrder(ctx, req.(*SkuSubmitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderDetail(ctx, req.(*OrderIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderCancel(ctx, req.(*OrderIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderPayNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayNotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderPayNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderPayNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderPayNotify(ctx, req.(*PayNotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderRefund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderRefund(ctx, req.(*RefundReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderConfirmReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderConfirmReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderConfirmReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderConfirmReceipt(ctx, req.(*OrderIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderComment(ctx, req.(*CommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrder",
			Handler:    _Order_SubmitOrder_Handler,
		},
		{
			MethodName: "SubmitSkuOrder",
			Handler:    _Order_SubmitSkuOrder_Handler,
		},
		{
			MethodName: "OrderDetail",
			Handler:    _Order_OrderDetail_Handler,
		},
		{
			MethodName: "OrderCancel",
			Handler:    _Order_OrderCancel_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _Order_OrderList_Handler,
		},
		{
			MethodName: "OrderPayNotify",
			Handler:    _Order_OrderPayNotify_Handler,
		},
		{
			MethodName: "OrderRefund",
			Handler:    _Order_OrderRefund_Handler,
		},
		{
			MethodName: "OrderConfirmReceipt",
			Handler:    _Order_OrderConfirmReceipt_Handler,
		},
		{
			MethodName: "OrderComment",
			Handler:    _Order_OrderComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/order.proto",
}
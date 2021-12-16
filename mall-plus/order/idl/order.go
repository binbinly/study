package idl

import (
	"strings"

	pb "common/proto/order"
	"common/util"
	"order/conf"
	"order/model"
)

//TransferOrder 转换订单详情输出
func TransferOrder(order *model.OrderModel) *pb.OrderInfo {
	res := &pb.OrderInfo{
		Id:              order.ID,
		OrderNo:         order.OrderNo,
		Note:            order.Note,
		TotalAmount:     util.ParseAmount(order.TotalAmount),
		Amount:          util.ParseAmount(order.Amount),
		CouponAmount:    util.ParseAmount(order.CouponAmount),
		FreightAmount:   util.ParseAmount(order.FreightAmount),
		PayAmount:       util.ParseAmount(order.PayAmount),
		PayType:         int32(order.PayType),
		PayAt:           order.PayAt,
		CreateAt:        order.CreatedAt,
		Status:          int32(order.Status),
		TradeNo:         order.TradeNo,
		DeliveryCompany: order.DeliveryCompany,
		DeliveryNo:      order.DeliveryNo,
		Integration:     int32(order.Integration),
		Growth:          int32(order.Growth),
		DeliveryAt:      order.DeliveryAt,
		ReceiveAt:       order.ReceiveAt,
		CommentAt:       order.CommentAt,
		Address: &pb.Address{
			Name:   order.AddressName,
			Phone:  order.AddressPhone,
			Area:   strings.Join([]string{order.AddressProvince, order.AddressCity, order.AddressCounty}, " "),
			Detail: order.AddressDetail,
		},
		Items: make([]*pb.OrderSku, 0, len(order.Items)),
	}
	for _, item := range order.Items {
		res.Items = append(res.Items, &pb.OrderSku{
			SkuId:     item.SkuID,
			Title:     item.SkuName,
			Cover:     util.BuildResUrl(conf.Conf.DFS, item.SkuImg),
			Price:     util.ParseAmount(item.SkuPrice),
			Num:       int32(item.Num),
			AttrValue: item.SkuAttrs,
		})
	}
	return res
}

//TransferOrderList 转换订单列表输出
func TransferOrderList(list []*model.OrderModel) []*pb.OrderList {
	if len(list) == 0 {
		return []*pb.OrderList{}
	}
	res := make([]*pb.OrderList, 0, len(list))
	for _, order := range list {
		ol := &pb.OrderList{
			Id:      order.ID,
			OrderNo: order.OrderNo,
			Amount:  util.ParseAmount(order.Amount),
			Status:  int32(order.Status),
			Time:    order.CreatedAt,
			Items:   make([]*pb.OrderSku, 0, len(order.Items)),
		}

		for _, item := range order.Items {
			ol.Items = append(ol.Items, &pb.OrderSku{
				SkuId:     item.SkuID,
				Title:     item.SkuName,
				Cover:     util.BuildResUrl(conf.Conf.DFS, item.SkuImg),
				Price:     util.ParseAmount(item.SkuPrice),
				Num:       int32(item.Num),
				AttrValue: item.SkuAttrs,
			})
		}
		res = append(res, ol)
	}
	return res
}

package idl

import (
	"mall/app/constvar"
	"mall/app/model"
)

//TransferOrder 转换订单详情
func TransferOrder(order *model.OrderModel) *model.Order {
	res := &model.Order{
		ID:           order.ID,
		OrderNo:      order.OrderNo,
		UserNote:     order.UserNote,
		TotalPrice:   constvar.ParseAmount(order.TotalPrice),
		Amount:       constvar.ParseAmount(order.Amount),
		CouponAmount: constvar.ParseAmount(order.CouponAmount),
		PayType:      order.PayType,
		PayStatus:    order.PayStatus,
		PayAmount:    constvar.ParseAmount(order.PayAmount),
		PayAt:        order.PayAt,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt,
		Address: &model.OrderAddress{
			Name:   order.Address.Name,
			Phone:  order.Address.Phone,
			Area:   order.Address.Area,
			Detail: order.Address.Detail,
		},
		Goods: make([]*model.OrderGoods, 0, len(order.Goods)),
	}
	for _, good := range order.Goods {
		res.Goods = append(res.Goods, &model.OrderGoods{
			GoodsID:   good.GoodsID,
			GoodsName: good.GoodsName,
			Cover:     constvar.BuildResUrl(good.GoodsCover),
			Price:     constvar.ParseAmount(good.GoodsPrice),
			Num:       good.BuyCount,
			SkuName:   good.Attrs,
		})
	}
	return res
}

//TransferOrderList 转换订单列表
func TransferOrderList(list []*model.OrderModel) []*model.OrderList {
	if len(list) == 0 {
		return make([]*model.OrderList, 0)
	}
	res := make([]*model.OrderList, 0, len(list))
	for _, order := range list {
		ol := &model.OrderList{
			ID:        order.ID,
			OrderNo:   order.OrderNo,
			Amount:    constvar.ParseAmount(order.Amount),
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			Goods:     make([]*model.OrderGoods, 0, len(order.Goods)),
		}

		for _, good := range order.Goods {
			ol.Goods = append(ol.Goods, &model.OrderGoods{
				GoodsID:   good.GoodsID,
				GoodsName: good.GoodsName,
				Cover:     constvar.BuildResUrl(good.GoodsCover),
				Price:     constvar.ParseAmount(good.GoodsPrice),
				Num:       good.BuyCount,
				SkuName:   good.Attrs,
			})
		}
		res = append(res, ol)
	}
	return res
}

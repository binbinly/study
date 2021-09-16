package idl

import (
	"mall/app/constvar"
	"mall/app/model"
)

//TransferCart 转换购物车接口输出
func TransferCart(cart *model.CartModel) *model.Cart {
	if cart == nil {
		return &model.Cart{}
	}

	return &model.Cart{
		ID:        cart.ID,
		GoodsID:   cart.GoodsID,
		GoodsName: cart.GoodsName,
		Price:     constvar.ParseAmount(cart.Price),
		Cover:     constvar.BuildResUrl(cart.Cover),
		SkuID:     cart.SkuID,
		SkuName:   cart.SkuName,
		Num:       cart.Num,
	}
}

//TransferCartList 转换购物车列表结构
func TransferCartList(list []*model.CartModel) []*model.Cart {
	if len(list) == 0 {
		return make([]*model.Cart, 0)
	}
	res := make([]*model.Cart, 0, len(list))
	for _, cart := range list {
		res = append(res, TransferCart(cart))
	}
	return res
}
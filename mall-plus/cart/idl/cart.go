package idl

import (
	"cart/conf"
	"cart/model"
	pb "common/proto/cart"
	"common/util"
)

//TransferCart 转换购物车接口输出
func TransferCart(cart *model.CartModel, external bool) *pb.CartItem {
	if cart == nil {
		return &pb.CartItem{}
	}

	item := &pb.CartItem{
		SkuId:   cart.SkuID,
		Title:   cart.Title,
		Price:   float64(cart.Price),
		Cover:   cart.Cover,
		SkuAttr: cart.SkuAttr,
		Num:     int32(cart.Num),
	}
	if external { //对外格式化数据
		item.Price = util.ParseAmount(cart.Price)
		item.Cover = util.BuildResUrl(conf.Conf.DFS, cart.Cover)
	}
	return item
}

//TransferCartList 转换购物车列表结构
func TransferCartList(list []*model.CartModel, external bool) []*pb.CartItem {
	if len(list) == 0 {
		return make([]*pb.CartItem, 0)
	}
	res := make([]*pb.CartItem, 0, len(list))
	for _, cart := range list {
		res = append(res, TransferCart(cart, external))
	}
	return res
}

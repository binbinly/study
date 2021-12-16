package service

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"cart/idl"
	"cart/model"
	pb "common/proto/cart"
)

//AddCart 添加购物车
func (s *Service) AddCart(ctx context.Context, userID, skuID int64, num int) error {
	//获取当前sku_id 的购物车信息
	cart, err := s.repo.GetCartByID(ctx, userID, skuID)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] cart uid: %v id: %v", userID, skuID)
	}
	if cart.SkuID != 0 { //商品已存在,加数量即可
		cart.Num += num
	} else {
		err = s.buildCart(ctx, skuID, num, cart)
		if err != nil {
			return err
		}
	}
	err = s.repo.AddCart(ctx, userID, cart)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] add cart")
	}
	return nil
}

//EditCart 更新购物车
func (s *Service) EditCart(ctx context.Context, userID, oldSkuID, newSkuID int64, num int) error {
	cart, err := s.repo.GetCartByID(ctx, userID, newSkuID)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] get by uid: %v, id: %v", userID, newSkuID)
	}
	if cart.SkuID == newSkuID && cart.Num == num { //购物车未更新，直接返回
		return nil
	}
	if oldSkuID == newSkuID { // 商品未修改更改数量即可
		cart.Num = num
	} else {
		err = s.buildCart(ctx, newSkuID, num, cart)
		if err != nil {
			return err
		}
	}
	err = s.repo.EditCart(ctx, userID, oldSkuID, cart)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] edit sku")
	}
	return nil
}

//EditCartNum 修改商品数量
func (s *Service) EditCartNum(ctx context.Context, userID, skuID int64, num int) error {
	cart, err := s.repo.GetCartByID(ctx, userID, skuID)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] get by uid: %v, id: %v", userID, skuID)
	}
	if cart.SkuID == 0 { // 商品不存在，直接添加
		err = s.buildCart(ctx, skuID, num, cart)
		if err != nil {
			return err
		}
	}
	if cart.Num == num { // 数量未修改，直接返回
		return nil
	}
	cart.Num = num
	err = s.repo.EditCart(ctx, userID, 0, cart)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] edit num")
	}
	return nil
}

//DelCart 删除购物车
func (s *Service) DelCart(ctx context.Context, userID int64, skuIds []int64) error {
	return s.repo.DelCart(ctx, userID, skuIds)
}

//ClearCart 清空购物车
func (s *Service) ClearCart(ctx context.Context, userID int64) error {
	return s.repo.EmptyCart(ctx, userID)
}

//CartList 购物车
func (s *Service) CartList(ctx context.Context, userID int64) ([]*pb.CartItem, error) {
	list, err := s.repo.CartList(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] list uid: %v", userID)
	}
	sort.Sort(model.CartSort(list))

	return idl.TransferCartList(list, true), nil
}

//BatchGetCarts 批量获取购物车信息
func (s *Service) BatchGetCarts(ctx context.Context, userID int64, skuIds []int64) ([]*pb.CartItem, error) {
	carts, err := s.repo.GetCartsByIds(ctx, userID, skuIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] get carts by uid: %v ids: %v", userID, skuIds)
	}
	return idl.TransferCartList(carts, false), nil
}

//buildCart 构建购物车结构
func (s *Service) buildCart(ctx context.Context, skuID int64, num int, cart *model.CartModel) error {
	sku, err := s.GetSkuByID(ctx, skuID)
	if err != nil {
		return err
	}
	cart.SkuID = skuID
	cart.Num = num
	cart.Cover = sku.Cover
	cart.Price = int(sku.Price)
	cart.Title = sku.Title
	cart.SkuAttr = sku.AttrValue
	return nil
}

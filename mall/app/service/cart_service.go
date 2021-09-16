package service

import (
	"context"
	"fmt"
	"mall/app/idl"
	"sort"
	"time"

	"github.com/pkg/errors"

	"mall/app/model"
)

//ICart 购物车服务接口
type ICart interface {
	AddCart(ctx context.Context, userID, goodsID, skuID, num int) (*model.Cart, error)
	EditCart(ctx context.Context, userID, skuID, num int, id string) (*model.Cart, error)
	EditCartNum(ctx context.Context, userID int, id string, num int) error
	DelCart(ctx context.Context, userID int, id string) error
	EmptyCart(ctx context.Context, userID int) error
	CartList(ctx context.Context, userID int) ([]*model.Cart, error)
}

//AddCart 添加购物车
func (s *Service) AddCart(ctx context.Context, userID, goodsID, skuID, num int) (*model.Cart, error) {
	//获取已添加的商品，匹配是否是新商品新规格
	carts, err := s.repo.CartList(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] carts uid: %v", userID)
	}
	var cart *model.CartModel
	for _, ct := range carts {
		if ct.GoodsID == goodsID && ct.SkuID == skuID { //商品规格已存在,加数量即可
			cart = ct
			cart.Num += num
		}
	}
	if cart == nil { //新商品规格添加
		goods, err := s.repo.GoodsDetail(ctx, goodsID)
		if err != nil {
			return nil, errors.Wrapf(err, "[service.cart] detail by goods_id: %v", goodsID)
		}
		if goods.ID == 0 {
			return nil, ErrGoodsNotFound
		}
		price := goods.Price
		skuName := ""
		if skuID > 0 {
			sku, err := s.repo.GetSkuByID(ctx, skuID)
			if err != nil {
				return nil, errors.Wrapf(err, "[service.cart] sku by id: %v", skuID)
			}
			if sku.ID == 0 {
				return nil, ErrGoodsSkuNotFound
			}
			price = sku.Price
			skuName = sku.ValueNames
		}
		cart = &model.CartModel{
			ID:        fmt.Sprintf("%d%d", goods.ID, time.Now().Unix()),
			GoodsID:   goods.ID,
			GoodsName: goods.Title,
			Price:     price,
			Cover:     goods.Cover,
			SkuID:     skuID,
			SkuName:   skuName,
			Num:       num,
			UTime:     time.Now().Unix(),
		}
	}
	err = s.repo.AddCart(ctx, userID, cart)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] edit cart")
	}
	return idl.TransferCart(cart), nil
}

//EditCart 更新购物车
func (s *Service) EditCart(ctx context.Context, userID, skuID, num int, id string) (*model.Cart, error) {
	sku, err := s.repo.GetSkuByID(ctx, skuID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] sku by id: %v", skuID)
	}
	if sku.ID == 0 {
		return nil, ErrGoodsSkuNotFound
	}
	cart, err := s.repo.GetCartByID(ctx, userID, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] get by uid: %v, id: %v", userID, id)
	}
	if cart.SkuID == skuID && cart.Num == num {
		return nil, ErrGoodsSkuNotEdit
	}
	if cart.SkuID != skuID {
		cart.Price = sku.Price
		cart.SkuID = skuID
		cart.SkuName = sku.ValueNames
	}
	cart.Num = num
	err = s.repo.EditCart(ctx, userID, id, cart)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] edit sku")
	}
	return idl.TransferCart(cart), nil
}

//EditCartNum 修改商品数量
func (s *Service) EditCartNum(ctx context.Context, userID int, id string, num int) error {
	cart, err := s.repo.GetCartByID(ctx, userID, id)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] get by uid: %v, id: %v", userID, id)
	}
	if cart.Num == num {
		return nil
	}
	cart.Num = num
	err = s.repo.EditCart(ctx, userID, id, cart)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] edit num")
	}
	return nil
}

//DelCart 删除购物车
func (s *Service) DelCart(ctx context.Context, userID int, id string) error {
	return s.repo.DelCart(ctx, userID, []string{id})
}

//EmptyCart 清空购物车
func (s *Service) EmptyCart(ctx context.Context, userID int) error {
	return s.repo.EmptyCart(ctx, userID)
}

//CartList 购物车
func (s *Service) CartList(ctx context.Context, userID int) ([]*model.Cart, error) {
	list, err := s.repo.CartList(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] list uid: %v", userID)
	}
	sort.Sort(model.CartSort(list))

	return idl.TransferCartList(list), nil
}

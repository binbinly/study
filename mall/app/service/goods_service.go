package service

import (
	"context"
	"github.com/pkg/errors"
	"mall/app/constvar"
	"mall/app/idl"
	"mall/app/model"
	"mall/pkg/redis"
)

type IGoods interface {
	GoodsSku(ctx context.Context, goodsID int) (*model.GoodsDetailSku, error)
	GoodsList(ctx context.Context, catID, orderType int, keyword, price, field, order string, offset, limit int) ([]*model.GoodsList, error)
	GoodsDetail(ctx context.Context, goodsID int) (*model.GoodsDetail, error)
}

var (
	//ErrGoodsNotFound 商品不存在
	ErrGoodsNotFound = errors.New("goods not found")
	//ErrGoodsSkuNotFound 商品销售属性不存在
	ErrGoodsSkuNotFound = errors.New("goods sku not found")
	//ErrGoodsSkuNotEdit 商品规格未修改
	ErrGoodsSkuNotEdit = errors.New("goods sku not edit")
)

//GoodsList 商品列表
func (s *Service) GoodsList(ctx context.Context, catID, orderType int, keyword, price, field, order string, offset, limit int) ([]*model.GoodsList, error) {
	var ids []int
	var err error
	if catID > 0 {
		//获取分类下所有子分类
		ids, err = s.repo.GoodsCategoryChild(ctx, catID)
		if err != nil {
			return nil, errors.Wrapf(err, "[service.goods] catID child by %v", catID)
		}
	}
	if keyword != "" {
		redis.Client.ZIncrBy(ctx, constvar.HotSearchKey, 1, keyword)
	}
	list, err := s.repo.GoodsList(ctx, ids, keyword, price, field, order, orderType, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] list by ids: %v, orderType: %v", ids, orderType)
	}
	return idl.TransferGoodsList(list), nil
}

//GoodsDetail 商品详情
func (s *Service) GoodsDetail(ctx context.Context, goodsID int) (*model.GoodsDetail, error) {
	goods, err := s.repo.GoodsDetail(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] detail by goods_id: %v", goodsID)
	}
	if goods.ID == 0 {
		return nil, ErrGoodsNotFound
	}
	goodsImage, err := s.repo.GetImagesByGID(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] images by goods_id: %v", goodsID)
	}
	goodsAttr, err := s.repo.GetAttrByGID(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] attrs by goods_id: %v", goodsID)
	}
	goodsSku, err := s.repo.GetSkusByGID(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] skus by goods_id: %v", goodsID)
	}
	goodsSkuAttr, err := s.repo.GetSkusAttrsByGID(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] sku attrs by goods_id: %v", goodsID)
	}
	input := &idl.TransferGoodsInput{
		Goods:        goods,
		GoodsAttr:    goodsAttr,
		GoodsSku:     goodsSku,
		GoodsSkuAttr: goodsSkuAttr,
		GoodsImage:   goodsImage,
	}
	return idl.TransFerGoodsDetail(input), nil
}

//GoodsSku 商品sku
func (s *Service) GoodsSku(ctx context.Context, goodsID int) (*model.GoodsDetailSku, error) {
	goods, err := s.repo.GoodsDetail(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] detail by goods_id: %v", goodsID)
	}
	if goods.ID == 0 {
		return nil, ErrGoodsNotFound
	}
	goodsSku, err := s.repo.GetSkusByGID(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] skus by goods_id: %v", goodsID)
	}
	goodsSkuAttr, err := s.repo.GetSkusAttrsByGID(ctx, goodsID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.goods] sku attrs by goods_id: %v", goodsID)
	}
	input := &idl.TransferGoodsInput{
		Goods:        goods,
		GoodsSku:     goodsSku,
		GoodsSkuAttr: goodsSkuAttr,
	}
	return idl.TransFerGoodsDetailSku(input), nil
}
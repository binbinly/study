package idl

import (
	"encoding/json"

	"mall/app/constvar"
	"mall/app/model"
)

// TransferGoodsInput 转换输入字段
type TransferGoodsInput struct {
	Goods        *model.GoodsModel
	GoodsAttr    []*model.GoodsAttrModel
	GoodsSku     []*model.GoodsSkuModel
	GoodsSkuAttr []*model.GoodsSkuAttrModel
	GoodsImage   []*model.GoodsImageModel
}

//TransferGoodsList 转换输出商品列表
func TransferGoodsList(list []*model.GoodsModel) (goods []*model.GoodsList) {
	if len(list) == 0 {
		return []*model.GoodsList{}
	}

	for _, m := range list {
		goods = append(goods, &model.GoodsList{
			ID:            m.ID,
			Title:         m.Title,
			Cover:         constvar.BuildResUrl(m.Cover),
			Price:         constvar.ParseAmount(m.Price),
			OriginalPrice: constvar.ParseAmount(m.OriginalPrice),
			Intro:         m.Intro,
		})
	}
	return
}

//TransFerGoodsDetail 商品详情转换输出
func TransFerGoodsDetail(input *TransferGoodsInput) *model.GoodsDetail {
	goods := &model.GoodsDetail{
		ID:            input.Goods.ID,
		CatID:         input.Goods.CatID,
		Title:         input.Goods.Title,
		Cover:         constvar.BuildResUrl(input.Goods.Cover),
		Price:         constvar.ParseAmount(input.Goods.Price),
		OriginalPrice: constvar.ParseAmount(input.Goods.OriginalPrice),
		Intro:         input.Goods.Intro,
		Unit:          input.Goods.Unit,
		Stock:         input.Goods.Stock,
		SkuMany:       input.Goods.SkuMany,
		Discount:      input.Goods.Discount,
		SaleCount:     input.Goods.SaleCount,
		ReviewCount:   input.Goods.ReviewCount,
		Attrs:         make(map[string]string),
		Skus:          make([]*model.GoodsSku, 0),
		BannerUrl:     make([]string, 0),
		MainUrl:       make([]string, 0),
		SkuAttrs:      make([]*model.GoodsSkuAttr, 0),
	}
	//商品图片
	for _, image := range input.GoodsImage {
		if image.Type == model.ImageTypeBanner {
			goods.BannerUrl = append(goods.BannerUrl, image.Url)
		} else {
			goods.MainUrl = append(goods.MainUrl, constvar.BuildResUrl(image.Url))
		}
	}

	//商品属性参数
	for _, attr := range input.GoodsAttr {
		goods.Attrs[attr.Name] = attr.Value
	}

	//商品SKU
	for _, sku := range input.GoodsSku {
		goods.Skus = append(goods.Skus, &model.GoodsSku{
			ID:            sku.ID,
			Attrs:         sku.Attrs,
			Values:        sku.Values,
			Stock:         sku.Stock,
			Price:         constvar.ParseAmount(sku.Price),
			OriginalPrice: constvar.ParseAmount(sku.OriginalPrice),
		})
	}

	//商品销售属性
	for _, skuAttr := range input.GoodsSkuAttr {
		goods.SkuAttrs = append(goods.SkuAttrs, &model.GoodsSkuAttr{
			ID:     skuAttr.AttrID,
			Name:   skuAttr.AttrName,
			Values: json.RawMessage(skuAttr.Values),
		})
	}
	return goods
}

//TransFerGoodsDetailSku 商品详情sku转换输出
func TransFerGoodsDetailSku(input *TransferGoodsInput) *model.GoodsDetailSku {
	goods := &model.GoodsDetailSku{
		ID:       input.Goods.ID,
		Stock:    input.Goods.Stock,
		SkuMany:  input.Goods.SkuMany,
		Skus:     make([]*model.GoodsSku, 0),
		SkuAttrs: make([]*model.GoodsSkuAttr, 0),
	}

	//商品SKU
	for _, sku := range input.GoodsSku {
		goods.Skus = append(goods.Skus, &model.GoodsSku{
			ID:            sku.ID,
			Attrs:         sku.Attrs,
			Values:        sku.Values,
			Stock:         sku.Stock,
			Price:         constvar.ParseAmount(sku.Price),
			OriginalPrice: constvar.ParseAmount(sku.OriginalPrice),
		})
	}

	//商品销售属性
	for _, skuAttr := range input.GoodsSkuAttr {
		goods.SkuAttrs = append(goods.SkuAttrs, &model.GoodsSkuAttr{
			ID:     skuAttr.AttrID,
			Name:   skuAttr.AttrName,
			Values: json.RawMessage(skuAttr.Values),
		})
	}
	return goods
}

package idl

import (
	"encoding/json"

	pb "common/proto/product"
	"common/util"
	"product/conf"
	"product/model"
)

//TransferSkuInput sku转换输入
type TransferSkuInput struct {
	Sku        *model.SkuModel
	Spu        *model.SpuModel
	SpuDesc    *model.SpuDescModel
	Skus       []*model.SkuModel
	SkuImages  []*model.SkuImageModel
	AttrGroups []*model.AttrGroupModel
	SkuAttrs   []*model.SkuAttrModel
	SpuAttrs   []*model.Attrs
	Stocks     map[int64]int32
}

//TransferSkuInfo 转换sku信息
func TransferSkuInfo(sku *model.SkuModel) *pb.SkuInfo {
	return &pb.SkuInfo{
		Id:        sku.ID,
		SpuId:     sku.SpuID,
		CatId:     sku.CatID,
		BrandId:   sku.BrandID,
		Title:     sku.Title,
		Desc:      sku.Desc,
		Cover:     sku.Cover,
		Subtitle:  sku.Subtitle,
		Price:     int64(sku.Price),
		SaleCount: int64(sku.SaleCount),
		AttrValue: sku.AttrValue,
	}
}

//TransferSku 转换sku数据输出
func TransferSku(input *TransferSkuInput) *pb.Sku {
	sku := new(pb.Sku)
	sku.Id = input.Sku.ID
	sku.CatId = input.Sku.CatID
	sku.SpuId = input.Sku.SpuID
	sku.BrandId = input.Sku.BrandID
	sku.Title = input.Sku.Title
	sku.Subtitle = input.Sku.Subtitle
	sku.Desc = input.Sku.Desc
	sku.Price = util.ParseAmount(input.Sku.Price)
	sku.Cover = util.BuildResUrl(conf.Conf.DFS, input.Sku.Cover)
	sku.SaleCount = int64(input.Sku.SaleCount)
	sku.IsMany = input.Spu.IsMany == 1
	sku.Stock = input.Stocks[input.Sku.ID]
	for _, image := range input.SkuImages {
		sku.Banners = append(sku.Banners, util.BuildResUrl(conf.Conf.DFS, image.Img))
	}
	if input.SpuDesc == nil || input.SpuDesc.Content == "" {
		sku.Mains = []string{}
	} else {
		var imgs []string
		_ = json.Unmarshal([]byte(input.SpuDesc.Content), &imgs)
		sku.Mains = buildResUrls(imgs)
	}
	sku.Attrs = transferAttrs(input.SpuAttrs, input.AttrGroups)
	sku.Skus, sku.SaleAttrs = transferSkuAttrs(input.SkuAttrs, input.Skus, input.Stocks)
	return sku
}

//TransferSkuSaleAttrs 转换sku销售属性
func TransferSkuSaleAttrs(input *TransferSkuInput) *pb.SkuSaleAttr {
	sku := new(pb.SkuSaleAttr)
	sku.Id = input.Sku.ID
	sku.IsMany = input.Spu.IsMany == 1
	sku.Skus, sku.SaleAttrs = transferSkuAttrs(input.SkuAttrs, input.Skus, input.Stocks)
	return sku
}

//transferAttrs 转换规格属性分组及其下属性
func transferAttrs(spuAttrs []*model.Attrs, attrGroups []*model.AttrGroupModel) (as []*pb.Attrs) {
	attrMap := make(map[int64]*pb.Attrs)
	for _, attr := range spuAttrs {
		// TODO 响应数据暂时不考虑属性分组
		attr.GroupID = 0
		if attrs, ok := attrMap[attr.GroupID]; ok { //分组已存在，追加其下的属性
			attrs.Items = append(attrs.Items, &pb.Attr{
				Id:    attr.AttrID,
				Name:  attr.AttrName,
				Value: attr.AttrValue,
			})
		} else { //添加新分组
			attrMap[attr.GroupID] = &pb.Attrs{
				GroupId:   attr.GroupID,
				GroupName: getGroupNameByID(attrGroups, attr.GroupID),
				Items: []*pb.Attr{
					{
						Id:    attr.AttrID,
						Name:  attr.AttrName,
						Value: attr.AttrValue,
					},
				},
			}
		}
	}
	for _, attrs := range attrMap {
		as = append(as, attrs)
	}
	return
}

//transferSkuAttrs 转换销售属性
func transferSkuAttrs(skuAttrs []*model.SkuAttrModel, skus []*model.SkuModel, stocks map[int64]int32) (resSkus []*pb.Skus, resAttrs []*pb.SaleAttrs) {
	attrsMap := make(map[int64][]*pb.SkuAttr)
	skuAttrsMap := make(map[int64]*pb.SaleAttrs)
	for _, skuAttr := range skuAttrs {
		if attr, ok := attrsMap[skuAttr.SkuID]; ok {
			attr = append(attr, &pb.SkuAttr{
				AttrId:    skuAttr.AttrID,
				ValueId:   skuAttr.ID,
				AttrName:  skuAttr.AttrName,
				ValueName: skuAttr.AttrValue,
			})
		} else {
			attrsMap[skuAttr.SkuID] = []*pb.SkuAttr{
				{
					AttrId:    skuAttr.AttrID,
					ValueId:   skuAttr.ID,
					AttrName:  skuAttr.AttrName,
					ValueName: skuAttr.AttrValue,
				},
			}
		}
		if attrs, ok := skuAttrsMap[skuAttr.AttrID]; ok {
			attrs.Values = append(attrs.Values, &pb.SkuValue{
				Id:   skuAttr.ID,
				Name: skuAttr.AttrValue,
			})
		} else {
			skuAttrsMap[skuAttr.AttrID] = &pb.SaleAttrs{
				AttrId:   skuAttr.AttrID,
				AttrName: skuAttr.AttrName,
				Values: []*pb.SkuValue{
					{
						Id:   skuAttr.ID,
						Name: skuAttr.AttrValue,
					},
				},
			}
		}
	}
	for _, attrs := range skuAttrsMap {
		resAttrs = append(resAttrs, attrs)
	}
	for _, sku := range skus {
		attrs, ok := attrsMap[sku.ID]
		if !ok {
			attrs = make([]*pb.SkuAttr, 0)
		}
		resSkus = append(resSkus, &pb.Skus{
			SkuId: sku.ID,
			Price: util.ParseAmount(sku.Price),
			Stock: stocks[sku.ID],
			Attrs: attrs,
		})
	}
	return
}

//getGroupNameByID 获取分组名
func getGroupNameByID(groups []*model.AttrGroupModel, id int64) string {
	for _, group := range groups {
		if group.ID == id {
			return group.Name
		}
	}
	return ""
}

//buildResUrl 构建资源图片路径列表
func buildResUrls(urls []string) (res []string) {
	for _, url := range urls {
		res = append(res, util.BuildResUrl(conf.Conf.DFS, url))
	}
	return
}

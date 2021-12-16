package idl

import (
	pb "common/proto/product"
	"common/util"
	"product/conf"
	"product/es"
	"product/model"
)

//TransferSearchInput 搜索结果集入参
type TransferSearchInput struct {
	Result *es.SearchRes
	Cats   map[int64]string
	Brands map[int64]*model.Brand
}

//TransferSkuList 转换sku列表
func TransferSkuList(list []*es.ProductEs) (res []*pb.SkuEs) {
	if len(list) == 0 {
		return []*pb.SkuEs{}
	}
	for _, product := range list {
		res = append(res, &pb.SkuEs{
			Id:        product.SkuID,
			Title:     product.SkuTitle,
			Price:     util.ParseAmount(product.SkuPrice),
			Cover:     util.BuildResUrl(conf.Conf.DFS, product.SkuImg),
			SaleCount: product.SaleCount,
			HasStock:  product.HasStock,
		})
	}
	return
}

//TransferSearchRes 转换搜索结果集
func TransferSearchRes(input *TransferSearchInput) *pb.SearchReply {
	reply := new(pb.SearchReply)
	for _, product := range input.Result.Products {
		reply.Data = append(reply.Data, &pb.SkuEs{
			Id:        product.SkuID,
			Title:     product.SkuTitle,
			Price:     util.ParseAmount(product.SkuPrice),
			Cover:     util.BuildResUrl(conf.Conf.DFS, product.SkuImg),
			SaleCount: product.SaleCount,
			HasStock:  product.HasStock,
		})
	}
	for id, name := range input.Cats {
		reply.Cats = append(reply.Cats, &pb.CatEs{
			Id:   id,
			Name: name,
		})
	}
	for id, brand := range input.Brands {
		reply.Brands = append(reply.Brands, &pb.BrandEs{
			Id:   id,
			Name: brand.Name,
			Logo: util.BuildResUrl(conf.Conf.DFS, brand.Logo),
		})
	}
	reply.Attrs = parseAttrs(input.Result)

	return reply
}

//parseAttrs 解析规格属性
func parseAttrs(res *es.SearchRes) (attrs []*pb.AttrEs) {
	for _, bucket := range res.Attrs.AttrIDAgg.Buckets {
		attr := &pb.AttrEs{
			Id: int64(bucket.Key),
		}
		for _, subBucket := range bucket.AttrNameAgg.Buckets {
			attr.Name = subBucket.Key
		}
		for _, subBucket := range bucket.AttrValueAgg.Buckets {
			attr.Values = append(attr.Values, subBucket.Key)
		}
		attrs = append(attrs, attr)
	}
	return
}

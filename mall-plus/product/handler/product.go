package handler

import (
	"context"
	"reflect"

	"google.golang.org/protobuf/types/known/emptypb"

	"common/constvar"
	"common/errno"
	pb "common/proto/product"
	"common/util"
	"product/service"
)

//Auth 产品服身份验证
func Auth(method string) bool {
	return false
}

//Product 产品服务处理器
type Product struct {
	srv service.IService
}

//New 实例化产品服务
func New(srv service.IService) *Product {
	return &Product{srv: srv}
}

//CategoryTree 产品分类树
func (p *Product) CategoryTree(ctx context.Context, empty *emptypb.Empty, reply *pb.CategoryReply) error {
	tree, err := p.srv.CategoryTree(ctx)
	if err != nil {
		return errno.ProductReplyErr(err)
	}
	reply.Data = tree
	return nil
}

//SkuDetail sku商品详情
func (p *Product) SkuDetail(ctx context.Context, req *pb.SkuReq, reply *pb.SkuReply) error {
	sku, err := p.srv.SkuDetail(ctx, req.SkuId)
	if err != nil {
		return errno.ProductReplyErr(err)
	}
	reply.Data = sku
	return nil
}

//GetSkuSaleAttrs 获取sku的销售属性
func (p *Product) GetSkuSaleAttrs(ctx context.Context, req *pb.SkuReq, reply *pb.SkuSaleAttrReply) error {
	info, err := p.srv.GetSkuSaleAttrs(ctx, req.SkuId)
	if err != nil {
		return errno.ProductReplyErr(err)
	}
	reply.Data = info
	return nil
}

//GetSkuByID 获取sku信息
func (p *Product) GetSkuByID(ctx context.Context, req *pb.SkuReq, reply *pb.SkuInfoInternal) error {
	info, err := p.srv.GetSkuInfo(ctx, req.SkuId)
	if err != nil {
		return errno.ProductReplyErr(err)
	}
	reply.Info = info
	return nil
}

//SpuComment 商品评价
func (p *Product) SpuComment(ctx context.Context, req *pb.CommentReq, empty *emptypb.Empty) error {
	err := p.srv.SpuComment(ctx, req.SkuIds, req.UserId, req.OrderId, int8(req.Star), req.Content, req.Resources)
	if err != nil {
		return errno.ProductReplyErr(err)
	}
	return nil
}

//SkuList sku商品列表
func (p *Product) SkuList(ctx context.Context, req *pb.SkuListReq, reply *pb.SkuListReply) error {
	list, err := p.srv.SkuList(ctx, req.CatId, util.GetPageOffset(req.Page), constvar.DefaultLimit)
	if err != nil {
		return errno.ProductReplyErr(err)
	}
	reply.Data = list
	return nil
}

//SkuSearch 搜索
func (p *Product) SkuSearch(ctx context.Context, req *pb.SearchReq, reply *pb.SearchReply) error {
	attrs := make(map[int64][]string)
	if len(req.Attrs) > 0 {
		for _, attr := range req.Attrs {
			attrs[attr.Id] = attr.Values
		}
	}
	data, err := p.srv.Search(ctx, req.Keyword, req.CatId, req.Field, req.Order,
		req.PriceS, req.PriceE, req.HasStock, req.BrandId, attrs, req.Page)
	if err != nil {
		return errno.ProductReplyErr(err)
	}
	reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(data).Elem())
	return nil
}

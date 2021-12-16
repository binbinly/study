package service

import (
	"context"

	"github.com/pkg/errors"

	"common/constvar"
	pb "common/proto/product"
	"common/util"
	"product/es"
	"product/idl"
	"product/model"
)

//SkuList sku商品列表
func (s *Service) SkuList(ctx context.Context, catID int64, offset, limit int) ([]*pb.SkuEs, error) {
	list, err := s.productEs.Query(ctx, catID, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.list] offset: %v, limit: %v", offset, limit)
	}
	return idl.TransferSkuList(list), nil
}

//Search 搜索
func (s *Service) Search(ctx context.Context, keyword string, catID int64, field, order, priceS, priceE int32,
	hasStock bool, brandIds []int64, attrs map[int64][]string, page int32) (*pb.SearchReply, error) {
	query := es.NewQuery()
	if keyword != "" {
		query.Keyword(keyword)
	}
	if catID > 0 {
		query.FilterCatID(catID)
	}
	if len(brandIds) > 0 {
		query.FilterBrandIds(brandIds)
	}
	query.FilterSkuPrice(priceS, priceE)
	if len(attrs) > 0 {
		for attrID, values := range attrs {
			query.FilterAttrs(attrID, values)
		}
	}
	res, err := s.productEs.Search(ctx, query, util.GetPageOffset(page), constvar.DefaultLimit, field, order)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.search] get result")
	}
	cats, err := s.parseCats(ctx, res)
	if err != nil {
		return nil, err
	}
	brands, err := s.parseBrands(ctx, res)
	if err != nil {
		return nil, err
	}

	return idl.TransferSearchRes(&idl.TransferSearchInput{
		Result: res,
		Cats:   cats,
		Brands: brands,
	}), nil
}

//parseCats 解析分类信息
func (s *Service) parseCats(ctx context.Context, res *es.SearchRes) (map[int64]string, error) {
	//结果集中所有分类id
	var catIds []int64
	for _, bucket := range res.Cats.Buckets {
		catIds = append(catIds, int64(bucket.Key))
	}
	//获取分类名
	names, err := s.repo.GetCategoryNamesByIds(ctx, catIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.search] get cat names")
	}
	return names, nil
}

//parseBrands 解析品牌信息
func (s *Service) parseBrands(ctx context.Context, res *es.SearchRes) (map[int64]*model.Brand, error) {
	//结果集中所有品牌id
	var brandIds []int64
	for _, bucket := range res.Brands.Buckets {
		brandIds = append(brandIds, int64(bucket.Key))
	}
	//获取品牌信息
	brands, err := s.repo.GetBrandsByIds(ctx, brandIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.search] get brand info")
	}
	return brands, nil
}

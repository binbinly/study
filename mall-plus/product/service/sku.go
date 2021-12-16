package service

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"common/errno"
	"common/orm"
	pb "common/proto/product"
	"product/idl"
	"product/model"
)

//SkuDetail sku商品详情
func (s *Service) SkuDetail(ctx context.Context, id int64) (*pb.Sku, error) {
	//并发执行
	g, ctx := errgroup.WithContext(ctx)
	//1, sku基本信息获取
	sku, err := s.repo.GetSkuByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.sku] sku info by id: %v", id)
	}
	if sku == nil || sku.ID == 0 {
		return nil, errno.ErrProductNotFound
	}
	input := &idl.TransferSkuInput{
		Sku: sku,
	}
	g.Go(func() error {
		//1.1，获取当前spu下所有sku
		skus, err := s.repo.GetSkusBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] skus by id: %v", sku.SpuID)
		}
		//1.2 获取库存
		stocks, err := s.getSpuStock(ctx, sku.SpuID, skus)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spuStock")
		}
		input.Stocks = stocks
		input.Skus = skus
		return nil
	})
	g.Go(func() error {
		//2, sku图片信息
		skuImages, err := s.repo.GetImagesBySkuID(ctx, id)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] sku images by id: %v", id)
		}
		input.SkuImages = skuImages
		return nil
	})
	g.Go(func() error {
		//3, sku分类下的属性分组
		attrGroups, err := s.repo.GetAttrGroupByCatID(ctx, sku.CatID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] attr groups by id: %v", sku.CatID)
		}
		input.AttrGroups = attrGroups
		return nil
	})
	g.Go(func() error {
		//4, 获取spu的详情
		spu, err := s.repo.GetSpuByID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spu info by id: %v", sku.SpuID)
		}
		if spu == nil || spu.ID == 0 {
			return errno.ErrProductNotFound
		}
		input.Spu = spu
		return nil
	})
	g.Go(func() error {
		//5, 获取sku的销售属性组合
		skuAttrs, err := s.repo.GetSkuAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spu skus by id: %v", sku.SpuID)
		}
		input.SkuAttrs = skuAttrs
		return nil
	})
	g.Go(func() error {
		//6, 获取spu介绍
		spuDesc, err := s.repo.GetSpuDescBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spu desc by id: %v", sku.SpuID)
		}
		input.SpuDesc = spuDesc
		return nil
	})
	g.Go(func() error {
		//7, 获取spu的规格参数
		spuAttrs, err := s.repo.GetAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spu attrs by id: %v", sku.SpuID)
		}
		input.SpuAttrs = spuAttrs
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return idl.TransferSku(input), nil
}

//GetSkuSaleAttrs 获取sku销售属性
func (s *Service) GetSkuSaleAttrs(ctx context.Context, id int64) (*pb.SkuSaleAttr, error) {
	//并发执行
	g, ctx := errgroup.WithContext(ctx)
	// sku基本信息获取
	sku, err := s.repo.GetSkuByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.sku] sku info by id: %v", id)
	}
	if sku == nil || sku.ID == 0 {
		return nil, errno.ErrProductNotFound
	}
	input := &idl.TransferSkuInput{
		Sku: sku,
	}
	g.Go(func() error {
		// 获取当前spu下所有sku
		skus, err := s.repo.GetSkusBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] skus by id: %v", sku.SpuID)
		}
		// 获取库存
		stocks, err := s.getSpuStock(ctx, sku.SpuID, skus)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spuStock")
		}
		input.Stocks = stocks
		input.Skus = skus
		return nil
	})
	g.Go(func() error {
		// 获取spu的详情
		spu, err := s.repo.GetSpuByID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spu info by id: %v", sku.SpuID)
		}
		if spu == nil || spu.ID == 0 {
			return errno.ErrProductNotFound
		}
		input.Spu = spu
		return nil
	})
	g.Go(func() error {
		//5, 获取spu的销售属性组合
		skuAttrs, err := s.repo.GetSkuAttrsBySpuID(ctx, sku.SpuID)
		if err != nil {
			return errors.Wrapf(err, "[service.sku] spu skus by id: %v", sku.SpuID)
		}
		input.SkuAttrs = skuAttrs
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return idl.TransferSkuSaleAttrs(input), nil
}

//GetSkuInfo 获取sku信息
func (s *Service) GetSkuInfo(ctx context.Context, id int64) (*pb.SkuInfo, error) {
	//1, sku基本信息获取
	sku, err := s.repo.GetSkuByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.sku] sku info by id: %v", id)
	}
	if sku == nil || sku.ID == 0 {
		return nil, errno.ErrProductNotFound
	}
	return idl.TransferSkuInfo(sku), nil
}

//SpuComment 商品评价
func (s *Service) SpuComment(ctx context.Context, skuIds []int64, memberID, orderID int64, star int8, content, resources string) error {
	//sku基本信息获取
	skus, err := s.repo.GetSkusByIds(ctx, skuIds)
	if err != nil {
		return errors.Wrapf(err, "[service.sku] sku info by ids: %v", skuIds)
	}
	if len(skus) == 0 {
		return errno.ErrProductNotFound
	}
	comments := make([]*model.SpuCommentModel, 0, len(skus))
	for _, sku := range skus {
		comments = append(comments, &model.SpuCommentModel{
			Spu:       orm.Spu{SpuID: sku.SpuID},
			Sku:       orm.Sku{SkuID: sku.ID},
			SkuName:   sku.Name,
			MemberID:  memberID,
			ReplyID:   0,
			OrderID:   orderID,
			Star:      star,
			SkuAttrs:  sku.AttrValue,
			Resources: resources,
			Content:   content,
			Release:   orm.Release{IsRelease: 1},
		})
	}

	if err = s.repo.CreateSpuComment(ctx, comments); err != nil {
		return errors.Wrapf(err, "[service.sku] batch comment")
	}
	return nil
}

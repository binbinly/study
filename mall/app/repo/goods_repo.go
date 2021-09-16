package repo

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"mall/app/model"
)

type IGoods interface {
	GetSkusByGID(ctx context.Context, goodsID int) (list []*model.GoodsSkuModel, err error)
	GetSkuByID(ctx context.Context, id int) (*model.GoodsSkuModel, error)
	GetSkusAttrsByGID(ctx context.Context, goodsID int) (list []*model.GoodsSkuAttrModel, err error)
	GetImagesByGID(ctx context.Context, goodsID int) (list []*model.GoodsImageModel, err error)
	GetAttrByGID(ctx context.Context, goodsID int) (list []*model.GoodsAttrModel, err error)
	CreateGoodsComment(ctx context.Context, tx *gorm.DB, comments []*model.GoodsCommentModel) error

	GoodsList(ctx context.Context, ids []int, keyword, price, field, order string, orderType, offset, limit int) (list []*model.GoodsModel, err error)
	GoodsDetail(ctx context.Context, id int) (*model.GoodsModel, error)
	GoodsCategoryChild(ctx context.Context, id int) ([]int, error)
	GoodsCategoryTree(ctx context.Context) (list []*model.GoodsCategoryTree, err error)
}

//GoodsList 商品列表
func (r *Repo) GoodsList(ctx context.Context, ids []int, keyword, price, field, order string, orderType, offset, limit int) (list []*model.GoodsModel, err error) {
	err = r.db.WithContext(ctx).Scopes(goodsScopesCat(ids),
		goodScopesKeyword(keyword),
		goodsScopesPrice(price),
		goodsScopesOrder(orderType, field, order),
		model.OffsetPage(offset, limit)).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.goods] list ids: %v", ids)
	}
	return
}

//GoodsDetail 商品详情
func (r *Repo) GoodsDetail(ctx context.Context, id int) (*model.GoodsModel, error) {
	goods := new(model.GoodsModel)
	err := r.db.WithContext(ctx).Model(&model.GoodsModel{}).First(&goods, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.goods] by id:%v", id)
	}
	return goods, nil
}

//goodsScopesCat 分类查询
func goodsScopesCat(ids []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) == 0 { //查询不限分类
			return db
		}
		return db.Where("cat_id IN (?)", ids)
	}
}

//goodScopesKeyword 关键字搜索
func goodScopesKeyword(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" {
			return db
		}
		return db.Where("title like ?", "%"+keyword+"%")
	}
}

//goodsScopesPrice 价格筛选
func goodsScopesPrice(price string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if price == "" {
			return db
		}
		s := strings.Split(price, ",")
		if len(s) != 2 {
			return db
		}
		min, err := strconv.Atoi(s[0])
		if err != nil {
			return db
		}
		max, err := strconv.Atoi(s[1])
		if err != nil {
			return db
		}
		if max == 0 {
			return db.Where("price >= ?", min*100)
		}
		return db.Where("price BETWEEN ? AND ?", min*100, max*100)
	}
}

//goodsScopesOrder 商品排序
func goodsScopesOrder(orderType int, field, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderType == model.GoodsOrderHot {
			return db.Order("sale_count DESC")
		} else if orderType == model.GoodsOrderNew {
			return db.Order(model.DefaultOrder)
		} else if field != "" {
			return db.Order(field + " " + order)
		}
		return db.Order(model.DefaultOrderSort)
	}
}

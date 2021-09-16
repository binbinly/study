package repo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"mall/app/model"
)

//CreateGoodsComment 创建商品评价
func (r *Repo) CreateGoodsComment(ctx context.Context, tx *gorm.DB, comments []*model.GoodsCommentModel) error {
	err := tx.WithContext(ctx).Create(&comments).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.goodsComment] batch create")
	}
	return nil
}
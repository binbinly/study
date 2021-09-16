package repo

import (
	"context"
	"github.com/pkg/errors"
	"mall/app/model"
)

//GetAttrsValByIds 商品销售属性值列表
func (r *Repo) GetAttrsValByIds(ctx context.Context, ids []int) (list []*model.SkuAttrValModel, err error) {
	err = r.db.WithContext(ctx).Where("id in ?", ids).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.skuAttrVal] db find by ids: %v", ids)
	}
	return
}
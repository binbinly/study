package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"common/orm"
	"product/model"
)

//GetAttrsBySpuID 批量获取spu下所有属性信息
func (r *Repo) GetAttrsBySpuID(ctx context.Context, spuID int64) (list []*model.Attrs, err error) {
	doKey := fmt.Sprintf("attrs:%v", spuID)
	if err = r.QueryCache(ctx, doKey, &list, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).Model(&model.AttrValueModel{}).Select("pms_attr_value.*, g.group_id").
			Joins("left join pms_attr_rel_group as g on pms_attr_value.attr_id=g.attr_id").Where("spu_id=?", spuID).
			Order(orm.DefaultOrderSort).Scan(data).Error; err != nil {
			return errors.Wrap(err, "[repo.attrValue] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.attrValue] query cache")
	}
	return
}

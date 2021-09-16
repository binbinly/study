package repo

import (
	"context"

	"github.com/pkg/errors"

	"mall/app/model"
)

//AppPageData 页面配置数据
func (r *Repo) AppPageData(ctx context.Context, page, catID int) (list []*model.AppSetting, err error) {
	err = r.db.WithContext(ctx).Model(&model.AppSettingModel{}).Where("page=? and cat_id=?", page, catID).Order(model.DefaultOrderSort).Scan(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.setting] home data by catID: %v", catID)
	}
	if len(list) == 0 {
		return []*model.AppSetting{}, nil
	}
	return
}

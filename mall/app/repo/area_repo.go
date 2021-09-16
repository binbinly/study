package repo

import (
	"context"

	"github.com/pkg/errors"

	"mall/app/model"
)

//GetAreaAll 获取三级地区城市
func (r *Repo) GetAreaAll(ctx context.Context) (list []*model.Area, err error) {
	err = r.db.WithContext(ctx).Model(&model.AreaModel{}).Where("level<3").Scan(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.area] scan")
	}
	return
}
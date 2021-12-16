package repo

import (
	"context"

	"github.com/pkg/errors"

	"common/orm"
	"market/model"
)

//GetNoticeList 公告列表
func (r *Repo) GetNoticeList(ctx context.Context, offset, limit int) (list []*model.AppNoticeModel, err error) {
	err = r.DB.WithContext(ctx).Scopes(orm.OffsetPage(offset, limit)).Order(orm.DefaultOrder).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "[repo.notice] db find")
	}
	return
}

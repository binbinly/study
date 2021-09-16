package repo

import (
	"context"

	"github.com/pkg/errors"

	"mall/app/model"
)

//GetNoticeList 公告列表
func (r *Repo) GetNoticeList(ctx context.Context, offset, limit int) (list []*model.AppNoticeModel, err error) {
	err = r.db.WithContext(ctx).Scopes(model.OffsetPage(offset, limit)).Order(model.DefaultOrder).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "[repo.notice] db find")
	}
	return
}
package repo

import (
	"context"

	"github.com/pkg/errors"

	"chat-micro/app/logic/model"
)

//IReport 用户举报
type IReport interface {
	// 创建
	ReportCreate(ctx context.Context, report *model.ReportModel) (id uint32, err error)
	// 是否还有待处理记录
	ReportExistPending(ctx context.Context, targetID uint32) (bool, error)
}

// ReportCreate 创建举报
func (r *Repo) ReportCreate(ctx context.Context, report *model.ReportModel) (id uint32, err error) {
	if err = r.db.WithContext(ctx).Create(&report).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.report] create report err")
	}
	return report.ID, nil
}

// ReportExistPending 待处理记录是否存在
func (r *Repo) ReportExistPending(ctx context.Context, targetID uint32) (bool, error) {
	var c int64
	if err := r.db.WithContext(ctx).Model(&model.ReportModel{}).
		Where("target_id=? && status=?", targetID, model.ReportStatusPending).Count(&c).Error; err != nil {
		return false, errors.Wrapf(err, "[repo.report] exist err")
	}
	return c > 0, nil
}

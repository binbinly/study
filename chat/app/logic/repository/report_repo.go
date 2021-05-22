package repository

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/logic/model"
)

type IReport interface {
	// 创建
	ReportCreate(ctx context.Context, report *model.ReportModel) (id uint32, err error)
	// 是否还有待处理记录
	ReportExistPending(ctx context.Context, targetId uint32) (bool, error)
}

// Create 创建举报
func (r *Repo) ReportCreate(ctx context.Context, report *model.ReportModel) (id uint32, err error) {
	err = r.db.Create(&report).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.report] create report err")
	}
	return report.ID, nil
}

// Exist 待处理记录是否存在
func (r *Repo) ReportExistPending(ctx context.Context, targetId uint32) (bool, error) {
	var c int64
	err := r.db.Model(&model.ReportModel{}).Where("target_id=? && status=?", targetId, model.ReportStatusPending).Count(&c).Error
	if err != nil {
		return false, errors.Wrapf(err, "[repo.report] exist err")
	}
	return c > 0, nil
}

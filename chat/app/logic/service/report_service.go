package service

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/logic/model"
)

//IReport 用户投诉接口
type IReport interface {
	// 用户举报
	ReportCreate(ctx context.Context, UserID, friendID uint32, cType int8, cat, content string) error
}

// ReportCreate 举报好友/群
func (s *Service) ReportCreate(ctx context.Context, UserID, friendID uint32, cType int8, cat, content string) error {
	is, err := s.repo.ReportExistPending(ctx, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.report] exist id:%d", friendID)
	}
	if is { // 已举报过
		return ErrReportExisted
	}
	report := &model.ReportModel{
		UID:        model.UID{UserID: UserID},
		TargetID:   friendID,
		TargetType: cType,
		Content:    content,
		Category:   cat,
		Status:     model.ReportStatusPending,
	}
	_, err = s.repo.ReportCreate(ctx, report)
	if err != nil {
		return errors.Wrapf(err, "[service.report] create report err")
	}
	return nil
}

package service

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/logic/model"
)

type IReport interface {
	// 用户举报
	ReportCreate(ctx context.Context, userId, friendId uint32, cType int8, cat, content string) error
}

// ReportCreate 举报好友/群
func (s *Service) ReportCreate(ctx context.Context, userId, friendId uint32, cType int8, cat, content string) error {
	is, err := s.repo.ReportExistPending(ctx, friendId)
	if err != nil {
		return errors.Wrapf(err, "[service.report] exist id:%d", friendId)
	}
	if is { // 已举报过
		return ErrReportExisted
	}
	report := &model.ReportModel{
		Uid:        model.Uid{UserId: userId},
		TargetId:   friendId,
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

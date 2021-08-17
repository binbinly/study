package chat

import (
	"context"

	"github.com/pkg/errors"

	"chat/app/chat/model"
	"chat/internal/orm"
)

//IUser 用户服务接口
type IUser interface {
	// 搜索用户
	UserSearch(ctx context.Context, keyword string) (users []*model.UserEs, err error)
	// 标签用户列表
	UserTagList(ctx context.Context, UserID uint32) (list []*model.UserTag, err error)
	// 用户举报
	ReportCreate(ctx context.Context, UserID, friendID uint32, cType int8, cat, content string) error
}

// UserSearch 搜索用户
func (s *Service) UserSearch(ctx context.Context, keyword string) (users []*model.UserEs, err error) {
	users, err = s.ec.UserSearch(ctx, keyword)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] search user keyword: %s", keyword)
	}
	if len(users) == 0 {
		return make([]*model.UserEs, 0), nil
	}
	return users, nil
}

// UserTagList 用户标签列表
func (s *Service) UserTagList(ctx context.Context, UserID uint32) (list []*model.UserTag, err error) {
	list, err = s.repo.GetTagsByUserID(ctx, UserID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] UserTagList id:%d", UserID)
	}
	return
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
		UID:        orm.UID{UserID: UserID},
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

package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"chat-micro/app/constvar"
	"chat-micro/app/logic/idl"
	"chat-micro/app/logic/model"
	"chat-micro/app/message"
	"chat-micro/internal/orm"
)

//IApply 申请好友接口
type IApply interface {
	// 申请好友
	ApplyFriend(ctx context.Context, userID, friendID uint32, nickname string, lookMe, lookHim int8) (err error)
	// 我的申请列表
	ApplyMyList(ctx context.Context, userID uint32, offset int) (list []*model.ApplyList, err error)
	// 待处理申请数
	ApplyPendingCount(ctx context.Context, userID uint32) (c int64, err error)
	// 申请处理
	ApplyHandle(ctx context.Context, userID, friendID uint32, nickname string, lookMe, lookHim int8) (err error)
}

// ApplyFriend 添加好友
func (s *Service) ApplyFriend(ctx context.Context, userID, friendID uint32, nickname string, lookMe, lookHim int8) error {
	info, err := s.applyInfo(ctx, userID, friendID)
	if err != nil {
		return err
	}
	if info.ID > 0 && info.Status == model.ApplyStatusPending { // 已存在
		return ErrApplyExisted
	}
	apply := model.ApplyModel{
		UID:      orm.UID{UserID: userID},
		FriendID: friendID,
		Nickname: nickname,
		LookMe:   lookMe,
		LookHim:  lookHim,
	}
	if _, err = s.repo.ApplyCreate(ctx, apply); err != nil {
		return errors.Wrapf(err, "[service.apply] create err")
	}
	// 通知被申请好友
	if err = s.pushUserIds(ctx, []uint32{friendID}, message.EventNotify, &message.Notify{Type: "apply"}); err != nil {
		return errors.Wrapf(err, "[service.apply] push notify err")
	}
	return nil
}

// ApplyMyList 用户申请列表
func (s *Service) ApplyMyList(ctx context.Context, userID uint32, offset int) (list []*model.ApplyList, err error) {
	applyList, err := s.repo.GetApplysByUserID(ctx, userID, offset, constvar.DefaultLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.apply] MyApplyList id:%d", userID)
	}
	// 用户id列表
	userIds := make([]uint32, 0)
	for _, apply := range applyList {
		userIds = append(userIds, apply.UserID)
	}
	if len(userIds) == 0 {
		return make([]*model.ApplyList, 0), nil
	}
	// 批量获取用户信息
	users, err := s.repo.GetUsersByIds(ctx, userIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.apply] users ids:%v", userIds)
	}
	transInput := &idl.TransferApplyInput{
		Apply: applyList,
		Users: users,
	}
	return idl.TransferApplyList(transInput), nil
}

// ApplyPendingCount 待处理申请数量
func (s *Service) ApplyPendingCount(ctx context.Context, userID uint32) (c int64, err error) {
	c, err = s.repo.ApplyPendingCount(ctx, userID)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.apply] pending count id:%d", userID)
	}
	return
}

// ApplyHandle 处理好友申请通过
func (s *Service) ApplyHandle(ctx context.Context, userID, friendID uint32, nickname string, lookMe, lookHim int8) (err error) {
	info, err := s.applyInfo(ctx, friendID, userID)
	if err != nil {
		return err
	}
	if info.ID == 0 || info.Status != model.ApplyStatusPending { // 未找到合法申请
		return ErrApplyNotFound
	}
	// 开启事务
	db := orm.GetDB()
	tx := db.Begin()
	// 我对好友的模型
	u := &model.FriendModel{
		UserID:   userID,
		FriendID: friendID,
		Nickname: nickname,
		LookMe:   lookMe,
		LookHim:  lookHim,
	}
	// 好友对我的模型
	f := &model.FriendModel{
		UserID:   info.UserID,
		FriendID: userID,
		Nickname: info.Nickname,
		LookMe:   info.LookMe,
		LookHim:  info.LookHim,
	}
	if err = s.repo.FriendBatchCreate(ctx, tx, []*model.FriendModel{u, f}); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.apply] insert into friend to friend err")
	}
	// 修改申请状态
	if err = s.repo.ApplyUpdateStatus(ctx, tx, info.ID, info.FriendID); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.apply] update apply status err")
	}
	auth, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.apply] auth user id:%d", userID)
	}
	fAuth, err := s.repo.GetUserByID(ctx, friendID)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.apply] auth user id:%d", friendID)
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.apply] tx commit err")
	}
	//推送消息 -> 好友
	if err = s.pushUserIds(ctx, []uint32{friendID}, message.EventChat, &message.Chat{
		From: &message.Sender{
			ID:     userID,
			Name:   info.Nickname,
			Avatar: auth.Avatar,
		},
		ChatType: model.MessageChatTypeUser,
		Type:     model.MessageTypeSystem,
		Content:  msgFriendCreate,
		T:        time.Now().Unix(),
	}); err != nil {
		return errors.Wrapf(err, "[service.apply] push friend: %v", friendID)
	}
	//推送消息 -> 自己
	if err = s.pushUserIds(ctx, []uint32{userID}, message.EventChat, &message.Chat{
		From: &message.Sender{
			ID:     friendID,
			Name:   nickname,
			Avatar: fAuth.Avatar,
		},
		ChatType: model.MessageChatTypeUser,
		Type:     model.MessageTypeSystem,
		Content:  msgFriendCreate,
		T:        time.Now().Unix(),
	}); err != nil {
		return errors.Wrapf(err, "[service.apply] push self: %v", userID)
	}
	return nil
}

// applyInfo 申请详情
func (s *Service) applyInfo(ctx context.Context, userID, friendID uint32) (apply *model.ApplyModel, err error) {
	apply, err = s.repo.GetApplyByFriendID(ctx, userID, friendID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.apply] info id:%d,fid:%d", userID, friendID)
	}
	return
}

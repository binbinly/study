package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"

	"chat-micro/app/logic/idl"
	"chat-micro/app/logic/model"
	"chat-micro/app/message"
	"chat-micro/pkg/util"
)

//IChat 聊天会话接口
type IChat interface {
	// 聊天回话详情
	ChatDetail(ctx context.Context, userID, id uint32, cType int) (*message.Sender, error)
	// 发送消息
	ChatSend(ctx context.Context, userID, toID uint32, t, cType int, content string, options json.RawMessage) (*message.Chat, error)
	// 撤回消息
	ChatRecall(ctx context.Context, userID, toID uint32, cType int, id string) error
}

// ChatDetail 会话详情
func (s *Service) ChatDetail(ctx context.Context, userID, id uint32, cType int) (*message.Sender, error) {
	if cType == model.MessageChatTypeUser {
		return s.detailUser(ctx, userID, id)
	} else if cType == model.MessageChatTypeGroup {
		return s.detailGroup(ctx, userID, id)
	}
	return nil, nil
}

// ChatSend 发送消息
func (s *Service) ChatSend(ctx context.Context, userID, toID uint32, t, cType int, content string, options json.RawMessage) (*message.Chat, error) {
	if cType == model.MessageChatTypeUser {
		return s.sendUser(ctx, userID, toID, t, content, options)
	} else if cType == model.MessageChatTypeGroup {
		return s.sendGroup(ctx, userID, toID, t, content, options)
	}
	return nil, nil
}

// ChatRecall 消息撤回
func (s *Service) ChatRecall(ctx context.Context, userID, toID uint32, cType int, id string) (err error) {
	if cType == model.MessageChatTypeUser { // 私聊
		// 发送消息
		if err = s.pushUserIds(ctx, []uint32{toID}, message.EventRecall, &message.Recall{
			ID:       id,
			FromID:   userID,
			ToID:     toID,
			ChatType: model.MessageChatTypeUser,
		}); err != nil {
			return errors.Wrapf(err, "[service.chat] push recall")
		}
		return nil
	} else if cType == model.MessageChatTypeGroup { // 群聊
		userAll, err := s.repo.GroupUserAll(ctx, toID)
		if err != nil {
			return errors.Wrapf(err, "[service.chat] group user all id:%d", toID)
		}
		userIds := make([]uint32, 0)
		for _, u := range userAll {
			if u.UserID == userID { // 不需要推送自己
				continue
			}
			userIds = append(userIds, u.UserID)
		}
		// 推送消息
		if err = s.pushUserIds(ctx, userIds, message.EventRecall, &message.Recall{
			ID:       id,
			FromID:   userID,
			ToID:     toID,
			ChatType: model.MessageChatTypeGroup,
		}); err != nil {
			return errors.Wrapf(err, "[service.group] send recall")
		}
		return nil
	}
	return nil
}

// detailUser 好友聊天详情
func (s *Service) detailUser(ctx context.Context, userID, id uint32) (*message.Sender, error) {
	// 好友->我关系详情
	f, err := s.repo.GetFriendInfo(ctx, id, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] friend id:%d, fid:%d", userID, id)
	}
	// 已经是好友，并且对方没有拉黑你
	if f.ID == 0 || f.IsBlack == 1 {
		return nil, ErrFriendNotFound
	}
	// 获取我备注的好友昵称
	mf, err := s.repo.GetFriendInfo(ctx, userID, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] my friend id:%d, fid:%d", userID, id)
	}
	if mf.ID == 0 {
		return nil, ErrFriendNotFound
	}
	// 好友用户详情
	u, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] user id: %d", id)
	}
	return &message.Sender{
		ID:     u.ID,
		Name:   mf.Nickname,
		Avatar: u.Avatar,
	}, nil
}

// detailGroup 好友聊天详情
func (s *Service) detailGroup(ctx context.Context, userID, id uint32) (*message.Sender, error) {
	group, err := s.repo.GetGroupByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] group id: %d", id)
	}
	if group.ID == 0 {
		return nil, ErrGroupNotFound
	}
	is, err := s.repo.GroupUserIsJoin(ctx, userID, id)
	if err != nil {
		return nil, err
	}
	if !is {
		return nil, ErrGroupUserNotJoin
	}
	return &message.Sender{
		ID:     group.ID,
		Name:   group.Name,
		Avatar: group.Avatar,
	}, nil
}

// sendUser 发送单聊消息
func (s *Service) sendUser(ctx context.Context, userID, toID uint32, t int, content string, options json.RawMessage) (*message.Chat, error) {
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, toID, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] friend id:%d, fid:%d", userID, toID)
	}
	// 已经是好友，并且对方没有拉黑你
	if f.ID == 0 || f.IsBlack == 1 {
		return nil, ErrFriendNotFound
	}
	// 我的用户详情
	u, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] user id: %d", userID)
	}
	//构建消息
	id, err := util.GenShortID()
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] gen user short id")
	}
	m := &message.Chat{
		ID: id,
		From: &message.Sender{
			ID:     u.ID,
			Name:   f.Nickname,
			Avatar: u.Avatar,
		},
		ChatType: model.MessageChatTypeUser,
		Type:     t,
		Content:  content,
		Options:  options,
		T:        time.Now().Unix(),
	}
	// 推送消息
	if err = s.pushUserIds(ctx, []uint32{toID}, message.EventChat, m); err != nil {
		return nil, errors.Wrapf(err, "[service.chat] send")
	}
	return m, nil
}

// sendGroup 发送群聊消息
func (s *Service) sendGroup(ctx context.Context, userID, toID uint32, t int, content string, options json.RawMessage) (*message.Chat, error) {
	group, err := s.repo.GetGroupByID(ctx, toID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] group info err id:%d", toID)
	}
	if group.ID == 0 {
		return nil, ErrGroupNotFound
	}
	userAll, err := s.repo.GroupUserAll(ctx, toID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] group user all id:%d", toID)
	}
	//验证是否还是群内成员
	if in := inGroup(userAll, userID); !in {
		return nil, ErrGroupUserNotJoin
	}
	//我的用户详情
	u, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] user id: %d", userID)
	}
	//构建消息
	id, err := util.GenShortID()
	if err != nil {
		return nil, errors.Wrapf(err, "[service.chat] gen group short id")
	}
	m := &message.Chat{
		ID: id,
		From: &message.Sender{
			ID:     u.ID,
			Name:   idl.TransferUserName(u),
			Avatar: u.Avatar,
		},
		To: &message.Sender{
			ID:     group.ID,
			Name:   group.Name,
			Avatar: group.Avatar,
		},
		ChatType: model.MessageChatTypeGroup,
		Type:     t,
		Content:  content,
		Options:  options,
		T:        time.Now().Unix(),
	}
	userIds := make([]uint32, 0)
	for _, user := range userAll {
		if user.UserID == userID { // 当前用户消息返回，不用推送
			if user.Nickname != "" { //设置了群昵称
				m.From.Name = user.Nickname
			}
			continue
		}
		userIds = append(userIds, user.UserID)
	}
	// 推送消息
	if err = s.pushUserIds(ctx, userIds, message.EventChat, m); err != nil {
		return nil, errors.Wrapf(err, "[service.group] send")
	}
	return m, nil
}

//是否在群组内
func inGroup(users []*model.GroupUserModel, uid uint32) bool {
	for _, user := range users {
		if user.UserID == uid {
			return true
		}
	}
	return false
}
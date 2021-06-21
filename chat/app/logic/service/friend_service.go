package service

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"chat/app/logic/idl"
	"chat/app/logic/model"
	"chat/pkg/utils"
)

//IFriend 好友服务接口
type IFriend interface {
	// 好友信息
	FriendInfo(ctx context.Context, userID, friendID uint32) (friend *model.FriendInfo, err error)
	// 我的好友
	FriendMyAll(ctx context.Context, userID uint32) (list []*model.UserBase, err error)
	// 我的指定好友
	FriendMyListByIds(ctx context.Context, userID uint32, ids []uint32) (list []*model.UserBase, err error)
	// 我的标签好友
	FriendMyListByTagID(ctx context.Context, userID, tagID uint32) (list []*model.UserBase, err error)
	// 设置黑名单
	FriendSetBlack(ctx context.Context, userID, friendID uint32, isBlack int8) error
	// 设置星标
	FriendSetStar(ctx context.Context, userID, friendID uint32, isStar int8) error
	// 设置朋友圈权限
	FriendSetMomentAuth(ctx context.Context, userID, friendID uint32, me, him int8) error
	// 设置备注标签
	FriendSetRemarkTag(ctx context.Context, userID, friendID uint32, nickname string, tags []string) error
	// 删除好友
	FriendDestroy(ctx context.Context, userID, friendID uint32) error
}

// FriendInfo 好友信息
func (s *Service) FriendInfo(ctx context.Context, userID, friendID uint32) (info *model.FriendInfo, err error) {
	// 好友用户详情
	u, err := s.repo.GetUserByID(ctx, friendID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] user id: %d", friendID)
	}
	if u.ID == 0 {
		return nil, ErrUserNotFound
	}
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, userID, friendID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] friend id:%d, fid:%d", userID, friendID)
	}
	tags := make([]string, 0)
	if f.Tags != "" {
		tags, err = s.repo.GetTagNamesByIds(ctx, userID, utils.SliceTouInt32(strings.Split(f.Tags, ",")))
		if err != nil {
			return nil, errors.Wrapf(err, "[service.friend] tag names err tags:%v", f.Tags)
		}
	}
	transInput := &idl.TransferFriendInput{
		User:       u,
		Friend:     f,
		FriendTags: tags,
	}
	return idl.TransferFriend(transInput), nil
}

// FriendMyAll 我的好友列表
func (s *Service) FriendMyAll(ctx context.Context, userID uint32) (list []*model.UserBase, err error) {
	l, err := s.repo.GetFriendAll(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] list uid:%d", userID)
	}
	return s.friendUserList(ctx, l)
}

// FriendMyListByIds 我的好友选中列表
func (s *Service) FriendMyListByIds(ctx context.Context, userID uint32, ids []uint32) (list []*model.UserBase, err error) {
	l, err := s.repo.GetFriendsByIds(ctx, userID, ids)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] list uid:%d, ids:%v", userID, ids)
	}
	return s.friendUserList(ctx, l)
}

// FriendMyListByTagID 我的标签好友
func (s *Service) FriendMyListByTagID(ctx context.Context, userID, tagID uint32) (list []*model.UserBase, err error) {
	l, err := s.repo.GetFriendsByTagID(ctx, userID, tagID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] tag list id:%d, tagID:%v", userID, tagID)
	}
	return s.friendUserList(ctx, l)
}

// friendUserList 好友用户信息列表
func (s *Service) friendUserList(ctx context.Context, friends []*model.FriendModel) (list []*model.UserBase, err error) {
	if len(friends) == 0 {
		return make([]*model.UserBase, 0), nil
	}
	// 好友id列表
	userIds := make([]uint32, 0)
	for _, f := range friends {
		userIds = append(userIds, f.FriendID)
	}
	// 批量获取用户信息
	users, err := s.repo.GetUsersByIds(ctx, userIds)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] get user ids:%v", userIds)
	}
	list = make([]*model.UserBase, 0)
	for _, u := range users {
		list = append(list, idl.TransferUserBase(u))
	}
	return list, nil
}

// FriendSetBlack 设置加入/移除黑名单
func (s *Service) FriendSetBlack(ctx context.Context, userID, friendID uint32, isBlack int8) error {
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, userID, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.friend] friend id:%d, fid:%d", userID, friendID)
	}
	if f.ID == 0 {
		return ErrFriendNotRecord
	}
	f.IsBlack = isBlack
	return s.repo.FriendSave(ctx, f)
}

// FriendSetStar 设置加入/移除星标
func (s *Service) FriendSetStar(ctx context.Context, userID, friendID uint32, isStar int8) error {
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, userID, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.friend] friend id:%d, fid:%d", userID, friendID)
	}
	if f.ID == 0 {
		return ErrFriendNotRecord
	}
	f.IsStar = isStar
	return s.repo.FriendSave(ctx, f)
}

// FriendSetMomentAuth 设置朋友圈权限
func (s *Service) FriendSetMomentAuth(ctx context.Context, userID, friendID uint32, me, him int8) error {
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, userID, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.friend] friend id:%d, fid:%d", userID, friendID)
	}
	if f.ID == 0 {
		return ErrFriendNotRecord
	}
	f.LookMe = me
	f.LookHim = him
	return s.repo.FriendSave(ctx, f)
}

// FriendSetRemarkTag 设置备注标签
func (s *Service) FriendSetRemarkTag(ctx context.Context, userID, friendID uint32, nickname string, tags []string) error {
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, userID, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.friend] friend id:%d, fid:%d", userID, friendID)
	}
	if f.ID == 0 {
		return ErrFriendNotRecord
	}
	if len(tags) > 0 {
		tagIds, err := s.getTagIds(ctx, userID, tags)
		if err != nil {
			return err
		}
		f.Tags = utils.SliceUInt32ToString2(tagIds)
	}
	f.Nickname = nickname
	return s.repo.FriendSave(ctx, f)
}

// FriendDestroy 删除好友
func (s *Service) FriendDestroy(ctx context.Context, userID, friendID uint32) error {
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, userID, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.friend] friend id:%d, fid:%d", userID, friendID)
	}
	if f.ID == 0 {
		return ErrFriendNotRecord
	}
	return s.repo.FriendDelete(ctx, f)
}

// getTagIds 获取标签id列表
func (s *Service) getTagIds(ctx context.Context, userID uint32, tags []string) (tagIds []uint32, err error) {
	// 获取我的所有标签
	myTags, err := s.repo.GetTagsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] tag all id:%d", userID)
	}
	newTags := make([]*model.UserTagModel, 0)
	for _, tag := range tags {
		var id uint32
		for _, myTag := range myTags {
			if myTag.Name == tag { // 该标签已存在
				id = myTag.ID
				break
			}
		}
		if id == 0 {
			newTags = append(newTags, &model.UserTagModel{
				UserID: userID,
				Name:   tag,
			})
		} else {
			tagIds = append(tagIds, id)
		}
	}
	if len(newTags) > 0 {
		// 新标签批量入库
		ids, err := s.repo.TagBatchCreate(ctx, newTags)
		if err != nil {
			return nil, errors.Wrapf(err, "[service.firned] batch create")
		}
		tagIds = append(tagIds, ids...)
	}
	return tagIds, nil
}

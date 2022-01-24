package repo

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat-micro/app/logic/model"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/util"
)

//IFriend 好友接口
type IFriend interface {
	// 创建
	FriendCreate(ctx context.Context, tx *gorm.DB, friend *model.FriendModel) (err error)
	// 批量创建
	FriendBatchCreate(ctx context.Context, tx *gorm.DB, friends []*model.FriendModel) (err error)
	// 好友信息
	GetFriendInfo(ctx context.Context, userID, friendID uint32) (friend *model.FriendModel, err error)
	// 我的全部好友
	GetFriendAll(ctx context.Context, userID uint32) (list []*model.FriendModel, err error)
	// 我的指定好友
	GetFriendsByIds(ctx context.Context, userID uint32, ids []uint32) (list []*model.FriendModel, err error)
	// 我的标签好友
	GetFriendsByTagID(ctx context.Context, userID uint32, tagID uint32) (list []*model.FriendModel, err error)
	// 保存信息
	FriendSave(ctx context.Context, friend *model.FriendModel) error
	// 删除好友
	FriendDelete(ctx context.Context, friend *model.FriendModel) error
}

// FriendCreate 创建好友关系
func (r *Repo) FriendCreate(ctx context.Context, tx *gorm.DB, friend *model.FriendModel) (err error) {
	if err = tx.WithContext(ctx).Create(&friend).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] create err")
	}
	r.delFriendAllCache(ctx, friend.UserID)
	r.delFriendCache(ctx, friend.UserID, friend.FriendID)
	return err
}

//FriendBatchCreate 批量创建
func (r *Repo) FriendBatchCreate(ctx context.Context, tx *gorm.DB, friends []*model.FriendModel) (err error) {
	if err = tx.WithContext(ctx).Model(&model.FriendModel{}).Create(&friends).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] batch create err")
	}
	for _, friend := range friends {
		r.delFriendAllCache(ctx, friend.UserID)
		r.delFriendCache(ctx, friend.UserID, friend.FriendID)
	}
	return err
}

//GetFriendInfo 好友信息
func (r *Repo) GetFriendInfo(ctx context.Context, userID, friendID uint32) (friend *model.FriendModel, err error) {
	if err = r.queryCache(ctx, friendCacheKey(userID, friendID), &friend, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Where("user_id=? && friend_id=?", userID, friendID).
			First(data).Error; err != nil {
			return errors.Wrap(err, "[repo.friend] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.friend] query cache")
	}
	return
}

// GetFriendAll 好友列表
func (r *Repo) GetFriendAll(ctx context.Context, userID uint32) (list []*model.FriendModel, err error) {
	if err = r.queryCache(ctx, friendAllCacheKey(userID), &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.FriendModel{}).
			Where("user_id=? and is_black=0", userID).Limit(5000).Find(data).Error; err != nil {
			return errors.Wrap(err, "[repo.friend] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.friend] query cache")
	}
	return
}

// GetFriendsByIds 获取指定的好友列表
func (r *Repo) GetFriendsByIds(ctx context.Context, userID uint32, ids []uint32) (list []*model.FriendModel, err error) {
	l, err := r.GetFriendAll(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.friend] get friend list by ids")
	}
	list = make([]*model.FriendModel, 0)
	for _, friend := range l {
		if util.InuInt32Slice(friend.FriendID, ids) {
			list = append(list, friend)
		}
	}
	return list, nil
}

// GetFriendsByTagID 获取指定标签好友列表
func (r *Repo) GetFriendsByTagID(ctx context.Context, userID, tagID uint32) (list []*model.FriendModel, err error) {
	l, err := r.GetFriendAll(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.friend] get friend list by ids")
	}
	list = make([]*model.FriendModel, 0)
	for _, friend := range l {
		if friend.Tags != "" {
			tags := strings.Split(friend.Tags, ",")
			if util.InStringSlice(strconv.Itoa(int(tagID)), tags) {
				list = append(list, friend)
			}
		}
	}
	return list, nil
}

// FriendSave 保存好友信息
func (r *Repo) FriendSave(ctx context.Context, friend *model.FriendModel) error {
	if err := r.db.WithContext(ctx).Save(friend).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] save err")
	}
	r.delFriendAllCache(ctx, friend.UserID)
	r.delFriendCache(ctx, friend.UserID, friend.FriendID)
	return nil
}

// FriendDelete 删除记录
func (r *Repo) FriendDelete(ctx context.Context, friend *model.FriendModel) error {
	if err := r.db.WithContext(ctx).Delete(friend).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] delete err")
	}
	r.delFriendAllCache(ctx, friend.UserID)
	r.delFriendCache(ctx, friend.UserID, friend.FriendID)
	return nil
}

//delFriendCache 删除缓存
func (r *Repo) delFriendCache(ctx context.Context, uid, fid uint32) {
	if err := r.cache.Del(ctx, friendCacheKey(uid, fid)); err != nil {
		logger.Warnf("[repo.friend] del cache key: %v", friendCacheKey(uid, fid))
	}
}

//delFriendAllCache 删除缓存
func (r *Repo) delFriendAllCache(ctx context.Context, id uint32) {
	if err := r.cache.Del(ctx, friendAllCacheKey(id)); err != nil {
		logger.Warnf("[repo.friend] del cache key: %v", friendAllCacheKey(id))
	}
}

func friendAllCacheKey(uid uint32) string {
	return fmt.Sprintf("friend:all:%d", uid)
}

func friendCacheKey(uid, fid uint32) string {
	return fmt.Sprintf("friend:%d_%d", uid, fid)
}

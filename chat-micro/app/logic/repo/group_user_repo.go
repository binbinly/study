package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat-micro/app/logic/model"
	"chat-micro/pkg/logger"
)

//IGroupUser 群成员数据仓库
type IGroupUser interface {
	// 创建群组
	GroupUserCreate(ctx context.Context, user *model.GroupUserModel) (err error)
	// 批量创建
	GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, users []*model.GroupUserModel) (err error)
	// 修改群成员昵称
	GroupUserUpdateNickname(ctx context.Context, userID, groupID uint32, nickname string) error
	// 删除成员
	GroupUserDelete(ctx context.Context, user *model.GroupUserModel) error
	// 删除群所有成员
	GroupUserDeleteByGroupID(ctx context.Context, tx *gorm.DB, groupID uint32) error
	// 获取成员信息
	GetGroupUserByID(ctx context.Context, userID, groupID uint32) (info *model.GroupUserModel, err error)
	// 获取所有群成员
	GroupUserAll(ctx context.Context, groupID uint32) (all []*model.GroupUserModel, err error)
	// 是否是群组成员
	GroupUserIsJoin(ctx context.Context, userID, groupID uint32) (bool, error)
}

// GroupUserCreate 创建群成员
func (r *Repo) GroupUserCreate(ctx context.Context, user *model.GroupUserModel) (err error) {
	if err = r.db.WithContext(ctx).Create(user).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] create err")
	}
	r.delGroupUserAllCache(ctx, user.GroupID)
	return nil
}

// GroupUserBatchCreate 批量创建群成员
func (r *Repo) GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, users []*model.GroupUserModel) (err error) {
	if err = tx.WithContext(ctx).Create(&users).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] multi create err")
	}
	return nil
}

// GroupUserUpdateNickname 修改群昵称
func (r *Repo) GroupUserUpdateNickname(ctx context.Context, userID, groupID uint32, nickname string) error {
	if err := r.db.WithContext(ctx).Model(&model.GroupUserModel{}).
		Where("user_id=? && group_id=?", userID, groupID).
		Update("nickname", nickname).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] update nickname err")
	}
	r.delGroupUserAllCache(ctx, groupID)
	r.delGroupUserCache(ctx, userID, groupID)
	return nil
}

// GroupUserDelete 删除群成员
func (r *Repo) GroupUserDelete(ctx context.Context, user *model.GroupUserModel) error {
	if err := r.db.WithContext(ctx).Delete(user).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] quit group err")
	}
	r.delGroupUserAllCache(ctx, user.GroupID)
	return nil
}

// GroupUserDeleteByGroupID 删除群下所有成员
func (r *Repo) GroupUserDeleteByGroupID(ctx context.Context, tx *gorm.DB, groupID uint32) (err error) {
	if err = tx.WithContext(ctx).Where("group_id=?", groupID).Delete(&model.GroupUserModel{}).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] delete users err")
	}
	r.delGroupUserAllCache(ctx, groupID)
	return nil
}

// GetGroupUserByID 获取群成员信息
func (r *Repo) GetGroupUserByID(ctx context.Context, userID, groupID uint32) (info *model.GroupUserModel, err error) {
	if err = r.queryCache(ctx, groupUserCacheKey(userID, groupID), &info, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Where("user_id=? and group_id=?", userID, groupID).
			First(data).Error; err != nil {
			return errors.Wrap(err, "[repo.group_user] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group_user] query cache")
	}
	return
}

// GroupUserAll 获取群所有成员
func (r *Repo) GroupUserAll(ctx context.Context, groupID uint32) (list []*model.GroupUserModel, err error) {
	if err = r.queryCache(ctx, groupUserAllCacheKey(groupID), &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.GroupUserModel{}).
			Where("group_id=?", groupID).Find(data).Error; err != nil {
			return errors.Wrap(err, "[repo.group_user] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group_user] query cache")
	}
	return
}

// GroupUserIsJoin 是否为群成员
func (r *Repo) GroupUserIsJoin(ctx context.Context, userID, groupID uint32) (bool, error) {
	list, err := r.GroupUserAll(ctx, groupID)
	if err != nil {
		return false, errors.Wrapf(err, "[repo.group_user] get all err")
	}
	for _, user := range list {
		if user.UserID == userID {
			return true, nil
		}
	}
	return false, nil
}

//delGroupUserCache 删除缓存
func (r *Repo) delGroupUserCache(ctx context.Context, uid, gid uint32) {
	if err := r.cache.Del(ctx, groupUserCacheKey(uid, gid)); err != nil {
		logger.Warnf("[repo.group_user] del cache key: %v", groupUserCacheKey(uid, gid))
	}
}

//delGroupUserAllCache 删除缓存
func (r *Repo) delGroupUserAllCache(ctx context.Context, id uint32) {
	if err := r.cache.Del(ctx, groupUserAllCacheKey(id)); err != nil {
		logger.Warnf("[repo.group_user] del cache key: %v", groupUserAllCacheKey(id))
	}
}

func groupUserCacheKey(uid, gid uint32) string {
	return fmt.Sprintf("group_user:%d_%d", uid, gid)
}

func groupUserAllCacheKey(id uint32) string {
	return fmt.Sprintf("group_user:all:%d", id)
}

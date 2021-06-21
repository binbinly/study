package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"chat/app/logic/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
)

//IGroupUser 群成员数据仓库
type IGroupUser interface {
	// 创建群组
	GroupUserCreate(ctx context.Context, user *model.GroupUserModel) (err error)
	// 批量创建
	GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, userID, groupID uint32, friends []*model.UserModel) (err error)
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
	err = r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] create err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, user.GroupID)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete all cache")
	}
	return nil
}

// GroupUserBatchCreate 批量创建群成员
func (r *Repo) GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, userID, groupID uint32, friends []*model.UserModel) (err error) {
	users := make([]model.GroupUserModel, 0)
	// 自己加入群组
	users = append(users, model.GroupUserModel{
		UID:     model.UID{UserID: userID},
		GroupID: groupID,
	})
	for _, friend := range friends {
		users = append(users, model.GroupUserModel{
			UID:     model.UID{UserID: friend.ID},
			GroupID: groupID,
		})
	}
	err = tx.WithContext(ctx).Create(&users).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] multi create err")
	}
	return nil
}

// GroupUserUpdateNickname 修改群昵称
func (r *Repo) GroupUserUpdateNickname(ctx context.Context, userID, groupID uint32, nickname string) error {
	err := r.db.WithContext(ctx).Model(&model.GroupUserModel{}).Where("user_id=? && group_id=?", userID, groupID).Update("nickname", nickname).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] update nickname err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, groupID)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete all cache")
	}
	// 删除缓存
	err = r.groupUserCache.DelCache(ctx, userID, groupID)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete info cache")
	}
	return nil
}

// GroupUserDelete 删除群成员
func (r *Repo) GroupUserDelete(ctx context.Context, user *model.GroupUserModel) error {
	err := r.db.WithContext(ctx).Delete(user).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] quit group err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, user.GroupID)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete all cache")
	}
	return nil
}

// GroupUserDeleteByGroupID 删除群下所有成员
func (r *Repo) GroupUserDeleteByGroupID(ctx context.Context, tx *gorm.DB, groupID uint32) (err error) {
	err = tx.WithContext(ctx).Where("group_id=?", groupID).Delete(&model.GroupUserModel{}).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group] delete users err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, groupID)
	if err != nil {
		log.Warnf("[repo.group] delete all cache err:%v", err)
	}
	return nil
}

// GetGroupUserByID 获取群成员信息
func (r *Repo) GetGroupUserByID(ctx context.Context, userID, groupID uint32) (info *model.GroupUserModel, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.group_user] group user uid: %d gid: %d cost: %d μs", userID, groupID, time.Since(start).Microseconds())
	}()
	// 从cache获取
	info, err = r.groupUserCache.GetCache(ctx, userID, groupID)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new(model.GroupUserModel), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.group_user] get group_user by uid: %d gid: %d", userID, groupID)
		}
	}
	// hit cache
	if info != nil {
		log.Debugf("[repo.group_user] get group_user data from cache, uid: %d gid: %d", userID, groupID)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := new(model.GroupUserModel)
		// 从数据库中获取
		err = r.db.WithContext(ctx).Where("user_id=? and group_id=?", userID, groupID).First(data).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = r.groupUserCache.SetCacheWithNotFound(ctx, userID, groupID)
			if err != nil {
				log.Warnf("[repo.group_user] set cache err, uid: %d gid: %d", userID, groupID)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo.group_user] query db err")
		}

		// set cache
		err = r.groupUserCache.SetCache(ctx, userID, groupID, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.group_user] set cache data err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	doKey := fmt.Sprintf("get_group_user_%d_%d", userID, groupID)
	val, err, _ := gr.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.group_user] get err via single flight do")
	}
	data := val.(*model.GroupUserModel)

	return data, nil
}

// GroupUserAll 获取群所有成员
func (r *Repo) GroupUserAll(ctx context.Context, groupID uint32) (all []*model.GroupUserModel, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.group_user] gid: %d cost: %d μs", groupID, time.Since(start).Microseconds())
	}()
	// 从cache获取
	all, err = r.groupUserCache.GetCacheAll(ctx, groupID)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.GroupUserModel, 0), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.group_user] get list by gid: %d", groupID)
		}
	}
	if len(all) > 0 {
		log.Debugf("[repo.group_user] get from cache, gid: %d", groupID)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.GroupUserModel, 0)
		err = r.db.WithContext(ctx).Model(&model.GroupUserModel{}).Where("group_id=?", groupID).Find(&data).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.group_user] query db err")
		}

		// set cache
		err = r.groupUserCache.SetCacheAll(ctx, groupID, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.group_user] set cache all err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	doKey := fmt.Sprintf("get_group_user_all_%d", groupID)
	val, err, _ := gr.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.group_user] get all err via single flight do")
	}
	data := val.([]*model.GroupUserModel)

	return data, nil
}

// GroupUserIsJoin 是否为群成员
func (r *Repo) GroupUserIsJoin(ctx context.Context, userID, groupID uint32) (bool, error) {
	list, err := r.GroupUserAll(ctx, groupID)
	if err != nil {
		return false, errors.Wrapf(err, "[repo.group_user] get all err")
	}
	for _, userModel := range list {
		if userModel.UserID == userID {
			return true, nil
		}
	}
	return false, nil
}

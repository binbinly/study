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

type IGroupUser interface {
	// 创建群组
	GroupUserCreate(ctx context.Context, user *model.GroupUserModel) (err error)
	// 批量创建
	GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, userId, groupId uint32, friends []*model.UserModel) (err error)
	// 修改群成员昵称
	GroupUserUpdateNickname(ctx context.Context, userId, groupId uint32, nickname string) error
	// 删除成员
	GroupUserDelete(ctx context.Context, user *model.GroupUserModel) error
	// 删除群所有成员
	GroupUserDeleteByGroupId(ctx context.Context, tx *gorm.DB, groupId uint32) error
	// 获取成员信息
	GetGroupUserById(ctx context.Context, userId, groupId uint32) (info *model.GroupUserModel, err error)
	// 获取所有群成员
	GroupUserAll(ctx context.Context, groupId uint32) (all []*model.GroupUserModel, err error)
	// 是否是群组成员
	GroupUserIsJoin(ctx context.Context, userId, groupId uint32) (bool, error)
}

// GroupUserCreate 创建群成员
func (r *Repo) GroupUserCreate(ctx context.Context, user *model.GroupUserModel) (err error) {
	err = r.db.Create(user).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] create err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, user.GroupId)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete all cache")
	}
	return nil
}

// GroupUserBatchCreate 批量创建群成员
func (r *Repo) GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, userId, groupId uint32, friends []*model.UserModel) (err error) {
	users := make([]model.GroupUserModel, 0)
	// 自己加入群组
	users = append(users, model.GroupUserModel{
		Uid:     model.Uid{UserId: userId},
		GroupId: groupId,
	})
	for _, friend := range friends {
		users = append(users, model.GroupUserModel{
			Uid:     model.Uid{UserId: friend.ID},
			GroupId: groupId,
		})
	}
	err = tx.Create(&users).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] multi create err")
	}
	return nil
}

// GroupUserUpdateNickname 修改群昵称
func (r *Repo) GroupUserUpdateNickname(ctx context.Context, userId, groupId uint32, nickname string) error {
	err := r.db.Model(&model.GroupUserModel{}).Where("user_id=? && group_id=?", userId, groupId).Update("nickname", nickname).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] update nickname err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, groupId)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete all cache")
	}
	// 删除缓存
	err = r.groupUserCache.DelCache(ctx, userId, groupId)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete info cache")
	}
	return nil
}

// GroupUserDelete 删除群成员
func (r *Repo) GroupUserDelete(ctx context.Context, user *model.GroupUserModel) error {
	err := r.db.Delete(user).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group_user] quit group err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, user.GroupId)
	if err != nil {
		return errors.Wrap(err, "[repo.group_user] delete all cache")
	}
	return nil
}

// GroupUserDeleteByGroupId 删除群下所有成员
func (r *Repo) GroupUserDeleteByGroupId(ctx context.Context, tx *gorm.DB, groupId uint32) (err error) {
	err = tx.Where("group_id=?", groupId).Delete(&model.GroupUserModel{}).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group] delete users err")
	}
	// 删除缓存
	err = r.groupUserCache.DelCacheAll(ctx, groupId)
	if err != nil {
		log.Errorf("[repo.group] delete all cache err:%v", err)
	}
	return nil
}

// GetGroupUserById 获取群成员信息
func (r *Repo) GetGroupUserById(ctx context.Context, userId, groupId uint32) (info *model.GroupUserModel, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.group_user] group user uid: %d gid: %d cost: %d μs", userId, groupId, time.Since(start).Microseconds())
	}()
	// 从cache获取
	info, err = r.groupUserCache.GetCache(ctx, userId, groupId)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new(model.GroupUserModel), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.group_user] get group_user by uid: %d gid: %d", userId, groupId)
		}
	}
	// hit cache
	if info != nil {
		log.Infof("[repo.group_user] get group_user data from cache, uid: %d gid: %d", userId, groupId)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := new(model.GroupUserModel)
		// 从数据库中获取
		err = r.db.Where("user_id=? and group_id=?", userId, groupId).First(data).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = r.groupUserCache.SetCacheWithNotFound(ctx, userId, groupId)
			if err != nil {
				log.Warnf("[repo.group_user] set cache err, uid: %d gid: %d", userId, groupId)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo.group_user] query db err")
		}

		// set cache
		err = r.groupUserCache.SetCache(ctx, userId, groupId, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.group_user] set cache data err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	doKey := fmt.Sprintf("get_group_user_%d_%d", userId, groupId)
	val, err, _ := gr.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.group_user] get err via single flight do")
	}
	data := val.(*model.GroupUserModel)

	return data, nil
}

// GroupUserAll 获取群所有成员
func (r *Repo) GroupUserAll(ctx context.Context, groupId uint32) (all []*model.GroupUserModel, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.group_user] gid: %d cost: %d μs", groupId, time.Since(start).Microseconds())
	}()
	// 从cache获取
	all, err = r.groupUserCache.GetCacheAll(ctx, groupId)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.GroupUserModel, 0), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.group_user] get list by gid: %d", groupId)
		}
	}
	if len(all) > 0 {
		log.Infof("[repo.group_user] get from cache, gid: %d", groupId)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.GroupUserModel, 0)
		err = r.db.Model(&model.GroupUserModel{}).Where("group_id=?", groupId).Find(&data).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.group_user] query db err")
		}

		// set cache
		err = r.groupUserCache.SetCacheAll(ctx, groupId, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.group_user] set cache all err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	doKey := fmt.Sprintf("get_group_user_all_%d", groupId)
	val, err, _ := gr.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.group_user] get all err via single flight do")
	}
	data := val.([]*model.GroupUserModel)

	return data, nil
}

// GroupUserIsJoin 是否为群成员
func (r *Repo) GroupUserIsJoin(ctx context.Context, userId, groupId uint32) (bool, error) {
	list, err := r.GroupUserAll(ctx, groupId)
	if err != nil {
		return false, errors.Wrapf(err, "[repo.group_user] get all err")
	}
	for _, userModel := range list {
		if userModel.UserId == userId {
			return true, nil
		}
	}
	return false, nil
}

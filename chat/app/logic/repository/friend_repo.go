package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"chat/app/logic/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
	"chat/pkg/utils"
)

type IFriend interface {
	// 创建
	FriendCreate(ctx context.Context, tx *gorm.DB, friend *model.FriendModel) (err error)
	// 好友信息
	GetFriendInfo(ctx context.Context, userId, friendId uint32) (friend *model.FriendModel, err error)
	// 我的全部好友
	GetFriendAll(ctx context.Context, userId uint32) (list []*model.FriendModel, err error)
	// 我的指定好友
	GetFriendsByIds(ctx context.Context, userId uint32, ids []uint32) (list []*model.FriendModel, err error)
	// 我的标签好友
	GetFriendsByTagId(ctx context.Context, userId uint32, tagId uint32) (list []*model.FriendModel, err error)
	// 保存信息
	FriendSave(ctx context.Context, friend *model.FriendModel) error
	// 删除好友
	FriendDelete(ctx context.Context, friend *model.FriendModel) error
}

// FriendCreate 创建好友关系
func (r *Repo) FriendCreate(ctx context.Context, tx *gorm.DB, friend *model.FriendModel) (err error) {
	err = tx.Create(&friend).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.friend] create err")
	}
	// 删除缓存
	err = r.friendCache.DelCacheAll(ctx, friend.UserId)
	if err != nil {
		return errors.Wrap(err, "[repo.friend] delete all cache")
	}
	err = r.friendCache.DelCache(ctx, friend.UserId, friend.FriendId)
	if err != nil {
		return errors.Wrap(err, "[repo.friend] delete info cache")
	}
	return err
}

//GetFriendInfo 好友信息
func (r *Repo) GetFriendInfo(ctx context.Context, userId, friendId uint32) (friend *model.FriendModel, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.friend] get friend by uid: %d fid: %d cost: %d μs", userId, friendId, time.Since(start).Microseconds())
	}()
	// 从cache获取
	friend, err = r.friendCache.GetCache(ctx, userId, friendId)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new(model.FriendModel), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.friend] get friend by uid: %d fid: %d", userId, friendId)
		}
	}
	// hit cache
	if friend != nil {
		log.Infof("[repo.friend] get friend from cache, uid: %d, fid: %d", userId, friendId)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := new(model.FriendModel)

		err = r.db.Where("user_id=? && friend_id=?", userId, friendId).First(data).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = r.friendCache.SetCacheWithNotFound(ctx, userId, friendId)
			if err != nil {
				log.Warnf("[repo.friend] SetCacheWithNotFound err, uid: %d fid: %d", userId, friendId)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo.friend] query db err")
		}

		// set cache
		err = r.friendCache.SetCache(ctx, userId, friendId, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.friend] set cache friend err")
		}
		return data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_friend_%d_%d", userId, friendId)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.friend] get friend err via single flight do")
	}
	data := val.(*model.FriendModel)

	return data, nil
}

// GetFriendAll 好友列表
func (r *Repo) GetFriendAll(ctx context.Context, userId uint32) (list []*model.FriendModel, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.friend] get friend all by uid: %d cost: %d μs", userId, time.Since(start).Microseconds())
	}()
	// 从cache获取
	list, err = r.friendCache.GetCacheAll(ctx, userId)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.FriendModel, 0), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.friend] get friend all cache by uid: %d", userId)
		}
	}
	// hit cache
	if len(list) > 0 {
		log.Infof("[repo.friend] get friend all from cache, uid: %d", userId)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.FriendModel, 0)
		err = r.db.Model(&model.FriendModel{}).Where("user_id=? and is_black=0", userId).Find(&data).Error
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.friend] query db err")
		}

		// set cache
		err = r.friendCache.SetCacheAll(ctx, userId, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.friend] set cache friend all err")
		}
		return data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_friend_all_%d", userId)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.friend] get friend all err via single flight do")
	}
	data := val.([]*model.FriendModel)

	return data, nil
}

// GetFriendsByIds 获取指定的好友列表
func (r *Repo) GetFriendsByIds(ctx context.Context, userId uint32, ids []uint32) (list []*model.FriendModel, err error) {
	l, err := r.GetFriendAll(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.friend] get friend list by ids")
	}
	list = make([]*model.FriendModel, 0)
	for _, friendModel := range l {
		if utils.InuInt32Slice(friendModel.FriendId, ids) {
			list = append(list, friendModel)
		}
	}
	return list, nil
}

// GetFriendsByTagId 获取指定标签好友列表
func (r *Repo) GetFriendsByTagId(ctx context.Context, userId, tagId uint32) (list []*model.FriendModel, err error) {
	l, err := r.GetFriendAll(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.friend] get friend list by ids")
	}
	list = make([]*model.FriendModel, 0)
	for _, friendModel := range l {
		if friendModel.Tags != "" {
			tags := strings.Split(friendModel.Tags, ",")
			if utils.InStringSlice(strconv.Itoa(int(tagId)), tags) {
				list = append(list, friendModel)
			}
		}
	}
	return list, nil
}

// FriendSave 保存好友信息
func (r *Repo) FriendSave(ctx context.Context, friend *model.FriendModel) error {
	err := r.db.Save(friend).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.friend] save err")
	}
	// 删除缓存
	err = r.friendCache.DelCacheAll(ctx, friend.UserId)
	if err != nil {
		return errors.Wrap(err, "[repo.friend] delete all cache")
	}
	// 删除缓存
	err = r.friendCache.DelCache(ctx, friend.UserId, friend.FriendId)
	if err != nil {
		return errors.Wrap(err, "[repo.friend] delete info cache")
	}
	return nil
}

// FriendDelete 删除记录
func (r *Repo) FriendDelete(ctx context.Context, friend *model.FriendModel) error {
	err := r.db.Delete(friend).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.friend] delete err")
	}
	// 删除缓存
	err = r.friendCache.DelCacheAll(ctx, friend.UserId)
	if err != nil {
		return errors.Wrap(err, "[repo.friend] delete all cache")
	}
	err = r.friendCache.DelCache(ctx, friend.UserId, friend.FriendId)
	if err != nil {
		return errors.Wrap(err, "[repo.friend] delete info cache")
	}
	return nil
}

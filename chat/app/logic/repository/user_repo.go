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

//IUser 用户数据仓库
type IUser interface {
	// 创建用户
	UserCreate(ctx context.Context, user *model.UserModel) (id uint32, err error)
	// 修改用户信息
	UserUpdate(ctx context.Context, id uint32, userMap map[string]interface{}) (err error)
	// id获取用户信息
	GetUserByID(ctx context.Context, id uint32) (*model.UserModel, error)
	// 批量获取用户信息
	GetUsersByIds(ctx context.Context, ids []uint32) ([]*model.UserModel, error)
	// 搜索用户
	GetUsersByKeyword(ctx context.Context, keyword string) ([]*model.UserModel, error)
	// username获取用户信息
	GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error)
	// phone获取用户信息
	GetUserByPhone(ctx context.Context, phone int64) (*model.UserModel, error)
	// 用户是否存在
	UserCheckExist(ctx context.Context, username string, phone int64) bool
}

// UserCreate 创建用户
func (r *Repo) UserCreate(ctx context.Context, user *model.UserModel) (id uint32, err error) {
	err = r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.user] Create err")
	}

	return user.ID, nil
}

// UserUpdate 更新用户信息
func (r *Repo) UserUpdate(ctx context.Context, id uint32, userMap map[string]interface{}) (err error) {
	info, err := r.GetUserByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[repo.user] Update err")
	}

	// 删除cache
	err = r.userCache.DelUserCache(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[repo.user] delete cache")
	}

	return r.db.WithContext(ctx).Model(info).Updates(userMap).Error
}

// GetUserByID 获取用户
// 缓存的更新策略使用 Cache Aside Pattern
// see: https://coolshell.cn/articles/17416.html
func (r *Repo) GetUserByID(ctx context.Context, id uint32) (user *model.UserModel, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.user] get user by uid: %d cost: %d μs", id, time.Since(start).Microseconds())
	}()
	// 从cache获取
	user, err = r.userCache.GetUserCache(ctx, id)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new(model.UserModel), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.user] get user by uid: %d", id)
		}
	}
	// hit cache
	if user != nil {
		log.Debugf("[repo.user] get user data from cache, uid: %d", id)
		return
	}

	// use sync/singleflight mode to get data
	// why not use redis lock? see this topic: https://redis.io/topics/distlock
	// demo see: https://github.com/go-demo/singleflight-demo/blob/master/main.go
	// https://juejin.cn/post/6844904084445593613
	getDataFn := func() (interface{}, error) {
		data := new(model.UserModel)
		// 从数据库中获取
		err = r.db.WithContext(ctx).First(data, id).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = r.userCache.SetCacheWithNotFound(ctx, id)
			if err != nil {
				log.Warnf("[repo.user] SetCacheWithNotFound err, uid: %d", id)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo.user] query db err")
		}

		// set cache
		err = r.userCache.SetUserCache(ctx, id, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.user] set cache data err")
		}
		return data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_user_%d", id)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.user] get user err via single flight do")
	}
	data := val.(*model.UserModel)

	return data, nil
}

// GetUsersByIds 批量获取用户
func (r *Repo) GetUsersByIds(ctx context.Context, ids []uint32) ([]*model.UserModel, error) {
	users := make([]*model.UserModel, 0)

	// 从cache批量获取
	userCacheMap, err := r.userCache.MultiGetUserCache(ctx, ids)
	if err != nil {
		return users, errors.Wrap(err, "[repo.user] multi get user cache data err")
	}

	// 查询未命中
	for _, uid := range ids {
		idx := r.userCache.GetUserCacheKey(uid)
		userModel, ok := userCacheMap[idx]
		if !ok {
			userModel, err = r.GetUserByID(ctx, uid)
			if err != nil {
				log.Warnf("[repo.user] get user model err: %v", err)
				continue
			}
		}
		users = append(users, userModel)
	}
	return users, nil
}

// GetUsersByKeyword 搜索用户
func (r *Repo) GetUsersByKeyword(ctx context.Context, keyword string) (users []*model.UserModel, err error) {
	err = r.db.WithContext(ctx).Where("username like ?", keyword+"%").Limit(10).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] search user err by keyword")
	}
	return users, nil
}

// GetUserByUsername 根据账号获取用户
func (r *Repo) GetUserByUsername(ctx context.Context, username string) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	err = r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] get user err by username")
	}
	return user, nil
}

// GetUserByPhone 根据手机号获取用户
func (r *Repo) GetUserByPhone(ctx context.Context, phone int64) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	err = r.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] get user err by phone")
	}
	return user, nil
}

// UserCheckExist 用户是否已存在
func (r *Repo) UserCheckExist(ctx context.Context, username string, phone int64) bool {
	var c int64
	err := r.db.WithContext(ctx).Model(&model.UserModel{}).Where("phone = ? or username=?", phone, username).Count(&c).Error
	if err != nil {
		log.Warnf("[repo.user] user exist err:%v", err)
		return false
	}
	return c > 0
}

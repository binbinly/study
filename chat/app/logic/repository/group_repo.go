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

type IGroup interface {
	// 创建群组
	GroupCreate(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (id uint32, err error)
	// 保存群组
	GroupSave(ctx context.Context, group *model.GroupModel) error
	// 删除群组
	GroupDelete(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (err error)
	// 获取群组信息
	GetGroupById(ctx context.Context, id uint32) (info *model.GroupModel, err error)
	// 获取我的群组列表
	GetGroupsByUserId(ctx context.Context, userId uint32) (list []*model.GroupList, err error)
}

// GroupCreate 创建群组
func (r *Repo) GroupCreate(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (id uint32, err error) {
	err = tx.Create(&group).Error
	if err != nil {
		return 0, errors.Wrapf(err, "[repo.group] create err")
	}
	// 删除缓存
	err = r.groupCache.DelCacheAll(ctx, group.UserId)
	if err != nil {
		return 0, errors.Wrap(err, "[repo.group] delete all cache")
	}
	return group.ID, nil
}

// GroupSave 保存群组信息
func (r *Repo) GroupSave(ctx context.Context, group *model.GroupModel) (err error) {
	err = r.db.Save(group).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group] save err")
	}
	// 删除缓存
	err = r.groupCache.DelCacheAll(ctx, group.UserId)
	if err != nil {
		return errors.Wrap(err, "[repo.group] delete all cache")
	}
	err = r.groupCache.DelCache(ctx, group.ID)
	if err != nil {
		return errors.Wrapf(err, "[repo.group] delete info cache")
	}
	return nil
}

// GroupDelete 删除群
func (r *Repo) GroupDelete(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (err error) {
	err = tx.Delete(group).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.group] delete err")
	}
	// 删除缓存
	err = r.groupCache.DelCacheAll(ctx, group.UserId)
	if err != nil {
		return errors.Wrap(err, "[repo.group] delete all cache")
	}
	err = r.groupCache.DelCache(ctx, group.ID)
	if err != nil {
		log.Warnf("[repo.group] delete info cache err:%v", err)
	}
	return err
}

// GetGroupById 获取群组信息
func (r *Repo) GetGroupById(ctx context.Context, id uint32) (info *model.GroupModel, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.group] id: %d cost: %d μs", id, time.Since(start).Microseconds())
	}()
	// 从cache获取
	info, err = r.groupCache.GetCache(ctx, id)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return new(model.GroupModel), nil
		} else if err != redis.Nil {
			// fail fast, if cache error return, don't request to db
			return nil, errors.Wrapf(err, "[repo.group] get by id: %d", id)
		}
	}
	// hit cache
	if info != nil {
		log.Infof("[repo.group] get data from cache, id: %d", id)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := new(model.GroupModel)
		// 从数据库中获取
		err = r.db.First(data, id).Error
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = r.groupCache.SetCacheWithNotFound(ctx, id)
			if err != nil {
				log.Warnf("[repo.group] set cache err, id: %d", id)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo.group] query db err")
		}

		// set cache
		err = r.groupCache.SetCache(ctx, id, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.group] set cache data err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	doKey := fmt.Sprintf("get_group_%d", id)
	val, err, _ := gr.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.group] get err via single flight do")
	}
	data := val.(*model.GroupModel)

	return data, nil
}

// GetGroupsByUserId 群组列表
func (r *Repo) GetGroupsByUserId(ctx context.Context, userId uint32) (list []*model.GroupList, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.group] uid: %d cost: %d μs", userId, time.Since(start).Microseconds())
	}()
	// 从cache获取
	list, err = r.groupCache.GetCacheAll(ctx, userId)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.GroupList, 0), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.group] get list by uid: %d", userId)
		}
	}
	if len(list) > 0 {
		log.Infof("[repo.group] get from cache, uid: %d", userId)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.GroupList, 0)
		err = r.db.Model(&model.GroupUserModel{}).Distinct().Select("`group`.id, `group`.name, `group`.avatar").
			Joins("left join `group` on `group`.id = group_user.group_id").
			Where("group_user.user_id=?", userId).Scan(&data).Error
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.group] query db err")
		}

		// set cache
		err = r.groupCache.SetCacheAll(ctx, userId, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.group] set cache all err")
		}
		return data, nil
	}

	gr := singleflight.Group{}
	doKey := fmt.Sprintf("get_group_all_%d", userId)
	val, err, _ := gr.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.group] get all err via single flight do")
	}
	data := val.([]*model.GroupList)

	return data, nil
}

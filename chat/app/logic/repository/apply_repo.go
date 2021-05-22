package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"chat/app/logic/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
)

type IApply interface {
	// 创建
	ApplyCreate(ctx context.Context, apply model.ApplyModel) (id uint32, err error)
	// 修改申请状态
	ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, userId, friendId uint32) (err error)
	// 获取申请列表
	GetApplysByUserId(ctx context.Context, userId uint32, offset, limit int) (list []*model.ApplyModel, err error)
	// 申请未处理数
	ApplyPendingCount(ctx context.Context, userId uint32) (c int64, err error)
	// 申请详情
	GetApplyByFriendId(ctx context.Context, userId, friendId uint32) (apply *model.ApplyModel, err error)
}

// ApplyCreate 创建申请记录
func (r *Repo) ApplyCreate(ctx context.Context, apply model.ApplyModel) (id uint32, err error) {
	err = r.db.Create(&apply).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.apply] create err")
	}
	// 删除列表缓存
	err = r.applyCache.DelCacheList(ctx, apply.FriendId)
	if err != nil {
		return 0, errors.Wrap(err, "[repo.apply] delete list cache")
	}

	return apply.ID, nil
}

// ApplyUpdateStatus 修改申请状态
func (r *Repo) ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, userId, friendId uint32) (err error) {
	err = tx.Model(&model.ApplyModel{}).Where("id=? && status=?",
		id, model.ApplyStatusPending).Update("status", model.ApplyStatusAgree).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.apply] update err")
	}
	// 删除缓存
	err = r.applyCache.DelCacheList(ctx, friendId)
	if err != nil {
		log.Errorf("[repo.apply] delete list cache err:%v", err)
	}
	return nil
}

// GetApplysByUserId 获取申请好友列表
func (r *Repo) GetApplysByUserId(ctx context.Context, userId uint32, offset, limit int) (list []*model.ApplyModel, err error) {
	start := time.Now()
	defer func() {
		log.Infof("[repo.apply] GetListByUserId uid: %d offset: %d cost: %d μs", userId, offset, time.Since(start).Microseconds())
	}()
	// 从cache获取
	list, err = r.applyCache.GetCacheList(ctx, userId, strconv.Itoa(offset))
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.ApplyModel, 0), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.apply] get apply list by uid: %d offset:%d", userId, offset)
		}
	}
	if len(list) > 0 {
		log.Infof("[repo.apply] get GetListByUserId from cache, uid: %d offset: %d", userId, offset)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.ApplyModel, 0)
		err = r.db.Scopes(model.OffsetPage(offset, limit)).Where("friend_id = ? ", userId).
			Order(model.DefaultOrder).Find(&data).Error
		if err != nil {
			return nil, errors.Wrap(err, "[repo.apply] get apply list err")
		}

		// set cache
		err = r.applyCache.SetCacheList(ctx, userId, strconv.Itoa(offset), data)
		if err != nil {
			return 0, errors.Wrap(err, "[repo.apply] set cache list err")
		}
		return data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_apply_list_%d_%d", userId, offset)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.apply] get apply list err via single flight do")
	}
	data := val.([]*model.ApplyModel)
	return data, nil
}

// ApplyPendingCount 待处理数量
func (r *Repo) ApplyPendingCount(ctx context.Context, userId uint32) (c int64, err error) {
	err = r.db.Model(&model.ApplyModel{}).Where("friend_id=? && status=?", userId, model.ApplyStatusPending).Count(&c).Error
	if err != nil {
		return 0, errors.Wrapf(err, "[repo.apply] pending count db err, uid: %d", userId)
	}
	return c, nil
}

// GetApplyByFriendId 获取申请详情
func (r *Repo) GetApplyByFriendId(ctx context.Context, userId, friendId uint32) (apply *model.ApplyModel, err error) {
	apply = new(model.ApplyModel)
	err = r.db.Where("user_id=? && friend_id=?", userId, friendId).Order(model.DefaultOrder).First(apply).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(err, "[repo.apply] query db err")
	}
	return apply, nil
}

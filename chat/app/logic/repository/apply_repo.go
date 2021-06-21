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

//IApply 申请好友接口
type IApply interface {
	// 创建
	ApplyCreate(ctx context.Context, apply model.ApplyModel) (id uint32, err error)
	// 修改申请状态
	ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, userID, friendID uint32) (err error)
	// 获取申请列表
	GetApplysByUserID(ctx context.Context, userID uint32, offset, limit int) (list []*model.ApplyModel, err error)
	// 申请未处理数
	ApplyPendingCount(ctx context.Context, userID uint32) (c int64, err error)
	// 申请详情
	GetApplyByFriendID(ctx context.Context, userID, friendID uint32) (apply *model.ApplyModel, err error)
}

// ApplyCreate 创建申请记录
func (r *Repo) ApplyCreate(ctx context.Context, apply model.ApplyModel) (id uint32, err error) {
	err = r.db.WithContext(ctx).Create(&apply).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.apply] create err")
	}
	// 删除列表缓存
	err = r.applyCache.DelCacheList(ctx, apply.FriendID)
	if err != nil {
		return 0, errors.Wrap(err, "[repo.apply] delete list cache")
	}

	return apply.ID, nil
}

// ApplyUpdateStatus 修改申请状态
func (r *Repo) ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, userID, friendID uint32) (err error) {
	err = tx.WithContext(ctx).Model(&model.ApplyModel{}).Where("id=? && status=?",
		id, model.ApplyStatusPending).Update("status", model.ApplyStatusAgree).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.apply] update err")
	}
	// 删除缓存
	err = r.applyCache.DelCacheList(ctx, friendID)
	if err != nil {
		log.Errorf("[repo.apply] delete list cache err:%v", err)
	}
	return nil
}

// GetApplysByUserID 获取申请好友列表
func (r *Repo) GetApplysByUserID(ctx context.Context, userID uint32, offset, limit int) (list []*model.ApplyModel, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.apply] GetListByUserId uid: %d offset: %d cost: %d μs", userID, offset, time.Since(start).Microseconds())
	}()
	// 从cache获取
	list, err = r.applyCache.GetCacheList(ctx, userID, strconv.Itoa(offset))
	if err != nil {
		if err == cache.ErrPlaceholder {
			return make([]*model.ApplyModel, 0), nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.apply] get apply list by uid: %d offset:%d", userID, offset)
		}
	}
	if len(list) > 0 {
		log.Debugf("[repo.apply] get GetListByUserId from cache, uid: %d offset: %d", userID, offset)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.ApplyModel, 0)
		err = r.db.WithContext(ctx).Scopes(model.OffsetPage(offset, limit)).Where("friend_id = ? ", userID).
			Order(model.DefaultOrder).Find(&data).Error
		if err != nil {
			return nil, errors.Wrap(err, "[repo.apply] get apply list err")
		}

		// set cache
		err = r.applyCache.SetCacheList(ctx, userID, strconv.Itoa(offset), data)
		if err != nil {
			return 0, errors.Wrap(err, "[repo.apply] set cache list err")
		}
		return data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_apply_list_%d_%d", userID, offset)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.apply] get apply list err via single flight do")
	}
	data := val.([]*model.ApplyModel)
	return data, nil
}

// ApplyPendingCount 待处理数量
func (r *Repo) ApplyPendingCount(ctx context.Context, userID uint32) (c int64, err error) {
	err = r.db.WithContext(ctx).Model(&model.ApplyModel{}).Where("friend_id=? && status=?", userID, model.ApplyStatusPending).Count(&c).Error
	if err != nil {
		return 0, errors.Wrapf(err, "[repo.apply] pending count db err, uid: %d", userID)
	}
	return c, nil
}

// GetApplyByFriendID 获取申请详情
func (r *Repo) GetApplyByFriendID(ctx context.Context, userID, friendID uint32) (apply *model.ApplyModel, err error) {
	apply = new(model.ApplyModel)
	err = r.db.WithContext(ctx).Where("user_id=? && friend_id=?", userID, friendID).Order(model.DefaultOrder).First(apply).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(err, "[repo.apply] query db err")
	}
	return apply, nil
}

package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat-micro/app/logic/model"
	"chat-micro/pkg/logger"
)

//IMoment 朋友圈
type IMoment interface {
	// 创建一条动态
	MomentCreate(ctx context.Context, tx *gorm.DB, message *model.MomentModel) (id uint32, err error)
	// 我的朋友圈列表
	GetMyMoments(ctx context.Context, userID uint32, offset, limit int) ([]*model.MomentModel, error)
	// 指定好友的朋友圈
	GetMomentsByUserID(ctx context.Context, myID, userID uint32, offset, limit int) ([]*model.MomentModel, error)
	// 获取动态信息
	GetMomentByID(ctx context.Context, id uint32) (*model.MomentModel, error)
	// 批量获取动态信息
	GetMomentsByIds(ctx context.Context, ids []uint32) ([]*model.MomentModel, error)
}

// MomentCreate 创建
func (r *Repo) MomentCreate(ctx context.Context, tx *gorm.DB, moment *model.MomentModel) (id uint32, err error) {
	err = tx.WithContext(ctx).Create(&moment).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.moment] create moment err")
	}
	return moment.ID, nil
}

// GetMyMoments 我的朋友圈列表
func (r *Repo) GetMyMoments(ctx context.Context, userID uint32, offset, limit int) (list []*model.MomentModel, err error) {
	var ids []uint32
	err = r.db.WithContext(ctx).Model(&model.MomentLike{}).Raw("select moment_id from moment_timeline where user_id = ? order by id desc limit ? offset ?", userID, limit, offset).Pluck("moment_id", &ids).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.moment] get db ids err")
	}
	return r.GetMomentsByIds(ctx, ids)
}

// GetMomentsByUserID 指定用户的朋友圈
func (r *Repo) GetMomentsByUserID(ctx context.Context, myID, userID uint32, offset, limit int) (list []*model.MomentModel, err error) {
	if myID == userID { // 查看自己
		err = r.db.WithContext(ctx).Raw("select * from moment where user_id=? order by id desc limit ? offset ?", myID, limit, offset).Find(&list).Error
	} else {
		err = r.db.WithContext(ctx).Raw("select * from moment where user_id=? and (see_type=1 or (see_type = 3 and FIND_IN_SET(?,see)) or (see_type = 4 and !FIND_IN_SET(?,see)) ) order by id desc limit ? offset ?", userID, myID, myID, limit, offset).Find(&list).Error
	}
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.moment] list err")
	}
	return list, nil
}

// GetMomentByID 获取动态信息
func (r *Repo) GetMomentByID(ctx context.Context, id uint32) (moment *model.MomentModel, err error) {
	if err = r.queryCache(ctx, momentCacheKey(id), &moment, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(data, id).Error; err != nil {
			return errors.Wrap(err, "[repo.moment] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment] query cache")
	}
	return
}

// GetMomentsByIds 批量获取动态信息
func (r *Repo) GetMomentsByIds(ctx context.Context, ids []uint32) (moments []*model.MomentModel, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, momentCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.MomentModel)
	if err = r.cache.MultiGet(ctx, keys, cacheMap, func() interface{} {
		return &model.MomentModel{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment] multi get cache data err")
	}

	// 查询未命中
	for _, id := range ids {
		moment, ok := cacheMap[momentCacheKey(id)]
		if !ok {
			moment, err = r.GetMomentByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.moment] get moment model err: %v", err)
				continue
			}
		}
		if moment == nil || moment.ID == 0 {
			continue
		}
		moments = append(moments, moment)
	}
	return moments, nil
}

func momentCacheKey(id uint32) string {
	return fmt.Sprintf("moment:%d", id)
}
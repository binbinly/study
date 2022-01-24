package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat-micro/app/logic/model"
	"chat-micro/internal/orm"
	"chat-micro/pkg/logger"
)

//IApply 申请好友接口
type IApply interface {
	// 创建
	ApplyCreate(ctx context.Context, apply model.ApplyModel) (id uint32, err error)
	// 修改申请状态
	ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, friendID uint32) (err error)
	// 获取申请列表
	GetApplysByUserID(ctx context.Context, userID uint32, offset, limit int) (list []*model.ApplyModel, err error)
	// 申请未处理数
	ApplyPendingCount(ctx context.Context, userID uint32) (c int64, err error)
	// 申请详情
	GetApplyByFriendID(ctx context.Context, userID, friendID uint32) (apply *model.ApplyModel, err error)
}

// ApplyCreate 创建申请记录
func (r *Repo) ApplyCreate(ctx context.Context, apply model.ApplyModel) (id uint32, err error) {
	if err = r.db.WithContext(ctx).Create(&apply).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.apply] create err")
	}
	r.delApplyCache(ctx, apply.FriendID)
	return apply.ID, nil
}

// ApplyUpdateStatus 修改申请状态
func (r *Repo) ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, friendID uint32) (err error) {
	if err = tx.WithContext(ctx).Model(&model.ApplyModel{}).Where("id=? && status=?",
		id, model.ApplyStatusPending).Update("status", model.ApplyStatusAgree).Error; err != nil {
		return errors.Wrapf(err, "[repo.apply] update err")
	}
	r.delApplyCache(ctx, friendID)
	return nil
}

// GetApplysByUserID 获取申请好友列表
func (r *Repo) GetApplysByUserID(ctx context.Context, userID uint32, offset, limit int) (list []*model.ApplyModel, err error) {
	if err = r.queryListCache(ctx, applyCacheKey(userID), offset, &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Scopes(orm.OffsetPage(offset, limit)).Where("friend_id = ? ", userID).
			Order(orm.DefaultOrder).Find(data).Error; err != nil {
			return errors.Wrap(err, "[repo.apply] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.apply] query cache")
	}
	return
}

// ApplyPendingCount 待处理数量
func (r *Repo) ApplyPendingCount(ctx context.Context, userID uint32) (c int64, err error) {
	if err = r.db.WithContext(ctx).Model(&model.ApplyModel{}).
		Where("friend_id=? && status=?", userID, model.ApplyStatusPending).Count(&c).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.apply] pending count db err, uid: %d", userID)
	}
	return c, nil
}

// GetApplyByFriendID 获取申请详情
func (r *Repo) GetApplyByFriendID(ctx context.Context, userID, friendID uint32) (apply *model.ApplyModel, err error) {
	apply = new(model.ApplyModel)
	if err = r.db.WithContext(ctx).Where("user_id=? && friend_id=?", userID, friendID).
		Order(orm.DefaultOrder).First(apply).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(err, "[repo.apply] query db err")
	}
	return apply, nil
}

//delApplyCache 删除缓存
func (r *Repo) delApplyCache(ctx context.Context, uid uint32) {
	if err := r.cache.Del(ctx, applyCacheKey(uid)); err != nil {
		logger.Warnf("[repo.apply] del cache key: %v", applyCacheKey(uid))
	}
}

func applyCacheKey(id uint32) string {
	return fmt.Sprintf("apply:list:%d", id)
}

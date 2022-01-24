package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat-micro/app/logic/model"
	"chat-micro/pkg/logger"
)

//IGroup 群组接口
type IGroup interface {
	// 创建群组
	GroupCreate(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (id uint32, err error)
	// 保存群组
	GroupSave(ctx context.Context, group *model.GroupModel) error
	// 删除群组
	GroupDelete(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (err error)
	// 获取群组信息
	GetGroupByID(ctx context.Context, id uint32) (info *model.GroupModel, err error)
	// 获取我的群组列表
	GetGroupsByUserID(ctx context.Context, userID uint32) (list []*model.GroupList, err error)
}

// GroupCreate 创建群组
func (r *Repo) GroupCreate(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (id uint32, err error) {
	if err = tx.WithContext(ctx).Create(&group).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.group] create err")
	}
	r.delGroupAllCache(ctx, group.UserID)
	return group.ID, nil
}

// GroupSave 保存群组信息
func (r *Repo) GroupSave(ctx context.Context, group *model.GroupModel) (err error) {
	if err = r.db.WithContext(ctx).Save(group).Error; err != nil {
		return errors.Wrapf(err, "[repo.group] save err")
	}
	r.delGroupCache(ctx, group.ID)
	r.delGroupAllCache(ctx, group.UserID)
	return nil
}

// GroupDelete 删除群
func (r *Repo) GroupDelete(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (err error) {
	if err = tx.WithContext(ctx).Delete(group).Error; err != nil {
		return errors.Wrapf(err, "[repo.group] delete err")
	}
	r.delGroupCache(ctx, group.ID)
	r.delGroupAllCache(ctx, group.UserID)
	return err
}

// GetGroupByID 获取群组信息
func (r *Repo) GetGroupByID(ctx context.Context, id uint32) (info *model.GroupModel, err error) {
	if err = r.queryCache(ctx, groupCacheKey(id), &info, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(data, id).Error; err != nil {
			return errors.Wrap(err, "[repo.group] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group] query cache")
	}
	return
}

// GetGroupsByUserID 群组列表
func (r *Repo) GetGroupsByUserID(ctx context.Context, userID uint32) (list []*model.GroupList, err error) {
	if err = r.queryCache(ctx, groupAllCacheKey(userID), &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.GroupUserModel{}).Distinct().Select("`group`.id, `group`.name, `group`.avatar").
			Joins("left join `group` on `group`.id = group_user.group_id").
			Where("group_user.user_id=?", userID).Scan(&data).Error; err != nil {
			return errors.Wrap(err, "[repo.group] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group] query cache")
	}
	return
}

//delGroupCache 删除缓存
func (r *Repo) delGroupCache(ctx context.Context, id uint32) {
	if err := r.cache.Del(ctx, groupCacheKey(id)); err != nil {
		logger.Warnf("[repo.group] del cache key: %v", groupCacheKey(id))
	}
}

//delGroupAllCache 删除缓存
func (r *Repo) delGroupAllCache(ctx context.Context, id uint32) {
	if err := r.cache.Del(ctx, groupAllCacheKey(id)); err != nil {
		logger.Warnf("[repo.group] del cache key: %v", groupAllCacheKey(id))
	}
}

func groupCacheKey(id uint32) string {
	return fmt.Sprintf("group:%d", id)
}

func groupAllCacheKey(uid uint32) string {
	return fmt.Sprintf("group:all:%d", uid)
}

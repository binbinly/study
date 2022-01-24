package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"chat-micro/app/logic/model"
	"chat-micro/internal/orm"
	"chat-micro/pkg/util"
)

//IUserTag 用户标签
type IUserTag interface {
	// 用户所有标签
	GetTagsByUserID(ctx context.Context, userID uint32) (list []*model.UserTag, err error)
	// 获取对应标签名
	GetTagNamesByIds(ctx context.Context, userID uint32, ids []uint32) (names []string, err error)
	// 批量创建
	TagBatchCreate(ctx context.Context, tags []*model.UserTagModel) (ids []uint32, err error)
}

// GetTagsByUserID 获取用户收藏列表
func (r *Repo) GetTagsByUserID(ctx context.Context, userID uint32) (list []*model.UserTag, err error) {
	if err = r.queryCache(ctx, tagCacheKey(userID), &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.UserTagModel{}).
			Where("user_id = ? ", userID).Order(orm.DefaultOrder).Scan(data).Error; err != nil {
			return errors.Wrap(err, "[repo.tag] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.tag] query cache")
	}
	return
}

// GetTagNamesByIds 标签id获取标签名列表
func (r *Repo) GetTagNamesByIds(ctx context.Context, userID uint32, ids []uint32) (names []string, err error) {
	tags, err := r.GetTagsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.tag] get all err")
	}
	names = make([]string, 0)
	for _, tag := range tags {
		if util.InuInt32Slice(tag.ID, ids) {
			names = append(names, tag.Name)
		}
	}
	return
}

// TagBatchCreate 批量创建标签
func (r *Repo) TagBatchCreate(ctx context.Context, tags []*model.UserTagModel) (ids []uint32, err error) {
	if err = r.db.WithContext(ctx).Create(&tags).Error; err != nil {
		return nil, errors.Wrapf(err, "[repo.tag] batch create err")
	}
	// 删除缓存
	if err = r.cache.Del(ctx, tagCacheKey(tags[0].UserID)); err != nil {
		return nil, errors.Wrap(err, "[repo.tag] delete all cache")
	}
	for _, tag := range tags {
		ids = append(ids, tag.ID)
	}
	return ids, nil
}

func tagCacheKey(uid uint32) string {
	return fmt.Sprintf("tag:all:%d", uid)
}

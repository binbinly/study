package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/prometheus/common/log"

	"chat-micro/app/logic/model"
	"chat-micro/pkg/logger"
)

//IMomentComment 朋友圈评论接口
type IMomentComment interface {
	// 创建
	CommentCreate(ctx context.Context, model *model.MomentCommentModel) (id uint32, err error)
	// 动态评论用户列表
	GetCommentsByMomentID(ctx context.Context, momentID uint32) ([]*model.MomentCommentModel, error)
	// 朋友圈动态下指定动态评论列表
	GetCommentsByMomentIds(ctx context.Context, momentIds []uint32) (map[uint32][]*model.MomentCommentModel, error)
}

// CommentCreate 创建
func (r *Repo) CommentCreate(ctx context.Context, model *model.MomentCommentModel) (id uint32, err error) {
	if err = r.db.WithContext(ctx).Create(model).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.moment_comment] create err")
	}
	// 删除cache
	if err = r.cache.Del(ctx, commentCacheKey(model.MomentID)); err != nil {
		logger.Warnf("[repo.moment_comment] delete cache err:%v", err)
	}
	return model.ID, nil
}

// GetCommentsByMomentID 获取动态下所有评论
func (r *Repo) GetCommentsByMomentID(ctx context.Context, momentID uint32) (list []*model.MomentCommentModel, err error) {
	if err = r.queryCache(ctx, commentCacheKey(momentID), &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Where("moment_id=?", momentID).
			Order("id asc").Find(data).Error; err != nil {
			return errors.Wrap(err, "[repo.moment_comment] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_comment] query cache")
	}
	return
}

// GetCommentsByMomentIds 朋友圈动态下指定动态评论列表
func (r *Repo) GetCommentsByMomentIds(ctx context.Context, ids []uint32) (mComments map[uint32][]*model.MomentCommentModel, err error) {
	mComments = make(map[uint32][]*model.MomentCommentModel, len(ids))

	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, commentCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*[]*model.MomentCommentModel)
	if err = r.cache.MultiGet(ctx, keys, cacheMap, func() interface{} {
		return &[]*model.MomentCommentModel{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_comment] multi get cache data err")
	}

	// 查询未命中
	for _, id := range ids {
		comments, ok := cacheMap[commentCacheKey(id)]
		if !ok {
			cs, err := r.GetCommentsByMomentID(ctx, id)
			if err != nil {
				log.Warnf("[repo.moment_comment] get comment err: %v", err)
				continue
			}
			comments = &cs
		}
		if len(*comments) == 0 {
			continue
		}
		mComments[id] = *comments
	}
	return mComments, nil
}

func commentCacheKey(mid uint32) string {
	return fmt.Sprintf("moment:comment:%d", mid)
}

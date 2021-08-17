package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"chat/app/chat/model"
	"chat/pkg/cache"
	"chat/pkg/log"
	"chat/pkg/redis"
)

//IMomentComment 朋友圈评论接口
type IMomentComment interface {
	// 创建
	CommentCreate(ctx context.Context, model *model.MomentCommentModel) (id uint32, err error)
	// 动态评论用户列表
	GetCommentsByMomentID(ctx context.Context, momentID uint32) (*[]*model.MomentCommentModel, error)
	// 朋友圈动态下指定动态评论列表
	GetCommentsByMomentIds(ctx context.Context, momentIds []uint32) (map[uint32]*[]*model.MomentCommentModel, error)
}

// CommentCreate 创建
func (r *Repo) CommentCreate(ctx context.Context, model *model.MomentCommentModel) (id uint32, err error) {
	err = r.db.WithContext(ctx).Create(model).Error
	if err != nil {
		return 0, errors.Wrapf(err, "[repo.moment_comment] create err")
	}
	// 删除cache
	err = r.commentCache.DelCache(ctx, model.MomentID)
	if err != nil {
		log.Warnf("[repo.moment_comment] delete cache err:%v", err)
	}
	return model.ID, nil
}

// GetCommentsByMomentID 获取动态下所有评论
func (r *Repo) GetCommentsByMomentID(ctx context.Context, momentID uint32) (comments *[]*model.MomentCommentModel, err error) {
	start := time.Now()
	defer func() {
		log.Debugf("[repo.moment_comment] get comments by mid: %d cost: %d μs", momentID, time.Since(start).Microseconds())
	}()
	// 从cache获取
	comments, err = r.commentCache.GetCache(ctx, momentID)
	if err != nil {
		if err == cache.ErrPlaceholder {
			return &[]*model.MomentCommentModel{}, nil
		} else if err != redis.Nil {
			return nil, errors.Wrapf(err, "[repo.moment_comment] get cache comments by mid: %d", momentID)
		}
	}
	if comments != nil {
		log.Debugf("[repo.moment_comment] get comments from cache, mid: %d", momentID)
		return
	}

	getDataFn := func() (interface{}, error) {
		data := make([]*model.MomentCommentModel, 0)
		// 从数据库中获取
		err = r.db.WithContext(ctx).Where("moment_id=?", momentID).Order("id asc").Find(&data).Error
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.moment_comment] query db err")
		}

		// set cache
		err = r.commentCache.SetCache(ctx, momentID, data)
		if err != nil {
			return data, errors.Wrap(err, "[repo.moment_comment] set cache err")
		}
		return &data, nil
	}

	g := singleflight.Group{}
	doKey := fmt.Sprintf("get_moment_comment_%d", momentID)
	val, err, _ := g.Do(doKey, getDataFn)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.moment_comment] get err via single flight do")
	}
	data := val.(*[]*model.MomentCommentModel)

	return data, nil
}

// GetCommentsByMomentIds 朋友圈动态下指定动态评论列表
func (r *Repo) GetCommentsByMomentIds(ctx context.Context, momentIds []uint32) (mComments map[uint32]*[]*model.MomentCommentModel, err error) {
	mComments = make(map[uint32]*[]*model.MomentCommentModel, 0)

	// 从cache批量获取
	commentMap, err := r.commentCache.MultiGetCache(ctx, momentIds)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.moment_comment] multi get comments cache data err")
	}

	// 查询未命中
	for _, momentID := range momentIds {
		idx := r.commentCache.BuildCacheKey([]interface{}{momentID})
		comments, ok := commentMap[idx]
		if !ok {
			comments, err = r.GetCommentsByMomentID(ctx, momentID)
			if err != nil {
				log.Warnf("[repo.moment_comment] get moments err: %v", err)
				continue
			}
		}
		mComments[momentID] = comments
	}
	return mComments, nil
}

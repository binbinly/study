package cache

import (
	"context"

	"chat/app/chat/model"
	"chat/app/constvar"
	"chat/internal"
)

const _momentCommentCacheKey = "moment:comment:%d"

// CommentCache 好友申请
type CommentCache struct {
	*internal.Cache
}

// NewCommentCache new一个朋友圈评论cache
func NewCommentCache() *CommentCache {
	return &CommentCache{internal.NewCache(_momentCommentCacheKey, func() interface{} {
		return &[]*model.MomentCommentModel{}
	})}
}

//SetCache 设置缓存
func (u *CommentCache) SetCache(ctx context.Context, id uint32, comments []*model.MomentCommentModel) error {
	if len(comments) == 0 {
		return u.Redis.SetCacheWithNotFound(ctx, u.BuildCacheKey([]interface{}{id}))
	}
	return u.Redis.Set(ctx, u.BuildCacheKey([]interface{}{id}), comments, constvar.CacheExpireTime)
}

//GetCache 获取缓存
func (u *CommentCache) GetCache(ctx context.Context, id uint32) (comments *[]*model.MomentCommentModel, err error) {
	err = u.Redis.Get(ctx, u.BuildCacheKey([]interface{}{id}), &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// MultiGetCache 批量获取缓存
func (u *CommentCache) MultiGetCache(ctx context.Context, ids []uint32) (map[string]*[]*model.MomentCommentModel, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := u.BuildCacheKey([]interface{}{v})
		keys = append(keys, cacheKey)
	}

	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	commentMap := make(map[string]*[]*model.MomentCommentModel)
	err := u.Redis.MultiGet(ctx, keys, commentMap)
	if err != nil {
		return nil, err
	}
	return commentMap, nil
}

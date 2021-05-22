package cache

import (
	"context"
	"fmt"

	"chat/app/logic/model"
	"chat/pkg/cache"
)

// CommentCache 好友申请
type CommentCache struct {
	cache cache.Driver
	//localCache cache.Driver
}

// NewCommentCache new一个朋友圈评论cache
func NewCommentCache() *CommentCache {
	return &CommentCache{
		cache: newCache(func() interface{} {
			return &[]*model.MomentCommentModel{}
		}),
	}
}

func (u *CommentCache) GetCacheKey(id uint32) string {
	return fmt.Sprintf(MomentCommentCacheKey, id)
}

func (u *CommentCache) SetCache(ctx context.Context, id uint32, comments []*model.MomentCommentModel) error {
	if len(comments) == 0 {
		return u.cache.SetCacheWithNotFound(u.GetCacheKey(id))
	}
	return u.cache.Set(u.GetCacheKey(id), comments, defaultExpireTime)
}

func (u *CommentCache) GetCache(ctx context.Context, id uint32) (comments *[]*model.MomentCommentModel, err error) {
	err = u.cache.Get(u.GetCacheKey(id), &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// MultiGetCache 批量获取缓存
func (u *CommentCache) MultiGetCache(ctx context.Context, ids []uint32) (map[string]*[]*model.MomentCommentModel, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := u.GetCacheKey(v)
		keys = append(keys, cacheKey)
	}

	// 需要在这里make实例化，如果在返回参数里直接定义会报 nil map
	commentMap := make(map[string]*[]*model.MomentCommentModel)
	err := u.cache.MultiGet(keys, commentMap)
	if err != nil {
		return nil, err
	}
	return commentMap, nil
}

// DelCache 删除列表缓存
func (u *CommentCache) DelCache(ctx context.Context, id uint32) error {
	return u.cache.Del(u.GetCacheKey(id))
}

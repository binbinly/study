package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"chat-micro/app/logic/model"
)

const (
	_emoticonCatAllCacheKey = "emoticon:cat:all"
)

//IEmoticon 表情包仓库接口
type IEmoticon interface {
	GetEmoticonCatAll(ctx context.Context) (list []*model.Emoticon, err error)
	GetEmoticonListByCat(ctx context.Context, cat string) (list []*model.Emoticon, err error)
}

//GetEmoticonCatAll 获取表情所有分类
func (r *Repo) GetEmoticonCatAll(ctx context.Context) (list []*model.Emoticon, err error) {
	if err = r.queryCache(ctx, _emoticonCatAllCacheKey, &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.EmoticonModel{}).
			Group("category").Scan(data).Error; err != nil {
			return errors.Wrap(err, "[repo.emoticon] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.emoticon] query cache")
	}
	return
}

//GetEmoticonListByCat 获取分类下所有表情
func (r *Repo) GetEmoticonListByCat(ctx context.Context, cat string) (list []*model.Emoticon, err error) {
	if err = r.queryCache(ctx, emoticonCacheKey(cat), &list, func(data interface{}) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.EmoticonModel{}).
			Where("category=?", cat).Scan(data).Error; err != nil {
			return errors.Wrap(err, "[repo.emoticon] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.emoticon] query cache")
	}
	return
}

func emoticonCacheKey(cat string) string {
	return fmt.Sprintf("emoticon:%s", cat)
}

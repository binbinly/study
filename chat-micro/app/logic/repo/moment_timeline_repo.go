package repo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat-micro/app/logic/model"
)

//IMomentTimeline 朋友圈时间线数据仓库
type IMomentTimeline interface {
	// 批量创建
	TimelineBatchCreate(ctx context.Context, tx *gorm.DB, models []*model.MomentTimelineModel) (ids []uint32, err error)
	// 记录是否存在
	TimelineExist(ctx context.Context, userID, momentID uint32) (bool, error)
}

// TimelineBatchCreate 批量创建
func (r *Repo) TimelineBatchCreate(ctx context.Context, tx *gorm.DB, models []*model.MomentTimelineModel) (ids []uint32, err error) {
	l := len(models)
	var end int
	batchSize := 500 //批处理大小
	// 按大小进行批处理插入，防好友太多，插入数据库失败
	for i := 0; i < l; i += batchSize {
		end = i + batchSize
		if end > l {
			end = l
		}
		sub := models[i:end]
		err = tx.WithContext(ctx).Create(&sub).Error
		if err != nil {
			return nil, errors.Wrapf(err, "[repo.moment_timeline] batch create err")
		}
		for _, m := range sub {
			ids = append(ids, m.ID)
		}
	}
	return ids, nil
}

// TimelineExist 记录是否存在
func (r *Repo) TimelineExist(ctx context.Context, userID, momentID uint32) (is bool, err error) {
	var c int64
	if err = r.db.WithContext(ctx).Model(&model.MomentTimelineModel{}).
		Where("user_id=? and moment_id=?", userID, momentID).Count(&c).Error; err != nil {
		return false, errors.Wrap(err, "[repo.moment_timeline] query db")
	}
	return c > 0, nil
}

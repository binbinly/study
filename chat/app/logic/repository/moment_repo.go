package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat/app/logic/model"
)

type IMoment interface {
	// 创建一条动态
	MomentCreate(ctx context.Context, tx *gorm.DB, message *model.MomentModel) (id uint32, err error)
	// 我的朋友圈列表
	GetMyMoments(ctx context.Context, userId uint32, offset, limit int) ([]*model.MomentModel, error)
	// 指定好友的朋友圈
	GetMomentsByUserId(ctx context.Context, myId, userId uint32, offset, limit int) ([]*model.MomentModel, error)
	// 获取动态信息
	GetMomentById(ctx context.Context, id uint32) (*model.MomentModel, error)
}

// MomentCreate 创建
func (r *Repo) MomentCreate(ctx context.Context, tx *gorm.DB, moment *model.MomentModel) (id uint32, err error) {
	err = tx.Create(&moment).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.moment] create moment err")
	}
	return moment.ID, nil
}

// GetMyMoments 我的朋友圈列表
func (r *Repo) GetMyMoments(ctx context.Context, userId uint32, offset, limit int) (list []*model.MomentModel, err error) {
	err = r.db.Raw("SELECT * from moment where id in (select moment_id from moment_timeline where user_id = ?) order by id desc limit ? offset ?", userId, limit, offset).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, errors.Wrapf(err, "[repo.moment] list err")
	}
	return list, nil
}

// GetMomentsByUserId 指定用户的朋友圈
func (r *Repo) GetMomentsByUserId(ctx context.Context, myId, userId uint32, offset, limit int) (list []*model.MomentModel, err error) {
	if myId == userId { // 查看自己
		err = r.db.Raw("select * from moment where user_id=? order by id desc limit ? offset ?", myId, limit, offset).Find(&list).Error
	} else {
		err = r.db.Raw("select * from moment where (see_type=1 and user_id=?) or (see_type=3 and FIND_IN_SET(?,see)) order by id desc limit ? offset ?", userId, myId, limit, offset).Find(&list).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, errors.Wrapf(err, "[repo.moment] list err")
	}
	return list, nil
}

// GetMomentById 获取动态信息
func (r *Repo) GetMomentById(ctx context.Context, id uint32) (moment *model.MomentModel, err error) {
	moment = new(model.MomentModel)
	err = r.db.First(moment, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(err, "[repo.moment] first err")
	}
	return moment, nil
}
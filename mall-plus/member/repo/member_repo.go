package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"

	"member/model"
)

//IMember 会员接口定义
type IMember interface {
	MemberCreate(ctx context.Context, member *model.MemberModel) error
	MemberUpdate(ctx context.Context, id int64, userMap map[string]interface{}) error
	GetMemberByID(ctx context.Context, id int64) (member *model.MemberModel, err error)
	MemberExist(ctx context.Context, username string, phone int64) (bool, error)
}

// MemberCreate 创建用户
func (r *Repo) MemberCreate(ctx context.Context, member *model.MemberModel) (err error) {
	if err = r.DB.WithContext(ctx).Create(member).Error; err != nil {
		return errors.Wrap(err, "[repo.member] Create err")
	}
	r.delMemberCache(ctx, member.ID)

	return nil
}

// MemberUpdate 更新用户信息
func (r *Repo) MemberUpdate(ctx context.Context, id int64, userMap map[string]interface{}) error {
	if err := r.DB.WithContext(ctx).Model(&model.MemberModel{}).Where("id=?", id).Updates(userMap).Error; err != nil {
		return errors.Wrapf(err, "[repo.member] update")
	}
	r.delMemberCache(ctx, id)

	return nil
}

// GetMemberByID 获取用户
func (r *Repo) GetMemberByID(ctx context.Context, id int64) (member *model.MemberModel, err error) {
	if err = r.QueryCache(ctx, buildMemberCacheKey(id), &member, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).First(data, id).Error; err != nil {
			return errors.Wrapf(err, "[repo.member] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.member] query cache")
	}

	return member, nil
}

// MemberExist 用户是否已存在
func (r *Repo) MemberExist(ctx context.Context, username string, phone int64) (bool, error) {
	var c int64
	if err := r.DB.WithContext(ctx).Model(&model.MemberModel{}).
		Where("phone = ? or username=?", phone, username).Count(&c).Error; err != nil {
		return false, errors.Wrapf(err, "[repo.member] username %v or phone %v does not exist", username, phone)
	}
	return c > 0, nil
}

//delMemberCache 删除会员缓存
func (r *Repo) delMemberCache(ctx context.Context, id int64)  {
	if err := r.Cache.DelCache(ctx, buildMemberCacheKey(id)); err != nil {
		logger.Warnf("[repo.member] del cache key: %v", buildMemberCacheKey(id))
	}
}

//buildMemberCacheKey 构建会员缓存key
func buildMemberCacheKey(id int64) string {
	return fmt.Sprintf("mall_member:%d", id)
}

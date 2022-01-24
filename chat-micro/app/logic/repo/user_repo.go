package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"chat-micro/app/logic/model"
	"chat-micro/pkg/logger"
)

//IUser 会员接口定义
type IUser interface {
	UserCreate(ctx context.Context, user *model.UserModel) (uint32, error)
	UserUpdate(ctx context.Context, id uint32, userMap map[string]interface{}) error
	UserUpdatePwd(ctx context.Context, user *model.UserModel) error
	GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error)
	GetUsersByKeyword(ctx context.Context, keyword string) ([]*model.UserModel, error)
	GetUserByPhone(ctx context.Context, phone int64) (*model.UserModel, error)
	GetUserByID(ctx context.Context, id uint32) (*model.UserModel, error)
	GetUsersByIds(ctx context.Context, ids []uint32) ([]*model.UserModel, error)
	UserExist(ctx context.Context, username string, phone int64) (bool, error)
}

// UserCreate 创建用户
func (r *Repo) UserCreate(ctx context.Context, user *model.UserModel) (id uint32, err error) {
	if err = r.db.WithContext(ctx).Create(user).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.user] Create err")
	}
	r.delUserCache(ctx, user.ID)

	return user.ID, nil
}

// UserUpdate 更新用户信息
func (r *Repo) UserUpdate(ctx context.Context, id uint32, userMap map[string]interface{}) error {
	if err := r.db.WithContext(ctx).Model(&model.UserModel{}).Where("id=?", id).Updates(userMap).Error; err != nil {
		return errors.Wrapf(err, "[repo.user] update")
	}
	r.delUserCache(ctx, id)

	return nil
}

// UserUpdatePwd 修改用户密码
func (r *Repo) UserUpdatePwd(ctx context.Context, user *model.UserModel) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return errors.Wrapf(err, "[repo.user] update pwd")
	}
	return nil
}

// GetUserByID 获取用户
func (r *Repo) GetUserByID(ctx context.Context, id uint32) (user *model.UserModel, err error) {
	if err = r.queryCache(ctx, userCacheKey(id), &user, func(data interface{}) error {
		// 从数据库中获取
		if err := r.db.WithContext(ctx).First(data, id).Error; err != nil {
			return errors.Wrapf(err, "[repo.user] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.user] query cache")
	}

	return
}

func (r *Repo) GetUsersByIds(ctx context.Context, ids []uint32) (users []*model.UserModel, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, userCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.UserModel)
	if err = r.cache.MultiGet(ctx, keys, cacheMap, func() interface{} {
		return &model.UserModel{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.user] multi get cache data err")
	}

	// 查询未命中
	for _, id := range ids {
		user, ok := cacheMap[userCacheKey(id)]
		if !ok {
			user, err = r.GetUserByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.user] get user err: %v", err)
				continue
			}
		}
		if user == nil || user.ID == 0 {
			continue
		}
		users = append(users, user)
	}
	return
}

// GetUserByUsername 根据账号获取用户
func (r *Repo) GetUserByUsername(ctx context.Context, username string) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	if err = r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error;
		err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] get user err by username")
	}
	return user, nil
}

// GetUsersByKeyword 关键字搜索用户
func (r *Repo) GetUsersByKeyword(ctx context.Context, keyword string) (users []*model.UserModel, err error) {
	//最多查询10个用户
	if err = r.db.WithContext(ctx).Where("username like ?", keyword+"%").Limit(10).Find(&users).Error;
		err != nil {
		return nil, errors.Wrap(err, "[repo.user] get user err by username")
	}
	return users, nil
}

// GetUserByPhone 根据手机号获取用户
func (r *Repo) GetUserByPhone(ctx context.Context, phone int64) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	if err = r.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error;
		err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] get user err by phone")
	}
	return user, nil
}

// UserExist 用户是否已存在
func (r *Repo) UserExist(ctx context.Context, username string, phone int64) (bool, error) {
	var c int64
	if err := r.db.WithContext(ctx).Model(&model.UserModel{}).
		Where("phone = ? or username=?", phone, username).Count(&c).Error; err != nil {
		return false, errors.Wrapf(err, "[repo.user] username %v or phone %v does not exist", username, phone)
	}
	return c > 0, nil
}

//delUserCache 删除会员缓存
func (r *Repo) delUserCache(ctx context.Context, id uint32) {
	if err := r.cache.Del(ctx, userCacheKey(id)); err != nil {
		logger.Warnf("[repo.user] del cache key: %v", userCacheKey(id))
	}
}

//userCacheKey 构建会员缓存key
func userCacheKey(id uint32) string {
	return fmt.Sprintf("user:%d", id)
}

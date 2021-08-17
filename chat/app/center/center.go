package center

import (
	"context"
	"errors"

	"chat/app/center/conf"
	"chat/app/center/repository"
	"chat/internal/orm"
	"chat/pkg/queue"
	"chat/pkg/queue/iqueue"
)

var (
	//ErrUserExisted 用户已存在
	ErrUserExisted = errors.New("user:existed")
	//ErrUserNotFound 用户不存在
	ErrUserNotFound = errors.New("user:not found")
	//ErrUserNotMatch 用户名密码不匹配
	ErrUserNotMatch = errors.New("user:not match")
	//ErrUserFrozen 账号已冻结
	ErrUserFrozen = errors.New("user:frozen")
	//ErrUserTokenExpired 用户令牌过期
	ErrUserTokenExpired = errors.New("user: token expired")
	//ErrUserTokenError 用户令牌错误
	ErrUserTokenError = errors.New("user: token error")
)

var _ ICenter = (*Center)(nil)

type ICenter interface {
	IUser
	IOnline
	IPush

	SendSMS(ctx context.Context, phone string) (string, error)
	CheckVCode(ctx context.Context, phone int64, vCode string) error
	Close() error
}

// Service struct
type Center struct {
	c     *conf.Config
	repo  repository.IRepo
	queue iqueue.Producer
}

// New init service
func New(c *conf.Config) ICenter {
	return &Center{
		c:     c,
		repo:  repository.New(orm.GetDB()),
		queue: queue.NewProducer(&c.Queue),
	}
}

// Close service
func (c *Center) Close() error {
	return c.repo.Close()
}

package service

import (
	"context"
	"errors"

	"chat-micro/app/logic/model"
	"chat-micro/app/logic/repo"
	"chat-micro/internal/cache"
	"chat-micro/internal/orm"
)

const (
	msgFriendCreate = "你们已经是好友，可以开始聊天啦"
	msgCloseClient  = "账号已在其他地方登录了!"
)

// 用于触发编译期的接口的合理性检查机制
var _ IService = (*Service)(nil)

var (
	//ErrMomentNotFound 动态不存在
	ErrMomentNotFound = errors.New("moment:not found")
	//ErrReportExisted 举报已存在
	ErrReportExisted = errors.New("report:existed")
	//ErrApplyExisted 申请已存在
	ErrApplyExisted = errors.New("apply:existed")
	//ErrApplyNotFound 申请不存在
	ErrApplyNotFound = errors.New("apply:not found")
	//ErrFriendNotRecord 未找到匹配好友记录
	ErrFriendNotRecord = errors.New("friend:not record")
	//ErrFriendNotFound 好友不存在或已被拉黑
	ErrFriendNotFound = errors.New("chat:friend not found")
	//ErrGroupNotFound 群组不存在
	ErrGroupNotFound = errors.New("group:not found")
	//ErrGroupUserNotJoin 非群组成员
	ErrGroupUserNotJoin = errors.New("group:not join")
	//ErrGroupUserTargetNotJoin 目标用户非群组成员
	ErrGroupUserTargetNotJoin = errors.New("group:target not join")
	//ErrGroupUserExisted 已经是群成员
	ErrGroupUserExisted = errors.New("group:existed")
	//ErrGroupDataUnmodified 数据未修改
	ErrGroupDataUnmodified = errors.New("group:data unmodified")
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
	//ErrVerifyCodeRuleMinute 分钟限制
	ErrVerifyCodeRuleMinute = errors.New("sms:minute limit")
	//ErrVerifyCodeRuleHour 小时限制
	ErrVerifyCodeRuleHour = errors.New("sms:hour limit")
	//ErrVerifyCodeRuleDay 天级限制
	ErrVerifyCodeRuleDay = errors.New("sms:day limit")
	//ErrVerifyCodeNotMatch 验证码不匹配
	ErrVerifyCodeNotMatch = errors.New("code:not match")
)

var Svc IService

// IService 服务接口定义
type IService interface {
	// user
	IUser
	// collect
	ICollect
	// user moment
	IMoment
	// apply
	IApply
	// friend
	IFriend
	// group
	IGroup
	// chat
	IChat

	SendSMS(ctx context.Context, phone string) (string, error)
	CheckVCode(ctx context.Context, phone int64, vCode string) error
	EmoticonCat(ctx context.Context) (list []*model.Emoticon, err error)
	Emoticon(ctx context.Context, cat string) (list []*model.Emoticon, err error)
	ReportCreate(ctx context.Context, UserID, friendID uint32, cType int8, cat, content string) error
	GetUploadSIgnUrl(ctx context.Context, name string) (string, error)
	//Close 关闭服务
	Close() error
}

// Service struct
type Service struct {
	opts options
	repo repo.IRepo
}

// New init service
func New(opts ...Option) (s *Service) {
	s = &Service{
		opts: newOptions(opts...),
		repo: repo.New(orm.GetDB(), cache.NewCache()),
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	return nil
}

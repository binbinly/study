package service

import "context"

// 用于触发编译期的接口的合理性检查机制
var _ IService = (*Service)(nil)

// IService 服务接口定义
type IService interface {
	IOnline
	// user
	IUser
	// collect
	ICollect
	// user moment
	IMoment
	// emoticon
	IEmoticon
	// apply
	IApply
	// friend
	IFriend
	// group
	IGroup
	// chat
	IChat
	// push
	IPush

	//SendSMS 发送短信
	SendSMS(ctx context.Context, phone string) (string, error)
	//CheckVCode 验证验证码是否正确
	CheckVCode(ctx context.Context, phone int64, vCode string) error
	//Close 关闭服务
	Close() error
}

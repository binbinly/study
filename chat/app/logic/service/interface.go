package service

// 用于触发编译期的接口的合理性检查机制
var _ IService = (*Service)(nil)

// service 用户接口定义
type IService interface {
	IOnline
	// user
	IUser
	// collect
	ICollect
	// user moment
	IMoment
	// user report
	IReport
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

	//发送短信
	SendSMS(phoneNumber string) (string, error)
	//验证验证码是否正确
	CheckVCode(phone int64, vCode string) error

	Close() error
}
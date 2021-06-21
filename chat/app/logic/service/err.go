package service

import "errors"

var (
	//ErrUserNotFound 用户不存在
	ErrUserNotFound = errors.New("user:not found")
	//ErrUserLogin 用户登录异常
	ErrUserLogin = errors.New("user:login")
	//ErrMomentNotFound 动态不存在
	ErrMomentNotFound = errors.New("moment:not found")
	//ErrReportExisted 举报已存在
	ErrReportExisted = errors.New("report:existed")
	//ErrUserKeyExisted 用户名或者手机已注册
	ErrUserKeyExisted = errors.New("user:existed")
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
	//ErrVerifyCodeNotMatch 验证码不匹配
	ErrVerifyCodeNotMatch = errors.New("code:empty")
	//ErrVerifyCodeRuleMinute 发送验证码受限
	ErrVerifyCodeRuleMinute = errors.New("sms:minute limit")
	//ErrVerifyCodeRuleHour 小时限制
	ErrVerifyCodeRuleHour = errors.New("sms:hour limit")
	//ErrVerifyCodeRuleDay 天级限制
	ErrVerifyCodeRuleDay = errors.New("sms:day limit")
)
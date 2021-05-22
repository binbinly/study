package service

import "errors"

var (
	// 用户不存在
	ErrUserNotFound = errors.New("user:not found")
	// 用户登录异常
	ErrUserLogin = errors.New("user:login")
	// 动态不存在
	ErrMomentNotFound = errors.New("moment:not found")
	// 举报已存在
	ErrReportExisted = errors.New("report:existed")
	// 用户名或者手机已注册
	ErrUserKeyExisted = errors.New("user:existed")
	// 申请已存在
	ErrApplyExisted = errors.New("apply:existed")
	// 申请不存在
	ErrApplyNotFound = errors.New("apply:not found")
	// 未找到匹配好友记录
	ErrFriendNotRecord = errors.New("friend:not record")
	// 好友不存在或已被拉黑
	ErrFriendNotFound = errors.New("chat:friend not found")
	// 群组不存在
	ErrGroupNotFound = errors.New("group:not found")
	// 非群组成员
	ErrGroupUserNotJoin = errors.New("group:not join")
	// 目标用户非群组成员
	ErrGroupUserTargetNotJoin = errors.New("group:target not join")
	// 已经是群成员
	ErrGroupUserExisted = errors.New("group:existed")
	// 数据未修改
	ErrGroupDataUnmodified = errors.New("group:data unmodified")
	// 验证码不匹配
	ErrVerifyCodeNotMatch = errors.New("code:empty")
	// 发送验证码受限
	ErrVerifyCodeRuleMinute = errors.New("sms:minute limit")
	ErrVerifyCodeRuleHour = errors.New("sms:hour limit")
	ErrVerifyCodeRuleDay = errors.New("sms:day limit")
)
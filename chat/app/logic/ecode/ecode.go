package ecode

import "chat/pkg/errno"

var (
	// user errors

	//ErrUserNotFound err
	ErrUserNotFound          = errno.NewError(20101, "The user was not found.")
	//ErrPasswordIncorrect err
	ErrPasswordIncorrect     = errno.NewError(20102, "账号或密码错误")
	//ErrAreaCodeEmpty err
	ErrAreaCodeEmpty         = errno.NewError(20103, "手机区号不能为空")
	//ErrPhoneEmpty err
	ErrPhoneEmpty            = errno.NewError(20104, "手机号不能为空")
	//ErrGenVCode err
	ErrGenVCode              = errno.NewError(20105, "生成验证码错误")
	//ErrSendSMS err
	ErrSendSMS               = errno.NewError(20106, "发送短信错误")
	//ErrSendSMSMinute err
	ErrSendSMSMinute         = errno.NewError(20107, "一分钟限制一次哦")
	//ErrSendSMSHour err
	ErrSendSMSHour           = errno.NewError(20108, "已超出小时限制，请稍后再试")
	//ErrSendSMSTooMany err
	ErrSendSMSTooMany        = errno.NewError(20109, "已超出当日限制，请明天再试")
	//ErrVerifyCode err
	ErrVerifyCode            = errno.NewError(20110, "验证码错误")
	//ErrEmailOrPassword err
	ErrEmailOrPassword       = errno.NewError(20111, "邮箱或密码错误")
	//ErrTwicePasswordNotMatch err
	ErrTwicePasswordNotMatch = errno.NewError(20112, "两次密码输入不一致")
	//ErrRegisterFailed err
	ErrRegisterFailed        = errno.NewError(20113, "注册失败")
	//ErrUserDisable err
	ErrUserDisable           = errno.NewError(20114, "用户已禁用")
	//ErrUserNoSelf err
	ErrUserNoSelf            = errno.NewError(20115, "不可以操作自己哦")
	//ErrPhoneValid err
	ErrPhoneValid            = errno.NewError(20116, "手机号不合法")
	//ErrUserKeyExisted err
	ErrUserKeyExisted        = errno.NewError(20117, "用户名或者手机号已注册")

	// apply errors

	//ErrApplyFailed err
	ErrApplyFailed         = errno.NewError(20201, "申请失败")
	//ErrApplyRepeatFailed err
	ErrApplyRepeatFailed   = errno.NewError(20202, "已申请过哦")
	//ErrHandleFailed err
	ErrHandleFailed        = errno.NewError(20203, "处理申请失败")
	//ErrApplyNotFoundFailed err
	ErrApplyNotFoundFailed = errno.NewError(20204, "申请未找到哦")

	// chat errors

	//ErrChatNotFound err
	ErrChatNotFound = errno.NewError(20301, "好友不存在或已被拉黑")
	//ErrChatOffline err
	ErrChatOffline  = errno.NewError(20302, "您已离线，请重连")

	//ErrFriendNotFound friend errors
	ErrFriendNotFound = errno.NewError(20401, "好友没有找到哦")

	// group errors

	//ErrGroupCreate err
	ErrGroupCreate         = errno.NewError(20501, "创建群组失败")
	//ErrGroupNotJoin err
	ErrGroupNotJoin        = errno.NewError(20502, "还不是群成员哦")
	//ErrGroupNotFound err
	ErrGroupNotFound       = errno.NewError(20503, "群聊不存在哦")
	//ErrGroupExisted err
	ErrGroupExisted        = errno.NewError(20504, "已经是群成员了哦")
	//ErrGroupSelectNotJoin err
	ErrGroupSelectNotJoin  = errno.NewError(20505, "请选择群成员哦")
	//ErrGroupDataUnmodified err
	ErrGroupDataUnmodified = errno.NewError(20506, "数据未修改哦")

	// collect errors

	//ErrCollectCreate err
	ErrCollectCreate  = errno.NewError(20601, "添加收藏失败")
	//ErrCollectDestroy err
	ErrCollectDestroy = errno.NewError(20602, "删除收藏失败")

	//ErrMomentNotFound err
	ErrMomentNotFound = errno.NewError(20702, "动态不存在哦")

	//ErrReportHanding err
	ErrReportHanding = errno.NewError(20801, "投诉正在处理中哦")
)

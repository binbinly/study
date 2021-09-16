package ecode

import "mall/pkg/errno"

var (
	// user errors
	//ErrUserNotFound err
	ErrUserNotFound = errno.NewError(20101, "用户不存在")
	//ErrPasswordIncorrect err
	ErrPasswordIncorrect = errno.NewError(20102, "账号或密码错误")
	//ErrAreaCodeEmpty err
	ErrAreaCodeEmpty = errno.NewError(20103, "手机区号不能为空")
	//ErrPhoneEmpty err
	ErrPhoneEmpty = errno.NewError(20104, "手机号不能为空")
	//ErrGenVCode err
	ErrGenVCode = errno.NewError(20105, "生成验证码错误")
	//ErrSendSMS err
	ErrSendSMS = errno.NewError(20106, "发送短信错误")
	//ErrSendSMSMinute err
	ErrSendSMSMinute = errno.NewError(20107, "一分钟限制一次哦")
	//ErrSendSMSHour err
	ErrSendSMSHour = errno.NewError(20108, "触发小时级限制")
	//ErrSendSMSTooMany err
	ErrSendSMSTooMany = errno.NewError(20109, "触发天级限制")
	//ErrVerifyCode err
	ErrVerifyCode = errno.NewError(20110, "验证码错误")
	//ErrUsernameOrPassword err
	ErrUsernameOrPassword = errno.NewError(20111, "用户名密码错误")
	//ErrTwicePasswordNotMatch err
	ErrTwicePasswordNotMatch = errno.NewError(20112, "两次密码输入不一致")
	//ErrRegisterFailed err
	ErrRegisterFailed = errno.NewError(20113, "注册失败")
	//ErrUserFrozen err
	ErrUserFrozen = errno.NewError(20114, "账号已被冻结")
	//ErrPhoneValid err
	ErrPhoneValid = errno.NewError(20116, "手机号不合法")
	//ErrUserKeyExisted err
	ErrUserKeyExisted = errno.NewError(20117, "用户名或手机号已存在哦")
	//ErrUserAddressNotFound err
	ErrUserAddressNotFound = errno.NewError(20118, "收货地址不存在")

	//ErrGoodsNotFound err
	ErrGoodsNotFound = errno.NewError(20201, "商品不存在")
	//ErrGoodsSkuNotFound err
	ErrGoodsSkuNotFound = errno.NewError(20202, "商品规格不存在")
	//ErrGoodsSkuNotEdit err
	ErrGoodsSkuNotEdit = errno.NewError(20203, "商品规格未修改")

	//ErrCouponNotFound err
	ErrCouponNotFound = errno.NewError(20301, "优惠券不存在哦")
	//ErrCouponNoNum err
	ErrCouponNoNum = errno.NewError(20302, "优惠券已领完啦")
	//ErrCouponReceived err
	ErrCouponReceived = errno.NewError(20303, "已领取了哦")

	//ErrOrderNotFound err
	ErrOrderNotFound = errno.NewError(20401, "订单不存在")
	//ErrGoodsEmpty err
	ErrGoodsEmpty = errno.NewError(20402, "请选择商品")
	//ErrCouponNotUse err
	ErrCouponNotUse = errno.NewError(20403, "优惠券不可用")
)

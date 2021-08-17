package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat/app/center"
	"chat/app/chat"
	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
	"chat/pkg/utils"
)

// SendCode 获取验证码
// @Summary 根据手机号获取校验码
// @Description 根据手机号获取校验码
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param body body SendCodeParams true "手机号"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /send_code [get]
func SendCode(c *gin.Context) {
	var req SendCodeParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	is := utils.ValidateMobile(req.Phone)
	phone := cast.ToInt64(req.Phone)
	if !is || phone == 0 {
		app.Error(c, ecode.ErrPhoneValid)
		return
	}
	code, err := chat.Svc.SendSMS(c.Request.Context(), phone)
	if errors.Is(err, center.ErrVerifyCodeRuleMinute) {
		app.Error(c, ecode.ErrSendSMSMinute)
		return
	} else if errors.Is(err, center.ErrVerifyCodeRuleHour) {
		app.Error(c, ecode.ErrSendSMSHour)
		return
	} else if errors.Is(err, center.ErrVerifyCodeRuleDay) {
		app.Error(c, ecode.ErrSendSMSTooMany)
		return
	} else if err != nil {
		app.Error(c, ecode.ErrSendSMS)
		log.Warnf("[http.code] send err:%v", err)
		return
	}
	if code != "" { //测试时，直接返回验证码，方便调试
		app.Success(c, code)
		return
	}
	app.Success(c, nil)
}

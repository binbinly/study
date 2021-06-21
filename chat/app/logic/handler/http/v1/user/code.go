package user

import (
	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
	"chat/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
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
	if !is {
		app.Error(c, ecode.ErrPhoneValid)
		return
	}

	code, err := service.Svc.SendSMS(c.Request.Context(), req.Phone)
	if errors.Is(err, service.ErrVerifyCodeRuleMinute) {
		app.Error(c, ecode.ErrSendSMSMinute)
		return
	} else if errors.Is(err, service.ErrVerifyCodeRuleHour) {
		app.Error(c, ecode.ErrSendSMSHour)
		return
	} else if errors.Is(err, service.ErrVerifyCodeRuleDay) {
		app.Error(c, ecode.ErrSendSMSTooMany)
		return
	} else if err != nil {
		app.Error(c, ecode.ErrSendSMS)
		log.Warnf("[http.code] send err:%v", err)
		return
	}
	if code != "" {
		app.Success(c, code)
		return
	}
	app.Success(c, nil)
}

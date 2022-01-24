package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/util"
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
	if v := app.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrBind)
		return
	}
	if is := util.ValidateMobile(req.Phone); !is {
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
		logger.Warnf("[http.code] send err:%v", err)
		return
	}
	if code != "" { //测试时，直接返回验证码，方便调试
		app.Success(c, code)
		return
	}
	app.Success(c, nil)
}

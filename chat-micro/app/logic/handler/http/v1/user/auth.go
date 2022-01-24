package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Login 用户名密码登录
// @Summary 用户登录接口
// @Description 通过用户名密码登录
// @Tags 用户
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @success 0 {object} app.Response{data=model.User} "调用成功结构"
// @Router /login [post]
func Login(c *gin.Context) {
	// Binding the data with the u struct.
	var req LoginParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	user, token, err := service.Svc.UsernameLogin(c.Request.Context(), req.Username, req.Password)
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if errors.Is(err, service.ErrUserFrozen) {
		app.Error(c, ecode.ErrUserFrozen)
		return
	} else if errors.Is(err, service.ErrUserNotMatch) {
		app.Error(c, ecode.ErrUsernameOrPassword)
		return
	} else if err != nil {
		logger.Warnf("[http.user] username login err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

// PhoneLogin 手机登录接口
// @Summary 用户登录接口
// @Description 仅限手机登录
// @Tags 用户
// @Produce  json
// @Param req body PhoneLoginParams true "phone"
// @success 0 {object} app.Response{data=model.User} "调用成功结构"
// @Router /login_phone [post]
func PhoneLogin(c *gin.Context) {
	var req PhoneLoginParams
	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	// 验证验证码
	err := service.Svc.CheckVCode(c.Request.Context(), req.Phone, req.VerifyCode)
	if errors.Is(err, service.ErrVerifyCodeNotMatch) {
		app.Error(c, ecode.ErrVerifyCode)
		return
	} else if err != nil {
		app.Error(c, errno.InternalServerError)
		logger.Warnf("[http.user] check code err:%v", err)
		return
	}

	user, token, err := service.Svc.UserPhoneLogin(c.Request.Context(), req.Phone)
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if errors.Is(err, service.ErrUserFrozen) {
		app.Error(c, ecode.ErrUserFrozen)
		return
	} else if err != nil {
		logger.Warnf("[http.user] phone login err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

// Logout 注销登录
// @Summary 用户注销登录
// @Description 用户注销登录
// @Tags 用户
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /logout [get]
func Logout(c *gin.Context) {
	err := service.Svc.UserLogout(c.Request.Context(), uint32(app.GetUserID(c)))
	if err != nil {
		logger.Warnf("[http.user] logout err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

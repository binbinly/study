package user

import (
	"chat/app/logic/service"
	"errors"
	"github.com/gin-gonic/gin"

	"chat/app/logic/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Login 用户名密码登录
// @Summary 用户登录接口
// @Description 通过用户名密码登录
// @Tags 用户
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik"}}"
// @Router /login [post]
func Login(c *gin.Context) {
	// Binding the data with the u struct.
	var req LoginParams

	valid := app.BindJson(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	user, err := service.Svc.UsernameLogin(c, req.Username, req.Password)
	if err != nil {
		log.Warnf("[http.user] username login err: %v", err)
		app.Error(c, ecode.ErrEmailOrPassword)
		return
	}
	app.Success(c, user)
}

// PhoneLogin 手机登录接口
// @Summary 用户登录接口
// @Description 仅限手机登录
// @Tags 用户
// @Produce  json
// @Param req body PhoneLoginParams true "phone"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik"}}"
// @Router /login_phone [post]
func PhoneLogin(c *gin.Context) {
	var req PhoneLoginParams
	valid := app.BindJson(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	// 验证验证码
	err := service.Svc.CheckVCode(req.Phone, req.VerifyCode)
	if errors.Is(err, service.ErrVerifyCodeNotMatch) {
		app.Error(c, ecode.ErrVerifyCode)
		return
	} else if err != nil {
		app.Error(c, errno.InternalServerError)
		log.Warnf("[http.user] check code err:%v", err)
		return
	}

	user, err := service.Svc.UserPhoneLogin(c, req.Phone)
	if err != nil {
		log.Warnf("[http.user] phone login err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, user)
}

// PhoneLogin 注销登录
// @Summary 用户注销登录
// @Description 用户注销登录
// @Tags 用户
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik"}}"
// @Router /logout [get]
func Logout(c *gin.Context) {
	err := service.Svc.UserLogout(c, uint32(app.GetUserId(c)))
	if err != nil {
		log.Warnf("[http.user] logout err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

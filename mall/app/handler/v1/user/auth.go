package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"mall/app/ecode"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// Login 用户名密码登录
// @Summary 用户登录接口
// @Description 通过用户名密码登录
// @Tags 用户
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @success 0 {object} app.Response{data=model.UserToken} "调用成功结构"
// @Router /login [post]
func Login(c *gin.Context) {
	// Binding the data with the u struct.
	var req LoginParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	user, err := service.Svc.UsernameLogin(c.Request.Context(), req.Username, req.Password)
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
		log.Warnf("[v1.user] username login err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, user)
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
	err := service.Svc.UserLogout(c.Request.Context(), app.GetUserID(c))
	if err != nil {
		log.Warnf("[v1.user] logout err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

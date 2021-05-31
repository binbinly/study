package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"chat/app/logic/ecode"
	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
	"chat/pkg/valid"
)

// Register 注册
// @Summary 注册
// @Description 用户注册
// @Tags 用户
// @Produce  json
// @Param phone body string true "手机号"
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param confirm_password body string true "确认密码"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6Ik"}}"
// @Router /register [post]
func Register(c *gin.Context) {
	var req RegisterParams
	v := app.BindJson(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}
	is := valid.ValidateMobile(req.Phone)
	phone := cast.ToInt64(req.Phone)
	if !is || phone == 0 {
		app.Error(c, ecode.ErrPhoneValid)
		return
	}
	err := service.Svc.UserRegister(c, req.Username, req.Password, phone)
	if errors.Is(err, service.ErrUserKeyExisted) {
		app.Error(c, ecode.ErrUserKeyExisted)
		return
	} else if err != nil {
		log.Warnf("[http.user] register err: %v", err)
		app.Error(c, ecode.ErrRegisterFailed)
		return
	}
	app.SuccessNil(c)
}
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

// Update 更新用户信息
// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user body UpdateParams true "The user info"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /user/edit [post]
func Update(c *gin.Context) {
	var req UpdateParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	userMap := make(map[string]interface{})
	if req.Avatar != "" {
		userMap["avatar"] = req.Avatar
	}
	if req.Nickname != "" {
		userMap["nickname"] = req.Nickname
	}
	if req.Sign != "" {
		userMap["sign"] = req.Sign
	}
	if len(userMap) == 0 {
		app.Error(c, errno.ErrParamsEmpty)
		return
	}
	err := service.Svc.UserEdit(c.Request.Context(), app.GetUInt32UserID(c), userMap)
	if errors.Is(err, service.ErrUserNotFound) {
		app.Error(c, ecode.ErrUserNotFound)
		return
	} else if errors.Is(err, service.ErrUserFrozen) {
		app.Error(c, ecode.ErrUserFrozen)
		return
	} else if err != nil {
		logger.Warnf("[http.user] update err, %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

package group

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Create 创建
// @Summary 创建群组
// @Description 创建群组
// @Tags 群组
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param ids body string true "用户id列表"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /group/create [post]
func Create(c *gin.Context) {
	var req IdsParams
	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.GroupCreate(c.Request.Context(), app.GetUInt32UserID(c), req.Ids)
	if errors.Is(err, service.ErrFriendNotRecord) {
		app.Error(c, ecode.ErrFriendNotFound)
		return
	} else if err != nil {
		logger.Warnf("[http.group] create err: %v", err)
		app.Error(c, ecode.ErrGroupCreate)
		return
	}
	app.SuccessNil(c)
}

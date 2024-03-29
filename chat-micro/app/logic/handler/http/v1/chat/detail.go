package chat

import (
	"errors"

	"github.com/gin-gonic/gin"

	"chat-micro/app/logic/ecode"
	"chat-micro/app/logic/service"
	"chat-micro/pkg/app"
	"chat-micro/pkg/errno"
	"chat-micro/pkg/logger"
)

// Detail 获取聊天信息
// @Summary 获取聊天信息
// @Description 获取聊天信息
// @Tags 聊天
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id body int true "用户/群组id"
// @Param type body int true "类型，1=用户，2=群组"
// @success 0 {object} app.Response{data=message.Sender} "调用成功结构"
// @Router /chat/detail [post]
func Detail(c *gin.Context) {
	var req DetailParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	info, err := service.Svc.ChatDetail(c.Request.Context(), app.GetUInt32UserID(c), req.ID, req.Type)
	if errors.Is(err, service.ErrFriendNotFound) {
		app.Error(c, ecode.ErrChatNotFound)
		return
	} else if errors.Is(err, service.ErrGroupNotFound) {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if err != nil {
		logger.Warnf("[http.chat] detail err:%v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, info)
}

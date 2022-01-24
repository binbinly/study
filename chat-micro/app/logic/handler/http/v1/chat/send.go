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

// Send 发送消息
// @Summary 发送消息
// @Description 发送消息
// @Tags 聊天
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body SendParams true "send"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /chat/send [post]
func Send(c *gin.Context) {
	var req SendParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	msg, err := service.Svc.ChatSend(c.Request.Context(), app.GetUInt32UserID(c), req.ToID, req.Type, req.ChatType, req.Content, req.Options)
	if errors.Is(err, service.ErrFriendNotFound) {
		app.Error(c, ecode.ErrChatNotFound)
		return
	} else if errors.Is(err, service.ErrGroupNotFound) {
		app.Error(c, ecode.ErrChatNotFound)
		return
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	} else if err != nil {
		logger.Warnf("[http.chat] send err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, msg)
}

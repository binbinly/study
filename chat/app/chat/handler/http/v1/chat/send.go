package chat

import (
	"chat/app/chat"
	"errors"

	"github.com/gin-gonic/gin"

	"chat/app/chat/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
)

// Send 发送消息
// @Summary 发送消息
// @Description 发送消息
// @Tags 聊天
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param type body int true "聊天信息类型"
// @Param chat_type body int true "聊天类型，1=用户，2=群组"
// @Param to_id body int true "用户/群组ID"
// @Param content body string true "内容"
// @Param options body []byte false "额外选项"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /chat/send [post]
func Send(c *gin.Context) {
	var req SendParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	msg, err := chat.Svc.ChatSend(c.Request.Context(), app.GetUInt32UserID(c), req.ToID, req.Type, req.ChatType, req.Content, req.Options)
	if errors.Is(err, chat.ErrFriendNotFound) {
		app.Error(c, ecode.ErrChatNotFound)
		return
	} else if errors.Is(err, chat.ErrGroupNotFound) {
		app.Error(c, ecode.ErrChatNotFound)
		return
	} else if err != nil {
		log.Warnf("[http.chat] send err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, msg)
}

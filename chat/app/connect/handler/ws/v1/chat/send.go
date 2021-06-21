package chat

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"chat/app/connect"
	"chat/app/logic/ecode"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
	"chat/proto"
	"chat/proto/logic"
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
// @Success 200 {string} json "{}"
// @Router /chat/send [post]
func Send(c *gin.Context) {
	var req SendParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	input := &TransferChatInput{
		Event:  proto.EventChatSend,
		UserID: uint32(app.GetUserId(c)),
		Send:   &req,
	}
	res, err := connect.Svc.Receive(c, TransChatReq(input))
	if err != nil {
		log.Warnf("[http.chat] send err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	if res.Code == logic.ReceiveReply_ErrFriendNotFound {
		app.Error(c, ecode.ErrChatNotFound)
		return
	} else if res.Code == logic.ReceiveReply_ErrGroupNotFound {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if res.Code == logic.ReceiveReply_ErrUserOffline {
		app.Error(c, ecode.ErrChatOffline)
		return
	}
	if res.Data == nil {
		app.Error(c, ecode.ErrChatNotFound)
		return
	}
	app.Success(c, json.RawMessage(res.Data))
}

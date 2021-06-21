package chat

import (
	"github.com/gin-gonic/gin"

	"chat/app/connect"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
	"chat/proto"
)

// Recall 消息撤回
// @Summary 消息撤回
// @Description 消息撤回
// @Tags 聊天
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body RecallParams true "recall"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{}"
// @Router /chat/recall [post]
func Recall(c *gin.Context) {
	var req RecallParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	input := &TransferChatInput{
		Event:  proto.EventChatRecall,
		UserID: uint32(app.GetUserId(c)),
		Recall: &req,
	}
	_, err := connect.Svc.Receive(c, TransChatReq(input))
	if err != nil {
		log.Warnf("[http.chat] recall err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

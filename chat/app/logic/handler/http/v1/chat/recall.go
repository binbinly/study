package chat

import (
	"github.com/gin-gonic/gin"

	"chat/app/logic/service"
	"chat/pkg/app"
	"chat/pkg/errno"
	"chat/pkg/log"
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

	valid := app.BindJson(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	err := service.Svc.ChatRecall(c, app.GetUInt32UserId(c), req.ToId, req.ChatType, req.Id)
	if err != nil {
		log.Warnf("[http.chat] recall err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

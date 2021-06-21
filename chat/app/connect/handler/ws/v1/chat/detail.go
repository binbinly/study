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

// Detail 获取聊天信息
// @Summary 获取聊天信息
// @Description 获取聊天信息
// @Tags 聊天
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id body int true "用户/群组id"
// @Param type body int true "类型，1=用户，2=群组"
// @Success 200 {string} json "{"code":0,"message":"OK","data":null}"
// @Router /chat/detail [post]
func Detail(c *gin.Context) {
	var req DetailParams

	valid := app.BindJSON(c, &req)
	if !valid {
		app.Error(c, errno.ErrBind)
		return
	}
	input := &TransferChatInput{
		Event:  proto.EventChatDetail,
		UserID: uint32(app.GetUserId(c)),
		Detail: &req,
	}
	res, err := connect.Svc.Receive(c, TransChatReq(input))
	if err != nil {
		log.Warnf("[http.chat] detail err:%v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	if res.Code == logic.ReceiveReply_ErrFriendNotFound {
		app.Error(c, ecode.ErrChatNotFound)
		return
	} else if res.Code == logic.ReceiveReply_ErrGroupNotFound {
		app.Error(c, ecode.ErrGroupNotFound)
		return
	} else if res.Code == logic.ReceiveReply_ErrGroupUserNotJoin {
		app.Error(c, ecode.ErrGroupNotJoin)
		return
	}
	app.Success(c, json.RawMessage(res.Data))
}

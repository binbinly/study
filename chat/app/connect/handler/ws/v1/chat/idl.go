package chat

import (
	"encoding/json"

	"chat/proto/logic"
)

//TransferChatInput 会话对外转化结构
type TransferChatInput struct {
	UserID uint32
	Event  string
	Detail *DetailParams
	Send   *SendParams
	Recall *RecallParams
}

//TransChatReq 转化会话结构
func TransChatReq(input *TransferChatInput) *logic.ReceiveReq {
	req := &logic.ReceiveReq{Uid: input.UserID, Event: input.Event}
	if input.Detail != nil {
		req.Body, _ = json.Marshal(input.Detail)
	} else if input.Send != nil {
		req.Body, _ = json.Marshal(input.Send)
	} else if input.Recall != nil {
		req.Body, _ = json.Marshal(input.Recall)
	}
	return req
}

package chat

import (
	"encoding/json"

	"chat/proto/logic"
)

type TransferChatInput struct {
	UserId uint32
	Event  string
	Detail *DetailParams
	Send   *SendParams
	Recall *RecallParams
}

func TransChatReq(input *TransferChatInput) *logic.ReceiveReq {
	req := &logic.ReceiveReq{Uid: input.UserId, Event: input.Event}
	if input.Detail != nil {
		req.Body, _ = json.Marshal(input.Detail)
	} else if input.Send != nil {
		req.Body, _ = json.Marshal(input.Send)
	} else if input.Recall != nil {
		req.Body, _ = json.Marshal(input.Recall)
	}
	return req
}

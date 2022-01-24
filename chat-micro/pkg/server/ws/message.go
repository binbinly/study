package ws

import "encoding/json"

type wsMessage struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

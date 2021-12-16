// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/katyusha.

package tracing

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

var (
	jsonPbMarshaller = &jsonpb.Marshaler{
		EmitDefaults: true,
	}
)

// MarshalPbMessageToJSONString marshals protobuf message to json string.
func MarshalPbMessageToJSONString(msg proto.Message) string {
	msgJSONStr, _ := jsonPbMarshaller.MarshalToString(msg)
	return msgJSONStr
}

//MarshalMessageToJSONStringForTracing marshal
func MarshalMessageToJSONStringForTracing(value interface{}, msgType string, maxBytes int) string {
	var messageContent string
	if msg, ok := value.(proto.Message); ok {
		if proto.Size(msg) <= maxBytes {
			messageContent = MarshalPbMessageToJSONString(msg)
		} else {
			messageContent = fmt.Sprintf(
				"[%s Message Too Large For Tracing, Max: %d bytes]",
				msgType,
				maxBytes,
			)
		}
	} else {
		messageContent = fmt.Sprintf("%v", value)
	}
	return messageContent
}

// generateID 生成随机字符串，eg: 76d27e8c-a80e-48c8-ad20-e5562e0f67e4
func generateID() string {
	reqID, _ := uuid.NewRandom()
	return reqID.String()
}

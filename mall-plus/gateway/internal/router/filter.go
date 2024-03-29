package router

import (
	"context"
	"net/http"
	"reflect"

	"google.golang.org/protobuf/proto"

	pb "gateway/proto/common"
)

//filter 过滤器，改变响应消息或设置响应头
func filter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	corsHeader(w)
	//fmt.Printf("req:%v\n", resp)
	// 填充默认值
	if _, ok := resp.(*pb.SuccessEmptyReply); ok {
		proto.Merge(resp, &pb.SuccessEmptyReply{
			Code:    0,
			Message: "ok",
		})
	} else {
		m := reflect.ValueOf(resp).Elem().FieldByName("Message")
		if m.IsValid() {
			m.SetString("ok")
		}
	}
	return nil
}

//corsHeader 跨域响应头设置
func corsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

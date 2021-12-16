package util

import (
	"context"
	"github.com/spf13/cast"
	"go-micro.dev/v4/metadata"
)

//GetUserID 获取用户id
func GetUserID(ctx context.Context) int64 {
	md, _ := metadata.FromContext(ctx)
	userID := md["User-Id"]
	return cast.ToInt64(userID)
}

//CHeckTokenEmpty 验证是否携带令牌
func CHeckTokenEmpty(ctx context.Context) bool {
	md, _ := metadata.FromContext(ctx)
	if md["Token"] == "" {
		return false
	}
	return true
}
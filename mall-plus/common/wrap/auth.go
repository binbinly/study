package wrap

import (
	"common/errno"
	"common/util"
	"context"
	"go-micro.dev/v4/server"
)

// AuthFunc 执行身份验证的可插入函数。
type AuthFunc func(method string) bool

//Auth 身份验证包装器
func Auth(authFunc AuthFunc) server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			if authFunc == nil || !authFunc(req.Method()) { // 不需要验证
				return fn(ctx, req, rsp)
			}
			if userID := util.GetUserID(ctx); userID == 0 {
				// 判断是否携带了令牌，以返回不同的错误信息
				if util.CHeckTokenEmpty(ctx) { // 令牌已过期或者不合法
					return errno.CenterReplyErr(errno.ErrUserTokenExpired)
				}
				return errno.CenterReplyErr(errno.ErrUserTokenEmpty)
			}
			return fn(ctx, req, rsp)
		}
	}
}

package router

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-micro.dev/v4/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pkg/logger"
)

//handleRoutingError 自定义路由错误
func handleRoutingError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	// 放行所有OPTIONS方法，因为有的模板是要请求两次的
	if r.Method == "OPTIONS" {
		corsHeader(w)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if httpStatus != http.StatusMethodNotAllowed {
		runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, httpStatus)
		return
	}

	// Use HTTPStatusError to customize the DefaultHTTPErrorHandler status code
	err := &runtime.HTTPStatusError{
		HTTPStatus: httpStatus,
		Err:        status.Error(codes.Unimplemented, http.StatusText(httpStatus)),
	}

	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}

//handleError 错误处理
func handleError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	code := codes.Code(500)
	msg := "Internal Server Error"
	if s, ok := status.FromError(err); ok {
		if s.Code() == codes.InvalidArgument {
			code = codes.Code(400)
			msg = "请求参数非法"
			logger.Infof("method: %v, code: %v, msg: %v", r.Method, s.Code(), s.Message())
		} else {
			e := errors.Parse(s.Message())
			if e.Code != 0 {
				code = codes.Code(e.Code)
				msg = e.Detail
			} else {
				if s.Code() == codes.InvalidArgument {
					// 此处不返回真实错误
					logger.Warnf("[gateway] code: %v，msg: %v", s.Code(), s.Message())
				} else {
					code = s.Code()
					msg = s.Message()
				}
			}
		}
	} else {
		logger.Warnf("[gateway]err: %v", err)
	}
	logger.Infof("method: %v, code: %v, msg: %v", r.Method, code, msg)

	err = &runtime.HTTPStatusError{
		HTTPStatus: http.StatusOK,
		Err:        status.Error(code, msg),
	}
	corsHeader(w)
	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}

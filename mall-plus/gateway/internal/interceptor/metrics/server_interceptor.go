// Copyright (c) The go-grpc-middleware Authors.
// Licensed under the Apache License 2.0.

package metrics

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// UnaryServerInterceptor is a gRPC server-side interceptor that provides Prometheus monitoring for Unary RPCs.
func UnaryServerInterceptor(serverMetrics *ServerMetrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reportable := &reportable{serverMetrics: serverMetrics}
		r := newReport(Unary, info.FullMethod)
		reporter, newCtx := reportable.ServerReporter(ctx, CallMeta{ReqProtoOrNil: req, Typ: r.rpcType, Service: r.service, Method: r.method})

		reporter.PostMsgReceive(req, nil, time.Since(r.startTime))
		resp, err := handler(newCtx, req)
		reporter.PostMsgSend(resp, err, time.Since(r.startTime))

		reporter.PostCall(err, time.Since(r.startTime))
		return resp, err
	}
}

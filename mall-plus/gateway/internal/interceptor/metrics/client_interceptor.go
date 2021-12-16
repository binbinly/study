// Copyright (c) The go-grpc-middleware Authors.
// Licensed under the Apache License 2.0.

package metrics

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"time"
)

// UnaryClientInterceptor is a gRPC client-side interceptor that provides Prometheus monitoring for Unary RPCs.
func UnaryClientInterceptor(clientMetrics *ClientMetrics) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		reportable := &reportable{clientMetrics: clientMetrics}
		r := newReport(Unary, method)
		reporter, newCtx := reportable.ClientReporter(ctx, CallMeta{ReqProtoOrNil: req, Typ: r.rpcType, Service: r.service, Method: r.method})

		reporter.PostMsgSend(req, nil, time.Since(r.startTime))
		err := invoker(newCtx, method, req, reply, cc, opts...)
		reporter.PostMsgReceive(reply, err, time.Since(r.startTime))

		reporter.PostCall(err, time.Since(r.startTime))
		return err
	}
}

type report struct {
	rpcType   grpcType
	service   string
	method    string
	startTime time.Time
}

func newReport(typ grpcType, fullMethod string) report {
	r := report{
		startTime: time.Now(),
		rpcType:   typ,
	}
	r.service, r.method = splitMethodName(fullMethod)
	return r
}

func splitMethodName(fullMethod string) (string, string) {
	fullMethod = strings.TrimPrefix(fullMethod, "/") // remove leading slash
	if i := strings.Index(fullMethod, "/"); i >= 0 {
		return fullMethod[:i], fullMethod[i+1:]
	}
	return "unknown", "unknown"
}
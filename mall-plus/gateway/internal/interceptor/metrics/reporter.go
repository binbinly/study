// Copyright (c) The go-grpc-middleware Authors.
// Licensed under the Apache License 2.0.

package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
)

var (
	//AllCodes codes all
	AllCodes = []codes.Code{
		codes.OK, codes.Canceled, codes.Unknown, codes.InvalidArgument, codes.DeadlineExceeded, codes.NotFound,
		codes.AlreadyExists, codes.PermissionDenied, codes.Unauthenticated, codes.ResourceExhausted,
		codes.FailedPrecondition, codes.Aborted, codes.OutOfRange, codes.Unimplemented, codes.Internal,
		codes.Unavailable, codes.DataLoss,
	}
)

type reporter struct {
	clientMetrics           *ClientMetrics
	serverMetrics           *ServerMetrics
	typ                     grpcType
	service, method         string
	startTime               time.Time
	kind                    Kind
	sendTimer, receiveTimer Timer
}

func (r *reporter) PostCall(err error, duration time.Duration) {
	// get status code from error
	status, _ := FromError(err)
	code := status.Code()

	// perform handling of metrics from code
	switch r.kind {
	case KindServer:
		r.serverMetrics.serverHandledCounter.WithLabelValues(string(r.typ), r.service, r.method, code.String()).Inc()
		if r.serverMetrics.serverHandledHistogram != nil {
			r.serverMetrics.serverHandledHistogram.WithLabelValues(string(r.typ), r.service, r.method).Observe(time.Since(r.startTime).Seconds())
		}

	case KindClient:
		r.clientMetrics.clientHandledCounter.WithLabelValues(string(r.typ), r.service, r.method, code.String()).Inc()
		if r.clientMetrics.clientHandledHistogram != nil {
			r.clientMetrics.clientHandledHistogram.WithLabelValues(string(r.typ), r.service, r.method).Observe(time.Since(r.startTime).Seconds())
		}
	}
}

func (r *reporter) PostMsgSend(_ interface{}, _ error, _ time.Duration) {
	switch r.kind {
	case KindServer:
		r.serverMetrics.serverStreamMsgSent.WithLabelValues(string(r.typ), r.service, r.method).Inc()
	case KindClient:
		r.clientMetrics.clientStreamMsgSent.WithLabelValues(string(r.typ), r.service, r.method).Inc()
		r.sendTimer.ObserveDuration()
	}
}

func (r *reporter) PostMsgReceive(_ interface{}, _ error, _ time.Duration) {
	switch r.kind {
	case KindServer:
		r.serverMetrics.serverStreamMsgReceived.WithLabelValues(string(r.typ), r.service, r.method).Inc()
	case KindClient:
		r.clientMetrics.clientStreamMsgReceived.WithLabelValues(string(r.typ), r.service, r.method).Inc()
		r.receiveTimer.ObserveDuration()
	}
}

type reportable struct {
	clientMetrics *ClientMetrics
	serverMetrics *ServerMetrics
}

func (rep *reportable) ServerReporter(ctx context.Context, meta CallMeta) (Reporter, context.Context) {
	return rep.reporter(ctx, rep.serverMetrics, nil, meta.Typ, meta.Service, meta.Method, KindServer)
}

func (rep *reportable) ClientReporter(ctx context.Context, meta CallMeta) (Reporter, context.Context) {
	return rep.reporter(ctx, nil, rep.clientMetrics, meta.Typ, meta.Service, meta.Method, KindClient)
}

func (rep *reportable) reporter(ctx context.Context, sm *ServerMetrics, cm *ClientMetrics, rpcType grpcType, service, method string, kind Kind) (Reporter, context.Context) {
	r := &reporter{
		clientMetrics: cm,
		serverMetrics: sm,
		typ:           rpcType,
		service:       service,
		method:        method,
		kind:          kind,
		sendTimer:     EmptyTimer,
		receiveTimer:  EmptyTimer,
	}

	switch kind {
	case KindClient:
		if r.clientMetrics.clientHandledHistogram != nil {
			r.startTime = time.Now()
		}
		r.clientMetrics.clientStartedCounter.WithLabelValues(string(r.typ), r.service, r.method).Inc()

		if r.clientMetrics.clientStreamSendHistogram != nil {
			hist := r.clientMetrics.clientStreamSendHistogram.WithLabelValues(string(r.typ), r.service, r.method)
			r.sendTimer = prometheus.NewTimer(hist)
		}

		if r.clientMetrics.clientStreamRecvHistogram != nil {
			hist := r.clientMetrics.clientStreamRecvHistogram.WithLabelValues(string(r.typ), r.service, r.method)
			r.receiveTimer = prometheus.NewTimer(hist)
		}
	case KindServer:
		if r.serverMetrics.serverHandledHistogram != nil {
			r.startTime = time.Now()
		}
		r.serverMetrics.serverStartedCounter.WithLabelValues(string(r.typ), r.service, r.method).Inc()
	}
	return r, ctx
}

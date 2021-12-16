// Copyright (c) The go-grpc-middleware Authors.
// Licensed under the Apache License 2.0.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type serverMetricsConfig struct {
	counterOpts counterOptions
	// serverHandledHistogram can be nil
	serverHandledHistogram *prometheus.HistogramVec
}

//ServerMetricsOption metrics option
type ServerMetricsOption func(*serverMetricsConfig)

func (c *serverMetricsConfig) apply(opts []ServerMetricsOption) {
	for _, o := range opts {
		o(c)
	}
}

//WithServerCounterOptions with counter
func WithServerCounterOptions(opts ...CounterOption) ServerMetricsOption {
	return func(o *serverMetricsConfig) {
		o.counterOpts = opts
	}
}

// WithServerHandlingTimeHistogram turns on recording of handling time of RPCs.
// Histogram metrics can be very expensive for Prometheus to retain and query.
func WithServerHandlingTimeHistogram(opts ...HistogramOption) ServerMetricsOption {
	return func(o *serverMetricsConfig) {
		o.serverHandledHistogram = prometheus.NewHistogramVec(
			histogramOptions(opts).apply(prometheus.HistogramOpts{
				Name:    "grpc_server_handling_seconds",
				Help:    "Histogram of response latency (seconds) of gRPC that had been application-level handled by the server.",
				Buckets: prometheus.DefBuckets,
			}),
			[]string{"grpc_type", "grpc_service", "grpc_method"},
		)
	}
}

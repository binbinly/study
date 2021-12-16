package tracing

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	"pkg/utils"
)

var commonAttrs = []attribute.KeyValue{
	attribute.String("hostname", utils.Hostname),
	attribute.String("local-ip", localIP),
}

func setClientSpan(ctx context.Context, span trace.Span, m interface{}) {
	attrs := commonAttrs
	if p, ok := m.(proto.Message); ok {
		attrs = append(attrs, attribute.Key("send_msg.size").Int(proto.Size(p)))
	}

	span.SetAttributes(attrs...)
}

func setServerSpan(ctx context.Context, span trace.Span, m interface{}) {
	attrs := commonAttrs
	if p, ok := m.(proto.Message); ok {
		attrs = append(attrs, attribute.Key("recv_msg.size").Int(proto.Size(p)))
	}

	span.SetAttributes(attrs...)
}

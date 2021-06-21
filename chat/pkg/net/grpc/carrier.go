package grpc

import "google.golang.org/grpc/metadata"

// MDCarrier custom carrier
type MDCarrier struct {
	MD metadata.MD
}

// ForeachKey conforms to the TextMapReader interface.
// 这里必须要实现这个 TextMapReader 这个接口
// TextMapReader is the Extract() carrier for the TextMap builtin format. With it,
// the caller can decode a propagated SpanContext as entries in a map of
// unicode strings.
//type TextMapReader interface {
//    // ForeachKey returns TextMap contents via repeated calls to the `handler`
//    // function. If any call to `handler` returns a non-nil error, ForeachKey
//    // terminates and returns that error.
//    //
//    // NOTE: The backing store for the TextMapReader may contain data unrelated
//    // to SpanContext. As such, Inject() and Extract() implementations that
//    // call the TextMapWriter and TextMapReader interfaces must agree on a
//    // prefix or other convention to distinguish their own key:value pairs.
//    //
//    // The "foreach" callback pattern reduces unnecessary copying in some cases
//    // and also allows implementations to hold locks while the map is read.
//    ForeachKey(handler func(key, val string) error) error
//}
func (m MDCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, strs := range m.MD {
		for _, v := range strs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set implements Set() of opentracing.TextMapWriter
// 这里也必须要实现
// TextMapWriter is the Inject() carrier for the TextMap builtin format. With
// it, the caller can encode a SpanContext for propagation as entries in a map
// of unicode strings.
//type TextMapWriter interface {
//    // Set a key:value pair to the carrier. Multiple calls to Set() for the
//    // same key leads to undefined behavior.
//    //
//    // NOTE: The backing store for the TextMapWriter may contain data unrelated
//    // to SpanContext. As such, Inject() and Extract() implementations that
//    // call the TextMapWriter and TextMapReader interfaces must agree on a
//    // prefix or other convention to distinguish their own key:value pairs.
//    Set(key, val string)
//}
func (m MDCarrier) Set(key, val string) {
	m.MD[key] = append(m.MD[key], val)
}

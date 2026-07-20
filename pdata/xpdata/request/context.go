package request

import (
	"context"
	"net"

	"go.opentelemetry.io/collector/pdata/internal"
)

func encodeContext(ctx context.Context) *internal.RequestContext {
	_ = "STUB: not implemented"
	return nil
}

func encodeSpanContext(ctx context.Context, rc *internal.RequestContext) {
	_ = "STUB: not implemented"
	return
}

func encodeClientMetadata(ctx context.Context, rc *internal.RequestContext) {
	_ = "STUB: not implemented"
	return
}

func encodeClientAddress(ctx context.Context, rc *internal.RequestContext) {
	_ = "STUB: not implemented"
	return
}

func decodeContext(ctx context.Context, rc *internal.RequestContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func decodeSpanContext(ctx context.Context, sc *internal.SpanContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func decodeClientMetadata(clientMetadata []internal.KeyValue) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func decodeClientAddress(rc *internal.RequestContext) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

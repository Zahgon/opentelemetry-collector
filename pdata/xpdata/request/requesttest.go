package request

import (
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func fakeSpanContext(tb testing.TB) trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type SpanKind int32

const (
	SpanKindUnspecified = SpanKind(internal.SpanKind_SPAN_KIND_UNSPECIFIED)

	SpanKindInternal = SpanKind(internal.SpanKind_SPAN_KIND_INTERNAL)

	SpanKindServer = SpanKind(internal.SpanKind_SPAN_KIND_SERVER)

	SpanKindClient = SpanKind(internal.SpanKind_SPAN_KIND_CLIENT)

	SpanKindProducer = SpanKind(internal.SpanKind_SPAN_KIND_PRODUCER)

	SpanKindConsumer = SpanKind(internal.SpanKind_SPAN_KIND_CONSUMER)
)

func (sk SpanKind) String() string { _ = "STUB: not implemented"; return "" }

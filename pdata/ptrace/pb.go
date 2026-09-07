package ptrace

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalTraces(td Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ProtoMarshaler) TracesSize(td Traces) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) ResourceSpansSize(td ResourceSpans) int {
	_ = "STUB: not implemented"
	return 0
}

func (e *ProtoMarshaler) ScopeSpansSize(td ScopeSpans) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) SpanSize(td Span) int { _ = "STUB: not implemented"; return 0 }

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalTraces(buf []byte) (Traces, error) {
	_ = "STUB: not implemented"
	return *new(Traces), nil
}

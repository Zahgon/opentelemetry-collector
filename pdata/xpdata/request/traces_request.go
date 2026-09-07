package request

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

func MarshalTraces(ctx context.Context, ld ptrace.Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalTraces(buf []byte) (context.Context, ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(ptrace.Traces), nil
}

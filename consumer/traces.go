package consumer

import (
	"context"

	"go.opentelemetry.io/collector/consumer/internal"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type Traces interface {
	internal.BaseConsumer

	ConsumeTraces(ctx context.Context, td ptrace.Traces) error
}

type ConsumeTracesFunc func(ctx context.Context, td ptrace.Traces) error

func (f ConsumeTracesFunc) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

type baseTraces struct {
	*internal.BaseImpl
	ConsumeTracesFunc
}

func NewTraces(consume ConsumeTracesFunc, options ...Option) (Traces, error) {
	_ = "STUB: not implemented"
	return *new(Traces), nil
}

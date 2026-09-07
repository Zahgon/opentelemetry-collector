package processorhelper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
)

type ProcessTracesFunc func(context.Context, ptrace.Traces) (ptrace.Traces, error)

type traces struct {
	component.StartFunc
	component.ShutdownFunc
	consumer.Traces
}

func NewTraces(
	_ context.Context,
	set processor.Settings,
	_ component.Config,
	nextConsumer consumer.Traces,
	tracesFunc ProcessTracesFunc,
	options ...Option,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

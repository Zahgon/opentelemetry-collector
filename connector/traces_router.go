package connector

import (
	"go.opentelemetry.io/collector/connector/internal"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pipeline"
)

type TracesRouterAndConsumer interface {
	consumer.Traces
	Consumer(...pipeline.ID) (consumer.Traces, error)
	PipelineIDs() []pipeline.ID
	privateFunc()
}

type tracesRouter struct {
	consumer.Traces
	internal.BaseRouter[consumer.Traces]
}

func NewTracesRouter(cm map[pipeline.ID]consumer.Traces) TracesRouterAndConsumer {
	_ = "STUB: not implemented"
	return *new(TracesRouterAndConsumer)
}

func (r *tracesRouter) privateFunc() { _ = "STUB: not implemented"; return }

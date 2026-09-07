package connector

import (
	"go.opentelemetry.io/collector/connector/internal"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pipeline"
)

type MetricsRouterAndConsumer interface {
	consumer.Metrics
	Consumer(...pipeline.ID) (consumer.Metrics, error)
	PipelineIDs() []pipeline.ID
	privateFunc()
}

type metricsRouter struct {
	consumer.Metrics
	internal.BaseRouter[consumer.Metrics]
}

func NewMetricsRouter(cm map[pipeline.ID]consumer.Metrics) MetricsRouterAndConsumer {
	_ = "STUB: not implemented"
	return *new(MetricsRouterAndConsumer)
}

func (r *metricsRouter) privateFunc() { _ = "STUB: not implemented"; return }

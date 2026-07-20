package processorhelper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor"
)

type ProcessMetricsFunc func(context.Context, pmetric.Metrics) (pmetric.Metrics, error)

type metrics struct {
	component.StartFunc
	component.ShutdownFunc
	consumer.Metrics
}

func NewMetrics(
	_ context.Context,
	set processor.Settings,
	_ component.Config,
	nextConsumer consumer.Metrics,
	metricsFunc ProcessMetricsFunc,
	options ...Option,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

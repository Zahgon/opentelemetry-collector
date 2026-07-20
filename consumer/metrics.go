package consumer

import (
	"context"

	"go.opentelemetry.io/collector/consumer/internal"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type Metrics interface {
	internal.BaseConsumer

	ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error
}

type ConsumeMetricsFunc func(ctx context.Context, md pmetric.Metrics) error

func (f ConsumeMetricsFunc) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

type baseMetrics struct {
	*internal.BaseImpl
	ConsumeMetricsFunc
}

func NewMetrics(consume ConsumeMetricsFunc, options ...Option) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

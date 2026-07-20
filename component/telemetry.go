package component

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type TelemetrySettings struct {
	Logger *zap.Logger

	TracerProvider trace.TracerProvider

	MeterProvider metric.MeterProvider

	Resource pcommon.Resource

	_ struct{}
}

package telemetrytest

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/service/telemetry"
)

func WithResource(res pcommon.Resource) telemetry.FactoryOption {
	_ = "STUB: not implemented"
	return *new(telemetry.FactoryOption)
}

func WithLogger(logger *zap.Logger, shutdownFunc component.ShutdownFunc) telemetry.FactoryOption {
	_ = "STUB: not implemented"
	return *new(telemetry.FactoryOption)
}

func WithMeterProvider(provider metric.MeterProvider) telemetry.FactoryOption {
	_ = "STUB: not implemented"
	return *new(telemetry.FactoryOption)
}

func WithTracerProvider(provider trace.TracerProvider) telemetry.FactoryOption {
	_ = "STUB: not implemented"
	return *new(telemetry.FactoryOption)
}

type ShutdownMeterProvider struct {
	metric.MeterProvider
	component.ShutdownFunc
}

type ShutdownTracerProvider struct {
	trace.TracerProvider
	component.ShutdownFunc
}

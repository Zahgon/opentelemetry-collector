package componenttest

import (
	"context"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"

	"go.opentelemetry.io/collector/component"
)

type TelemetryOption interface {
	apply(*telemetryOption)
}

type telemetryOption struct {
	metricOpts []sdkmetric.Option
	traceOpts  []sdktrace.TracerProviderOption
}

type telemetryOptionFunc func(*telemetryOption)

func (f telemetryOptionFunc) apply(o *telemetryOption) { _ = "STUB: not implemented"; return }

func WithMetricOptions(opts ...sdkmetric.Option) TelemetryOption {
	_ = "STUB: not implemented"
	return *new(TelemetryOption)
}

func WithTraceOptions(opts ...sdktrace.TracerProviderOption) TelemetryOption {
	_ = "STUB: not implemented"
	return *new(TelemetryOption)
}

type Telemetry struct {
	Reader        *sdkmetric.ManualReader
	SpanRecorder  *tracetest.SpanRecorder
	meterProvider *sdkmetric.MeterProvider
	traceProvider *sdktrace.TracerProvider
}

func NewTelemetry(opts ...TelemetryOption) *Telemetry { _ = "STUB: not implemented"; return nil }

func (tt *Telemetry) NewTelemetrySettings() component.TelemetrySettings {
	_ = "STUB: not implemented"
	return *new(component.TelemetrySettings)
}

func (tt *Telemetry) GetMetric(name string) (metricdata.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(metricdata.Metrics), nil
}

func (tt *Telemetry) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

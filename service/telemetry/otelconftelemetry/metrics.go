package otelconftelemetry

import (
	"context"

	noopmetric "go.opentelemetry.io/otel/metric/noop"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service/telemetry"
)

func createMeterProvider(
	ctx context.Context,
	set telemetry.MeterSettings,
	componentConfig component.Config,
) (telemetry.MeterProvider, error) {
	_ = "STUB: not implemented"
	return *new(telemetry.MeterProvider), nil
}

type noopMeterProvider struct {
	noopmetric.MeterProvider
}

func (noopMeterProvider) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

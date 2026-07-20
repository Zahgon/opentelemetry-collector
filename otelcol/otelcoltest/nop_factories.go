package otelcoltest

import (
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/service/telemetry"
)

func NopFactories() (otelcol.Factories, error) {
	_ = "STUB: not implemented"
	return *new(otelcol.Factories), nil
}

func nopTelemetryFactory() telemetry.Factory {
	_ = "STUB: not implemented"
	return *new(telemetry.Factory)
}

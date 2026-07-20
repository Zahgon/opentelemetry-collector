package otelconftelemetry

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service/telemetry"
)

func NewFactory() telemetry.Factory { _ = "STUB: not implemented"; return *new(telemetry.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func ptr[T any](v T) *T { _ = "STUB: not implemented"; return nil }

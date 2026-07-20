package otelconftelemetry

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service/telemetry"
)

func createLogger(
	ctx context.Context,
	set telemetry.LoggerSettings,
	componentConfig component.Config,
) (*zap.Logger, component.ShutdownFunc, error) {
	_ = "STUB: not implemented"
	return nil, *new(component.ShutdownFunc), nil
}

func warnLegacyResourceAttributes(logger *zap.Logger, cfg *Config) {
	_ = "STUB: not implemented"
	return
}

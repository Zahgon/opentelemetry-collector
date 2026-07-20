package debugexporter

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/xexporter"
)

var componentType = component.MustNewType("debug")

const (
	defaultSamplingInitial    = 2
	defaultSamplingThereafter = 1
)

func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTraces(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createMetrics(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createLogs(ctx context.Context, set exporter.Settings, config component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func createProfiles(ctx context.Context, set exporter.Settings, config component.Config) (xexporter.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xexporter.Profiles), nil
}

func createLogger(cfg *Config, logger *zap.Logger) *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

func createCustomLogger(exporterConfig *Config) *zap.Logger { _ = "STUB: not implemented"; return nil }

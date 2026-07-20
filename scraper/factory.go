package scraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
)

type Settings struct {
	ID component.ID

	component.TelemetrySettings

	BuildInfo component.BuildInfo

	_ struct{}
}

type Factory interface {
	component.Factory

	CreateLogs(ctx context.Context, set Settings, cfg component.Config) (Logs, error)

	CreateMetrics(ctx context.Context, set Settings, cfg component.Config) (Metrics, error)

	LogsStability() component.StabilityLevel

	MetricsStability() component.StabilityLevel

	unexportedFactoryFunc()
}

type FactoryOption interface {
	applyOption(o *factory)
}

var _ FactoryOption = (*factoryOptionFunc)(nil)

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) applyOption(o *factory) { _ = "STUB: not implemented"; return }

type factory struct {
	cfgType component.Type
	component.CreateDefaultConfigFunc
	createLogsFunc        CreateLogsFunc
	createMetricsFunc     CreateMetricsFunc
	logsStabilityLevel    component.StabilityLevel
	metricsStabilityLevel component.StabilityLevel
}

func (f *factory) Type() component.Type { _ = "STUB: not implemented"; return *new(component.Type) }

func (f *factory) unexportedFactoryFunc() { _ = "STUB: not implemented"; return }

func (f *factory) LogsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) MetricsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) CreateLogs(ctx context.Context, set Settings, cfg component.Config) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

func (f *factory) CreateMetrics(ctx context.Context, set Settings, cfg component.Config) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

type CreateLogsFunc func(context.Context, Settings, component.Config) (Logs, error)

type CreateMetricsFunc func(context.Context, Settings, component.Config) (Metrics, error)

func WithLogs(createLogs CreateLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetrics(createMetrics CreateMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func NewFactory(cfgType component.Type, createDefaultConfig component.CreateDefaultConfigFunc, options ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

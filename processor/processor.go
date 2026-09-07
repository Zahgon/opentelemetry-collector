package processor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/internal/componentalias"
)

type Traces interface {
	component.Component
	consumer.Traces
}

type Metrics interface {
	component.Component
	consumer.Metrics
}

type Logs interface {
	component.Component
	consumer.Logs
}

type Settings struct {
	ID component.ID

	component.TelemetrySettings

	BuildInfo component.BuildInfo

	_ struct{}
}

type Factory interface {
	component.Factory

	CreateTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Traces, error)

	TracesStability() component.StabilityLevel

	CreateMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Metrics, error)

	MetricsStability() component.StabilityLevel

	CreateLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Logs, error)

	LogsStability() component.StabilityLevel

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
	componentalias.TypeAliasHolder
	createTracesFunc      CreateTracesFunc
	tracesStabilityLevel  component.StabilityLevel
	createMetricsFunc     CreateMetricsFunc
	metricsStabilityLevel component.StabilityLevel
	createLogsFunc        CreateLogsFunc
	logsStabilityLevel    component.StabilityLevel
}

func (f *factory) Type() component.Type { _ = "STUB: not implemented"; return *new(component.Type) }

func (f *factory) unexportedFactoryFunc() { _ = "STUB: not implemented"; return }

func (f *factory) TracesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) MetricsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) LogsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) CreateTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Traces, error) {
	_ = "STUB: not implemented"
	return *new(Traces), nil
}

func (f *factory) CreateMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

func (f *factory) CreateLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

type CreateTracesFunc func(context.Context, Settings, component.Config, consumer.Traces) (Traces, error)

type CreateMetricsFunc func(context.Context, Settings, component.Config, consumer.Metrics) (Metrics, error)

type CreateLogsFunc func(context.Context, Settings, component.Config, consumer.Logs) (Logs, error)

func WithTraces(createTraces CreateTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetrics(createMetrics CreateMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogs(createLogs CreateLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func NewFactory(cfgType component.Type, createDefaultConfig component.CreateDefaultConfigFunc, options ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

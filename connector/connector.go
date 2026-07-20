package connector

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

	CreateDefaultConfig() component.Config

	CreateTracesToTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Traces, error)
	CreateTracesToMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Traces, error)
	CreateTracesToLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Traces, error)

	CreateMetricsToTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Metrics, error)
	CreateMetricsToMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Metrics, error)
	CreateMetricsToLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Metrics, error)

	CreateLogsToTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Logs, error)
	CreateLogsToMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Logs, error)
	CreateLogsToLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Logs, error)

	TracesToTracesStability() component.StabilityLevel
	TracesToMetricsStability() component.StabilityLevel
	TracesToLogsStability() component.StabilityLevel

	MetricsToTracesStability() component.StabilityLevel
	MetricsToMetricsStability() component.StabilityLevel
	MetricsToLogsStability() component.StabilityLevel

	LogsToTracesStability() component.StabilityLevel
	LogsToMetricsStability() component.StabilityLevel
	LogsToLogsStability() component.StabilityLevel

	unexportedFactoryFunc()
}

type FactoryOption interface {
	apply(o *factory)
}

var _ FactoryOption = (*factoryOptionFunc)(nil)

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) apply(o *factory) { _ = "STUB: not implemented"; return }

type CreateTracesToTracesFunc func(context.Context, Settings, component.Config, consumer.Traces) (Traces, error)

type CreateTracesToMetricsFunc func(context.Context, Settings, component.Config, consumer.Metrics) (Traces, error)

type CreateTracesToLogsFunc func(context.Context, Settings, component.Config, consumer.Logs) (Traces, error)

type CreateMetricsToTracesFunc func(context.Context, Settings, component.Config, consumer.Traces) (Metrics, error)

type CreateMetricsToMetricsFunc func(context.Context, Settings, component.Config, consumer.Metrics) (Metrics, error)

type CreateMetricsToLogsFunc func(context.Context, Settings, component.Config, consumer.Logs) (Metrics, error)

type CreateLogsToTracesFunc func(context.Context, Settings, component.Config, consumer.Traces) (Logs, error)

type CreateLogsToMetricsFunc func(context.Context, Settings, component.Config, consumer.Metrics) (Logs, error)

type CreateLogsToLogsFunc func(context.Context, Settings, component.Config, consumer.Logs) (Logs, error)

func WithTracesToTraces(createTracesToTraces CreateTracesToTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithTracesToMetrics(createTracesToMetrics CreateTracesToMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithTracesToLogs(createTracesToLogs CreateTracesToLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetricsToTraces(createMetricsToTraces CreateMetricsToTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetricsToMetrics(createMetricsToMetrics CreateMetricsToMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetricsToLogs(createMetricsToLogs CreateMetricsToLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogsToTraces(createLogsToTraces CreateLogsToTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogsToMetrics(createLogsToMetrics CreateLogsToMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogsToLogs(createLogsToLogs CreateLogsToLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type factory struct {
	cfgType component.Type
	component.CreateDefaultConfigFunc
	componentalias.TypeAliasHolder

	createTracesToTracesFunc  CreateTracesToTracesFunc
	createTracesToMetricsFunc CreateTracesToMetricsFunc
	createTracesToLogsFunc    CreateTracesToLogsFunc

	createMetricsToTracesFunc  CreateMetricsToTracesFunc
	createMetricsToMetricsFunc CreateMetricsToMetricsFunc
	createMetricsToLogsFunc    CreateMetricsToLogsFunc

	createLogsToTracesFunc  CreateLogsToTracesFunc
	createLogsToMetricsFunc CreateLogsToMetricsFunc
	createLogsToLogsFunc    CreateLogsToLogsFunc

	tracesToTracesStabilityLevel  component.StabilityLevel
	tracesToMetricsStabilityLevel component.StabilityLevel
	tracesToLogsStabilityLevel    component.StabilityLevel

	metricsToTracesStabilityLevel  component.StabilityLevel
	metricsToMetricsStabilityLevel component.StabilityLevel
	metricsToLogsStabilityLevel    component.StabilityLevel

	logsToTracesStabilityLevel  component.StabilityLevel
	logsToMetricsStabilityLevel component.StabilityLevel
	logsToLogsStabilityLevel    component.StabilityLevel
}

func (f *factory) Type() component.Type { _ = "STUB: not implemented"; return *new(component.Type) }

func (f *factory) unexportedFactoryFunc() { _ = "STUB: not implemented"; return }

func (f *factory) TracesToTracesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) TracesToMetricsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) TracesToLogsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) MetricsToTracesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) MetricsToMetricsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) MetricsToLogsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) LogsToTracesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) LogsToMetricsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) LogsToLogsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) CreateTracesToTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Traces, error) {
	_ = "STUB: not implemented"
	return *new(Traces), nil
}

func (f *factory) CreateTracesToMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Traces, error) {
	_ = "STUB: not implemented"
	return *new(Traces), nil
}

func (f *factory) CreateTracesToLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Traces, error) {
	_ = "STUB: not implemented"
	return *new(Traces), nil
}

func (f *factory) CreateMetricsToTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

func (f *factory) CreateMetricsToMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

func (f *factory) CreateMetricsToLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

func (f *factory) CreateLogsToTraces(ctx context.Context, set Settings, cfg component.Config, next consumer.Traces) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

func (f *factory) CreateLogsToMetrics(ctx context.Context, set Settings, cfg component.Config, next consumer.Metrics) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

func (f *factory) CreateLogsToLogs(ctx context.Context, set Settings, cfg component.Config, next consumer.Logs) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

func NewFactory(cfgType component.Type, createDefaultConfig component.CreateDefaultConfigFunc, options ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

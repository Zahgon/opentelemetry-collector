package xconnector

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/internal/componentalias"
)

type Factory interface {
	connector.Factory

	CreateTracesToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (connector.Traces, error)
	CreateMetricsToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (connector.Metrics, error)
	CreateLogsToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (connector.Logs, error)

	TracesToProfilesStability() component.StabilityLevel
	MetricsToProfilesStability() component.StabilityLevel
	LogsToProfilesStability() component.StabilityLevel

	CreateProfilesToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (Profiles, error)
	CreateProfilesToTraces(ctx context.Context, set connector.Settings, cfg component.Config, next consumer.Traces) (Profiles, error)
	CreateProfilesToMetrics(ctx context.Context, set connector.Settings, cfg component.Config, next consumer.Metrics) (Profiles, error)
	CreateProfilesToLogs(ctx context.Context, set connector.Settings, cfg component.Config, next consumer.Logs) (Profiles, error)

	ProfilesToProfilesStability() component.StabilityLevel
	ProfilesToTracesStability() component.StabilityLevel
	ProfilesToMetricsStability() component.StabilityLevel
	ProfilesToLogsStability() component.StabilityLevel
}

type Profiles interface {
	component.Component
	xconsumer.Profiles
}

type CreateTracesToProfilesFunc func(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (connector.Traces, error)

type CreateMetricsToProfilesFunc func(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (connector.Metrics, error)

type CreateLogsToProfilesFunc func(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (connector.Logs, error)

type CreateProfilesToProfilesFunc func(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (Profiles, error)

type CreateProfilesToTracesFunc func(context.Context, connector.Settings, component.Config, consumer.Traces) (Profiles, error)

type CreateProfilesToMetricsFunc func(context.Context, connector.Settings, component.Config, consumer.Metrics) (Profiles, error)

type CreateProfilesToLogsFunc func(context.Context, connector.Settings, component.Config, consumer.Logs) (Profiles, error)

type FactoryOption interface {
	applyOption(o *factory)
}

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) applyOption(o *factory) { _ = "STUB: not implemented"; return }

func WithTracesToTraces(createTracesToTraces connector.CreateTracesToTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithTracesToMetrics(createTracesToMetrics connector.CreateTracesToMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithTracesToLogs(createTracesToLogs connector.CreateTracesToLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetricsToTraces(createMetricsToTraces connector.CreateMetricsToTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetricsToMetrics(createMetricsToMetrics connector.CreateMetricsToMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetricsToLogs(createMetricsToLogs connector.CreateMetricsToLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogsToTraces(createLogsToTraces connector.CreateLogsToTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogsToMetrics(createLogsToMetrics connector.CreateLogsToMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogsToLogs(createLogsToLogs connector.CreateLogsToLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithTracesToProfiles(createTracesToProfiles CreateTracesToProfilesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetricsToProfiles(createMetricsToProfiles CreateMetricsToProfilesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogsToProfiles(createLogsToProfiles CreateLogsToProfilesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithProfilesToProfiles(createProfilesToProfiles CreateProfilesToProfilesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithProfilesToTraces(createProfilesToTraces CreateProfilesToTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithProfilesToMetrics(createProfilesToMetrics CreateProfilesToMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithProfilesToLogs(createProfilesToLogs CreateProfilesToLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithDeprecatedTypeAlias(alias component.Type) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type factory struct {
	connector.Factory
	componentalias.TypeAliasHolder
	opts []connector.FactoryOption

	createTracesToProfilesFunc  CreateTracesToProfilesFunc
	createMetricsToProfilesFunc CreateMetricsToProfilesFunc
	createLogsToProfilesFunc    CreateLogsToProfilesFunc

	createProfilesToProfilesFunc CreateProfilesToProfilesFunc
	createProfilesToTracesFunc   CreateProfilesToTracesFunc
	createProfilesToMetricsFunc  CreateProfilesToMetricsFunc
	createProfilesToLogsFunc     CreateProfilesToLogsFunc

	tracesToProfilesStabilityLevel  component.StabilityLevel
	metricsToProfilesStabilityLevel component.StabilityLevel
	logsToProfilesStabilityLevel    component.StabilityLevel

	profilesToProfilesStabilityLevel component.StabilityLevel
	profilesToTracesStabilityLevel   component.StabilityLevel
	profilesToMetricsStabilityLevel  component.StabilityLevel
	profilesToLogsStabilityLevel     component.StabilityLevel
}

func (f *factory) TracesToProfilesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) MetricsToProfilesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) LogsToProfilesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) ProfilesToProfilesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) ProfilesToTracesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) ProfilesToMetricsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) ProfilesToLogsStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) CreateTracesToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

func (f *factory) CreateMetricsToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

func (f *factory) CreateLogsToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

func (f *factory) CreateProfilesToProfiles(ctx context.Context, set connector.Settings, cfg component.Config, next xconsumer.Profiles) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

func (f *factory) CreateProfilesToTraces(ctx context.Context, set connector.Settings, cfg component.Config, next consumer.Traces) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

func (f *factory) CreateProfilesToMetrics(ctx context.Context, set connector.Settings, cfg component.Config, next consumer.Metrics) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

func (f *factory) CreateProfilesToLogs(ctx context.Context, set connector.Settings, cfg component.Config, next consumer.Logs) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

func NewFactory(cfgType component.Type, createDefaultConfig component.CreateDefaultConfigFunc, options ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

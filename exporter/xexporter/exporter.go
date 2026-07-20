package xexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/internal/componentalias"
)

type Profiles interface {
	component.Component
	xconsumer.Profiles
}

type Factory interface {
	exporter.Factory

	CreateProfiles(ctx context.Context, set exporter.Settings, cfg component.Config) (Profiles, error)

	ProfilesStability() component.StabilityLevel
}

type FactoryOption interface {
	applyOption(o *factory)
}

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) applyOption(o *factory) { _ = "STUB: not implemented"; return }

type CreateProfilesFunc func(context.Context, exporter.Settings, component.Config) (Profiles, error)

func WithTraces(createTraces exporter.CreateTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetrics(createMetrics exporter.CreateMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogs(createLogs exporter.CreateLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithProfiles(createProfiles CreateProfilesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithDeprecatedTypeAlias(alias component.Type) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type factory struct {
	exporter.Factory
	componentalias.TypeAliasHolder
	opts                   []exporter.FactoryOption
	createProfilesFunc     CreateProfilesFunc
	profilesStabilityLevel component.StabilityLevel
}

func (f *factory) ProfilesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) CreateProfiles(ctx context.Context, set exporter.Settings, cfg component.Config) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

func NewFactory(cfgType component.Type, createDefaultConfig component.CreateDefaultConfigFunc, options ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

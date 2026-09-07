package xreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/internal/componentalias"
	"go.opentelemetry.io/collector/receiver"
)

type Profiles interface {
	component.Component
}

type Factory interface {
	receiver.Factory

	CreateProfiles(ctx context.Context, set receiver.Settings, cfg component.Config, next xconsumer.Profiles) (Profiles, error)

	ProfilesStability() component.StabilityLevel
}

type CreateProfilesFunc func(context.Context, receiver.Settings, component.Config, xconsumer.Profiles) (Profiles, error)

type FactoryOption interface {
	applyOption(o *factory)
}

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) applyOption(o *factory) { _ = "STUB: not implemented"; return }

type factory struct {
	receiver.Factory
	componentalias.TypeAliasHolder
	opts                   []receiver.FactoryOption
	createProfilesFunc     CreateProfilesFunc
	profilesStabilityLevel component.StabilityLevel
}

func (f *factory) ProfilesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) CreateProfiles(ctx context.Context, set receiver.Settings, cfg component.Config, next xconsumer.Profiles) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

func WithTraces(createTraces receiver.CreateTracesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetrics(createMetrics receiver.CreateMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithLogs(createLogs receiver.CreateLogsFunc, sl component.StabilityLevel) FactoryOption {
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

func NewFactory(cfgType component.Type, createDefaultConfig component.CreateDefaultConfigFunc, options ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

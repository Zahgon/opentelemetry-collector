package xscraper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/scraper"
)

type Factory interface {
	scraper.Factory

	CreateProfiles(ctx context.Context, set scraper.Settings, cfg component.Config) (Profiles, error)

	ProfilesStability() component.StabilityLevel
}

type FactoryOption interface {
	applyOption(o *factory)
}

var _ FactoryOption = (*factoryOptionFunc)(nil)

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) applyOption(o *factory) { _ = "STUB: not implemented"; return }

type factory struct {
	scraper.Factory
	opts []scraper.FactoryOption

	createProfilesFunc     CreateProfilesFunc
	profilesStabilityLevel component.StabilityLevel
}

func (f *factory) ProfilesStability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) CreateProfiles(ctx context.Context, set scraper.Settings, cfg component.Config) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

func WithLogs(createLogs scraper.CreateLogsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func WithMetrics(createMetrics scraper.CreateMetricsFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

type CreateProfilesFunc func(context.Context, scraper.Settings, component.Config) (Profiles, error)

func WithProfiles(createProfiles CreateProfilesFunc, sl component.StabilityLevel) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func NewFactory(cfgType component.Type, createDefaultConfig component.CreateDefaultConfigFunc, options ...FactoryOption) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

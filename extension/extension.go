package extension

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/internal/componentalias"
)

type Extension interface {
	component.Component
}

type Settings struct {
	ID component.ID

	component.TelemetrySettings

	BuildInfo component.BuildInfo

	_ struct{}
}

type CreateFunc func(context.Context, Settings, component.Config) (Extension, error)

type Factory interface {
	component.Factory

	Create(ctx context.Context, set Settings, cfg component.Config) (Extension, error)

	Stability() component.StabilityLevel

	unexportedFactoryFunc()
}

type factory struct {
	cfgType component.Type
	component.CreateDefaultConfigFunc
	componentalias.TypeAliasHolder
	createFunc         CreateFunc
	extensionStability component.StabilityLevel
}

func (f *factory) Type() component.Type { _ = "STUB: not implemented"; return *new(component.Type) }

func (f *factory) unexportedFactoryFunc() { _ = "STUB: not implemented"; return }

func (f *factory) Stability() component.StabilityLevel {
	_ = "STUB: not implemented"
	return *new(component.StabilityLevel)
}

func (f *factory) Create(ctx context.Context, set Settings, cfg component.Config) (Extension, error) {
	_ = "STUB: not implemented"
	return *new(Extension), nil
}

func NewFactory(
	cfgType component.Type,
	createDefaultConfig component.CreateDefaultConfigFunc,
	createServiceExtension CreateFunc,
	sl component.StabilityLevel,
) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

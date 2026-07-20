package xextension

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/internal/componentalias"
)

type Factory interface {
	extension.Factory
}

type FactoryOption interface {
	applyOption(o *factory)
}

type factoryOptionFunc func(*factory)

func (f factoryOptionFunc) applyOption(o *factory) { _ = "STUB: not implemented"; return }

type factory struct {
	extension.Factory
	componentalias.TypeAliasHolder
}

func WithDeprecatedTypeAlias(alias component.Type) FactoryOption {
	_ = "STUB: not implemented"
	return *new(FactoryOption)
}

func NewFactory(
	cfgType component.Type,
	createDefaultConfig component.CreateDefaultConfigFunc,
	createServiceExtension extension.CreateFunc,
	sl component.StabilityLevel,
	options ...FactoryOption,
) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

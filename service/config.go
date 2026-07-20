package service

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service/extensions"
	"go.opentelemetry.io/collector/service/pipelines"
)

type Config struct {
	Telemetry component.Config `mapstructure:"telemetry"`

	Extensions extensions.Config `mapstructure:"extensions,omitempty"`

	Pipelines pipelines.Config `mapstructure:"pipelines"`

	_ struct{}
}

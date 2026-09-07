package otelcol

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
)

var (
	errMissingExporters       = errors.New("no exporter configuration specified in config")
	errMissingReceivers       = errors.New("no receiver configuration specified in config")
	errEmptyConfigurationFile = errors.New("empty configuration file")
)

type Config struct {
	Receivers map[component.ID]component.Config `mapstructure:"receivers"`

	Exporters map[component.ID]component.Config `mapstructure:"exporters"`

	Processors map[component.ID]component.Config `mapstructure:"processors"`

	Connectors map[component.ID]component.Config `mapstructure:"connectors"`

	Extensions map[component.ID]component.Config `mapstructure:"extensions"`

	Service service.Config `mapstructure:"service"`

	_ struct{}
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

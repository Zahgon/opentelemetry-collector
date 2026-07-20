package otlpreceiver

import (
	"encoding"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configoptional"
)

type SanitizedURLPath string

var _ encoding.TextUnmarshaler = (*SanitizedURLPath)(nil)

func (s *SanitizedURLPath) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

type HTTPConfig struct {
	ServerConfig confighttp.ServerConfig `mapstructure:",squash"`

	TracesURLPath SanitizedURLPath `mapstructure:"traces_url_path,omitempty"`

	MetricsURLPath SanitizedURLPath `mapstructure:"metrics_url_path,omitempty"`

	LogsURLPath SanitizedURLPath `mapstructure:"logs_url_path,omitempty"`

	_ struct{}
}

type Protocols struct {
	GRPC configoptional.Optional[configgrpc.ServerConfig] `mapstructure:"grpc"`
	HTTP configoptional.Optional[HTTPConfig]              `mapstructure:"http"`

	_ struct{}
}

type Config struct {
	Protocols Protocols `mapstructure:"protocols"`

	_ struct{}
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

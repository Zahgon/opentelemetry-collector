package zpagesextension

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

type Config struct {
	confighttp.ServerConfig `mapstructure:",squash"`

	Expvar ExpvarConfig `mapstructure:"expvar"`

	_ struct{}
}

type ExpvarConfig struct {
	Enabled bool `mapstructure:"enabled"`

	_ struct{}
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

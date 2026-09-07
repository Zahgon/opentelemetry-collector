package configretry

import (
	"time"
)

type BackOffConfig struct {
	Enabled bool `mapstructure:"enabled"`

	InitialInterval time.Duration `mapstructure:"initial_interval"`

	MaxElapsedTime time.Duration `mapstructure:"max_elapsed_time"`

	MaxInterval time.Duration `mapstructure:"max_interval"`

	Multiplier float64 `mapstructure:"multiplier"`

	RandomizationFactor float64 `mapstructure:"randomization_factor"`

	_ struct{}
}

func (c *BackOffConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func NewDefaultBackOffConfig() BackOffConfig { _ = "STUB: not implemented"; return *new(BackOffConfig) }

package configretry

import (
	"time"
)

func NewDefaultBackOffConfig() BackOffConfig { _ = "STUB: not implemented"; return *new(BackOffConfig) }

type BackOffConfig struct {
	Enabled bool `mapstructure:"enabled"`

	InitialInterval time.Duration `mapstructure:"initial_interval"`

	RandomizationFactor float64 `mapstructure:"randomization_factor"`

	Multiplier float64 `mapstructure:"multiplier"`

	MaxInterval time.Duration `mapstructure:"max_interval"`

	MaxElapsedTime time.Duration `mapstructure:"max_elapsed_time"`

	_ struct{}
}

func (bs *BackOffConfig) Validate() error { _ = "STUB: not implemented"; return nil }

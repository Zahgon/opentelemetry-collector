package batchprocessor

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

type Config struct {
	Timeout time.Duration `mapstructure:"timeout"`

	SendBatchSize uint32 `mapstructure:"send_batch_size"`

	SendBatchMaxSize uint32 `mapstructure:"send_batch_max_size"`

	MetadataKeys []string `mapstructure:"metadata_keys"`

	MetadataCardinalityLimit uint32 `mapstructure:"metadata_cardinality_limit"`

	_ struct{}
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

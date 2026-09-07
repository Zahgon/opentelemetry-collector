package debugexporter

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var supportedLevels map[configtelemetry.Level]struct{} = map[configtelemetry.Level]struct{}{
	configtelemetry.LevelBasic:    {},
	configtelemetry.LevelNormal:   {},
	configtelemetry.LevelDetailed: {},
}

type Config struct {
	Verbosity configtelemetry.Level `mapstructure:"verbosity,omitempty"`

	SamplingInitial int `mapstructure:"sampling_initial"`

	SamplingThereafter int `mapstructure:"sampling_thereafter"`

	UseInternalLogger bool `mapstructure:"use_internal_logger"`

	OutputPaths []string `mapstructure:"output_paths"`

	QueueConfig configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`

	_ struct{}
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

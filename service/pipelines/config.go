package pipelines

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/service/internal/metadata"
)

var (
	errMissingServicePipelines         = errors.New("service must have at least one pipeline")
	errMissingServicePipelineReceivers = errors.New("must have at least one receiver")
	errMissingServicePipelineExporters = errors.New("must have at least one exporter")

	AllowNoPipelines = metadata.ServiceAllowNoPipelinesFeatureGate
)

type Config map[pipeline.ID]*PipelineConfig

func (cfg Config) Validate() error { _ = "STUB: not implemented"; return nil }

type PipelineConfig struct {
	Receivers  []component.ID `mapstructure:"receivers"`
	Processors []component.ID `mapstructure:"processors"`
	Exporters  []component.ID `mapstructure:"exporters"`
}

func (cfg *PipelineConfig) Validate() error { _ = "STUB: not implemented"; return nil }

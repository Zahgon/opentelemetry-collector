package otlphttpexporter

import (
	"encoding"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

type EncodingType string

const (
	EncodingProto EncodingType = "proto"
	EncodingJSON  EncodingType = "json"
)

var _ encoding.TextUnmarshaler = (*EncodingType)(nil)

func (e *EncodingType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

type Config struct {
	ClientConfig confighttp.ClientConfig                                  `mapstructure:",squash"`
	QueueConfig  configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	RetryConfig  configretry.BackOffConfig                                `mapstructure:"retry_on_failure"`

	TracesEndpoint string `mapstructure:"traces_endpoint"`

	MetricsEndpoint string `mapstructure:"metrics_endpoint"`

	LogsEndpoint string `mapstructure:"logs_endpoint"`

	ProfilesEndpoint string `mapstructure:"profiles_endpoint"`

	Encoding EncodingType `mapstructure:"encoding"`
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

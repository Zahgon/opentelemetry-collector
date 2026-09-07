package otelconftelemetry

import (
	"go.opentelemetry.io/collector/service/telemetry/otelconftelemetry/internal/migration"
)

type Config struct {
	Logs    LogsConfig    `mapstructure:"logs"`
	Metrics MetricsConfig `mapstructure:"metrics"`
	Traces  TracesConfig  `mapstructure:"traces,omitempty"`

	Resource ResourceConfig `mapstructure:"resource,omitempty"`
}

type LogsConfig = migration.LogsConfigV030

type LogsSamplingConfig = migration.LogsSamplingConfig

type MetricsConfig = migration.MetricsConfigV030

type TracesConfig = migration.TracesConfigV030

type ResourceConfig = migration.ResourceConfigV030

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

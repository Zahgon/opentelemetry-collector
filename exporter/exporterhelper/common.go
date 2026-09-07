package exporterhelper

import (
	"go.opentelemetry.io/otel/attribute"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal"
)

type Option = internal.Option

func WithStart(start component.StartFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithShutdown(shutdown component.ShutdownFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTimeout(timeoutConfig TimeoutConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRetry(config configretry.BackOffConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCapabilities(capabilities consumer.Capabilities) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithAttrs(attrs ...attribute.KeyValue) Option { _ = "STUB: not implemented"; return *new(Option) }

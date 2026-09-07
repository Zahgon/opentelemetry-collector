package otelcol

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/service/telemetry"
)

type Factories struct {
	Receivers map[component.Type]receiver.Factory

	Processors map[component.Type]processor.Factory

	Exporters map[component.Type]exporter.Factory

	Extensions map[component.Type]extension.Factory

	Connectors map[component.Type]connector.Factory

	Telemetry telemetry.Factory

	ReceiverModules map[component.Type]string

	ProcessorModules map[component.Type]string

	ExporterModules map[component.Type]string

	ExtensionModules map[component.Type]string

	ConnectorModules map[component.Type]string
}

func MakeFactoryMap[T component.Factory](factories ...T) (map[component.Type]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

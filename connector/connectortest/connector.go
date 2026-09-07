package connectortest

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/xconnector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/consumer/xconsumer"
)

var NopType = component.MustNewType("nop")

func NewNopSettings(typ component.Type) connector.Settings {
	_ = "STUB: not implemented"
	return *new(connector.Settings)
}

type nopConfig struct{}

func NewNopFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

func createTracesToTracesConnector(context.Context, connector.Settings, component.Config, consumer.Traces) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

func createTracesToMetricsConnector(context.Context, connector.Settings, component.Config, consumer.Metrics) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

func createTracesToLogsConnector(context.Context, connector.Settings, component.Config, consumer.Logs) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

func createTracesToProfilesConnector(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

func createMetricsToTracesConnector(context.Context, connector.Settings, component.Config, consumer.Traces) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

func createMetricsToMetricsConnector(context.Context, connector.Settings, component.Config, consumer.Metrics) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

func createMetricsToLogsConnector(context.Context, connector.Settings, component.Config, consumer.Logs) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

func createMetricsToProfilesConnector(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

func createLogsToTracesConnector(context.Context, connector.Settings, component.Config, consumer.Traces) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

func createLogsToMetricsConnector(context.Context, connector.Settings, component.Config, consumer.Metrics) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

func createLogsToLogsConnector(context.Context, connector.Settings, component.Config, consumer.Logs) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

func createLogsToProfilesConnector(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

func createProfilesToTracesConnector(context.Context, connector.Settings, component.Config, consumer.Traces) (xconnector.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xconnector.Profiles), nil
}

func createProfilesToMetricsConnector(context.Context, connector.Settings, component.Config, consumer.Metrics) (xconnector.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xconnector.Profiles), nil
}

func createProfilesToLogsConnector(context.Context, connector.Settings, component.Config, consumer.Logs) (xconnector.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xconnector.Profiles), nil
}

func createProfilesToProfilesConnector(context.Context, connector.Settings, component.Config, xconsumer.Profiles) (xconnector.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xconnector.Profiles), nil
}

type nopConnector struct {
	component.StartFunc
	component.ShutdownFunc
	consumertest.Consumer
}

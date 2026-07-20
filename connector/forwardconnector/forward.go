package forwardconnector

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/xconnector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
)

func NewFactory() xconnector.Factory { _ = "STUB: not implemented"; return *new(xconnector.Factory) }

type Config struct{}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesToTraces(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Traces,
) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

func createMetricsToMetrics(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Metrics,
) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

func createLogsToLogs(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Logs,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

func createProfilesToProfiles(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer xconsumer.Profiles,
) (xconnector.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xconnector.Profiles), nil
}

type forward struct {
	consumer.Traces
	consumer.Metrics
	consumer.Logs
	xconsumer.Profiles
	component.StartFunc
	component.ShutdownFunc
}

func (c *forward) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

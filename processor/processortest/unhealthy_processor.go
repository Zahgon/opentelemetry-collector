package processortest

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/processor"
)

func NewUnhealthyProcessorFactory() processor.Factory {
	_ = "STUB: not implemented"
	return *new(processor.Factory)
}

func createUnhealthyTraces(_ context.Context, set processor.Settings, _ component.Config, _ consumer.Traces) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func createUnhealthyMetrics(_ context.Context, set processor.Settings, _ component.Config, _ consumer.Metrics) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createUnhealthyLogs(_ context.Context, set processor.Settings, _ component.Config, _ consumer.Logs) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

type unhealthy struct {
	component.StartFunc
	component.ShutdownFunc
	consumertest.Consumer
	telemetry component.TelemetrySettings
}

func (p unhealthy) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

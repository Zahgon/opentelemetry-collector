package processortest

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/xprocessor"
)

var NopType = component.MustNewType("nop")

func NewNopSettings(typ component.Type) processor.Settings {
	_ = "STUB: not implemented"
	return *new(processor.Settings)
}

func NewNopFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createTraces(context.Context, processor.Settings, component.Config, consumer.Traces) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func createMetrics(context.Context, processor.Settings, component.Config, consumer.Metrics) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createLogs(context.Context, processor.Settings, component.Config, consumer.Logs) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func createProfiles(context.Context, processor.Settings, component.Config, xconsumer.Profiles) (xprocessor.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xprocessor.Profiles), nil
}

type nopConfig struct{}

var nopInstance = &nop{
	Consumer: consumertest.NewNop(),
}

type nop struct {
	component.StartFunc
	component.ShutdownFunc
	consumertest.Consumer
}

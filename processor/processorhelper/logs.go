package processorhelper

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
)

type ProcessLogsFunc func(context.Context, plog.Logs) (plog.Logs, error)

type logs struct {
	component.StartFunc
	component.ShutdownFunc
	consumer.Logs
}

func NewLogs(
	_ context.Context,
	set processor.Settings,
	_ component.Config,
	nextConsumer consumer.Logs,
	logsFunc ProcessLogsFunc,
	options ...Option,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

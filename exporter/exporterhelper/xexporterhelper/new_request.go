package xexporterhelper

import (
	"context"

	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/queuebatch"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func NewLogsRequest(
	ctx context.Context,
	set exporter.Settings,
	converter RequestConverterFunc[plog.Logs],
	pusher RequestConsumeFunc,
	options ...exporterhelper.Option,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

func NewMetricsRequest(
	ctx context.Context,
	set exporter.Settings,
	converter RequestConverterFunc[pmetric.Metrics],
	pusher RequestConsumeFunc,
	options ...exporterhelper.Option,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func NewTracesRequest(
	ctx context.Context,
	set exporter.Settings,
	converter RequestConverterFunc[ptrace.Traces],
	pusher RequestConsumeFunc,
	options ...exporterhelper.Option,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

type QueueBatchSettings = queuebatch.Settings[Request]

func NewMetricsQueueBatchSettings() QueueBatchSettings {
	_ = "STUB: not implemented"
	return *new(QueueBatchSettings)
}

func NewLogsQueueBatchSettings() QueueBatchSettings {
	_ = "STUB: not implemented"
	return *new(QueueBatchSettings)
}

func NewTracesQueueBatchSettings() QueueBatchSettings {
	_ = "STUB: not implemented"
	return *new(QueueBatchSettings)
}

func WithQueueBatch(cfg configoptional.Optional[exporterhelper.QueueBatchConfig], set QueueBatchSettings) exporterhelper.Option {
	_ = "STUB: not implemented"
	return *new(exporterhelper.Option)
}

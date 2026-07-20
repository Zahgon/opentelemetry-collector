package connector

import (
	"go.opentelemetry.io/collector/connector/internal"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pipeline"
)

type LogsRouterAndConsumer interface {
	consumer.Logs
	Consumer(...pipeline.ID) (consumer.Logs, error)
	PipelineIDs() []pipeline.ID
	privateFunc()
}

type logsRouter struct {
	consumer.Logs
	internal.BaseRouter[consumer.Logs]
}

func NewLogsRouter(cm map[pipeline.ID]consumer.Logs) LogsRouterAndConsumer {
	_ = "STUB: not implemented"
	return *new(LogsRouterAndConsumer)
}

func (r *logsRouter) PipelineIDs() []pipeline.ID { _ = "STUB: not implemented"; return nil }

func (r *logsRouter) Consumer(pipelineIDs ...pipeline.ID) (consumer.Logs, error) {
	_ = "STUB: not implemented"
	return *new(consumer.Logs), nil
}

func (r *logsRouter) privateFunc() { _ = "STUB: not implemented"; return }

package consumer

import (
	"context"

	"go.opentelemetry.io/collector/consumer/internal"
	"go.opentelemetry.io/collector/pdata/plog"
)

type Logs interface {
	internal.BaseConsumer

	ConsumeLogs(ctx context.Context, ld plog.Logs) error
}

type ConsumeLogsFunc func(ctx context.Context, ld plog.Logs) error

func (f ConsumeLogsFunc) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

type baseLogs struct {
	*internal.BaseImpl
	ConsumeLogsFunc
}

func NewLogs(consume ConsumeLogsFunc, options ...Option) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

package exporterhelper

import (
	"context"

	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/queue"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/queuebatch"
)

func WithQueue(config configoptional.Optional[QueueBatchConfig]) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type QueueBatchConfig = queuebatch.Config

type BatchConfig = queuebatch.BatchConfig

type QueueBatchEncoding[T any] interface {
	Marshal(context.Context, T) ([]byte, error)

	Unmarshal([]byte) (context.Context, T, error)
}

var ErrQueueIsFull = queue.ErrQueueIsFull

var NewDefaultQueueConfig = internal.NewDefaultQueueConfig

package batchprocessor

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
)

var errTooManyBatchers = consumererror.NewPermanent(errors.New("too many batcher metadata-value combinations"))

type batchProcessor[T any] struct {
	logger           *zap.Logger
	timeout          time.Duration
	sendBatchSize    int
	sendBatchMaxSize int

	batchFunc func() batch[T]

	shutdownC  chan struct{}
	goroutines sync.WaitGroup

	telemetry *batchProcessorTelemetry

	batcher batcher[T]
}

type batcher[T any] interface {
	start(ctx context.Context) error

	consume(ctx context.Context, data T) error

	currentMetadataCardinality() int
}

type shard[T any] struct {
	processor *batchProcessor[T]

	exportCtx context.Context

	timer *time.Timer

	newItem chan T

	batch batch[T]
}

type batch[T any] interface {
	export(ctx context.Context, req T) error

	split(sendBatchMaxSize int) (sentBatchSize int, req T)

	itemCount() int

	add(item T)

	sizeBytes(item T) int
}

func newBatchProcessor[T any](set processor.Settings, cfg *Config, batchFunc func() batch[T]) (*batchProcessor[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bp *batchProcessor[T]) newShard(md map[string][]string) *shard[T] {
	_ = "STUB: not implemented"
	return nil
}

func (bp *batchProcessor[T]) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (bp *batchProcessor[T]) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (bp *batchProcessor[T]) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (b *shard[T]) start() { _ = "STUB: not implemented"; return }

func (b *shard[T]) startLoop() { _ = "STUB: not implemented"; return }

func (b *shard[T]) processItem(item T) { _ = "STUB: not implemented"; return }

func (b *shard[T]) hasTimer() bool { _ = "STUB: not implemented"; return false }

func (b *shard[T]) stopTimer() { _ = "STUB: not implemented"; return }

func (b *shard[T]) resetTimer() { _ = "STUB: not implemented"; return }

func (b *shard[T]) sendItems(trigger trigger) { _ = "STUB: not implemented"; return }

type singleShardBatcher[T any] struct {
	processor *batchProcessor[T]
	single    *shard[T]
}

func (sb *singleShardBatcher[T]) start(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (sb *singleShardBatcher[T]) consume(_ context.Context, data T) error {
	_ = "STUB: not implemented"
	return nil
}

func (sb *singleShardBatcher[T]) currentMetadataCardinality() int {
	_ = "STUB: not implemented"
	return 0
}

type multiShardBatcher[T any] struct {
	metadataKeys []string

	metadataLimit int

	processor *batchProcessor[T]
	batchers  sync.Map

	lock sync.Mutex
	size int
}

func (mb *multiShardBatcher[T]) start(context.Context) error { _ = "STUB: not implemented"; return nil }

func (mb *multiShardBatcher[T]) consume(ctx context.Context, data T) error {
	_ = "STUB: not implemented"
	return nil
}

func (mb *multiShardBatcher[T]) currentMetadataCardinality() int {
	_ = "STUB: not implemented"
	return 0
}

type tracesBatchProcessor struct {
	*batchProcessor[ptrace.Traces]
}

func newTracesBatchProcessor(set processor.Settings, next consumer.Traces, cfg *Config) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func (t *tracesBatchProcessor) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

type metricsBatchProcessor struct {
	*batchProcessor[pmetric.Metrics]
}

func newMetricsBatchProcessor(set processor.Settings, next consumer.Metrics, cfg *Config) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func (m *metricsBatchProcessor) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

type logsBatchProcessor struct {
	*batchProcessor[plog.Logs]
}

func newLogsBatchProcessor(set processor.Settings, next consumer.Logs, cfg *Config) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func (l *logsBatchProcessor) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

type batchTraces struct {
	nextConsumer consumer.Traces
	traceData    ptrace.Traces
	spanCount    int
	sizer        ptrace.Sizer
}

func newBatchTraces(nextConsumer consumer.Traces) *batchTraces {
	_ = "STUB: not implemented"
	return nil
}

func (bt *batchTraces) add(td ptrace.Traces) { _ = "STUB: not implemented"; return }

func (bt *batchTraces) sizeBytes(td ptrace.Traces) int { _ = "STUB: not implemented"; return 0 }

func (bt *batchTraces) export(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (bt *batchTraces) split(sendBatchMaxSize int) (int, ptrace.Traces) {
	_ = "STUB: not implemented"
	return 0, *new(ptrace.Traces)
}

func (bt *batchTraces) itemCount() int { _ = "STUB: not implemented"; return 0 }

type batchMetrics struct {
	nextConsumer   consumer.Metrics
	metricData     pmetric.Metrics
	dataPointCount int
	sizer          pmetric.Sizer
}

func newMetricsBatch(nextConsumer consumer.Metrics) *batchMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (bm *batchMetrics) sizeBytes(md pmetric.Metrics) int { _ = "STUB: not implemented"; return 0 }

func (bm *batchMetrics) export(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (bm *batchMetrics) split(sendBatchMaxSize int) (int, pmetric.Metrics) {
	_ = "STUB: not implemented"
	return 0, *new(pmetric.Metrics)
}

func (bm *batchMetrics) itemCount() int { _ = "STUB: not implemented"; return 0 }

func (bm *batchMetrics) add(md pmetric.Metrics) { _ = "STUB: not implemented"; return }

type batchLogs struct {
	nextConsumer consumer.Logs
	logData      plog.Logs
	logCount     int
	sizer        plog.Sizer
}

func newBatchLogs(nextConsumer consumer.Logs) *batchLogs { _ = "STUB: not implemented"; return nil }

func (bl *batchLogs) sizeBytes(ld plog.Logs) int { _ = "STUB: not implemented"; return 0 }

func (bl *batchLogs) export(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (bl *batchLogs) split(sendBatchMaxSize int) (int, plog.Logs) {
	_ = "STUB: not implemented"
	return 0, *new(plog.Logs)
}

func (bl *batchLogs) itemCount() int { _ = "STUB: not implemented"; return 0 }

func (bl *batchLogs) add(ld plog.Logs) { _ = "STUB: not implemented"; return }

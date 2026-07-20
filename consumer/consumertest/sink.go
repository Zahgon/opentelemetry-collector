package consumertest

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type TracesSink struct {
	nonMutatingConsumer
	mu        sync.Mutex
	traces    []ptrace.Traces
	contexts  []context.Context
	spanCount int
}

var _ consumer.Traces = (*TracesSink)(nil)

func (ste *TracesSink) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (ste *TracesSink) AllTraces() []ptrace.Traces { _ = "STUB: not implemented"; return nil }

func (ste *TracesSink) Contexts() []context.Context { _ = "STUB: not implemented"; return nil }

func (ste *TracesSink) SpanCount() int { _ = "STUB: not implemented"; return 0 }

func (ste *TracesSink) Reset() { _ = "STUB: not implemented"; return }

type MetricsSink struct {
	nonMutatingConsumer
	mu             sync.Mutex
	metrics        []pmetric.Metrics
	contexts       []context.Context
	dataPointCount int
}

var _ consumer.Metrics = (*MetricsSink)(nil)

func (sme *MetricsSink) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (sme *MetricsSink) AllMetrics() []pmetric.Metrics { _ = "STUB: not implemented"; return nil }

func (sme *MetricsSink) Contexts() []context.Context { _ = "STUB: not implemented"; return nil }

func (sme *MetricsSink) DataPointCount() int { _ = "STUB: not implemented"; return 0 }

func (sme *MetricsSink) Reset() { _ = "STUB: not implemented"; return }

type LogsSink struct {
	nonMutatingConsumer
	mu             sync.Mutex
	logs           []plog.Logs
	contexts       []context.Context
	logRecordCount int
}

var _ consumer.Logs = (*LogsSink)(nil)

func (sle *LogsSink) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (sle *LogsSink) AllLogs() []plog.Logs { _ = "STUB: not implemented"; return nil }

func (sle *LogsSink) LogRecordCount() int { _ = "STUB: not implemented"; return 0 }

func (sle *LogsSink) Reset() { _ = "STUB: not implemented"; return }

func (sle *LogsSink) Contexts() []context.Context { _ = "STUB: not implemented"; return nil }

type ProfilesSink struct {
	nonMutatingConsumer
	mu           sync.Mutex
	profiles     []pprofile.Profiles
	contexts     []context.Context
	sampleCount  int
	profileCount int
}

var _ xconsumer.Profiles = (*ProfilesSink)(nil)

func (ste *ProfilesSink) ConsumeProfiles(ctx context.Context, td pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

func (ste *ProfilesSink) AllProfiles() []pprofile.Profiles { _ = "STUB: not implemented"; return nil }

func (ste *ProfilesSink) SampleCount() int { _ = "STUB: not implemented"; return 0 }

func (ste *ProfilesSink) ProfileCount() int { _ = "STUB: not implemented"; return 0 }

func (ste *ProfilesSink) Reset() { _ = "STUB: not implemented"; return }

func (ste *ProfilesSink) Contexts() []context.Context { _ = "STUB: not implemented"; return nil }

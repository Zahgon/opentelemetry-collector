package consumertest

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type Consumer interface {
	Capabilities() consumer.Capabilities

	ConsumeTraces(context.Context, ptrace.Traces) error

	ConsumeMetrics(context.Context, pmetric.Metrics) error

	ConsumeLogs(context.Context, plog.Logs) error

	ConsumeProfiles(context.Context, pprofile.Profiles) error

	unexported()
}

var (
	_ consumer.Logs      = Consumer(nil)
	_ consumer.Metrics   = Consumer(nil)
	_ consumer.Traces    = Consumer(nil)
	_ xconsumer.Profiles = Consumer(nil)
)

type nonMutatingConsumer struct{}

func (bc nonMutatingConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

type baseConsumer struct {
	nonMutatingConsumer
	consumer.ConsumeTracesFunc
	consumer.ConsumeMetricsFunc
	consumer.ConsumeLogsFunc
	xconsumer.ConsumeProfilesFunc
}

func (bc baseConsumer) unexported() { _ = "STUB: not implemented"; return }

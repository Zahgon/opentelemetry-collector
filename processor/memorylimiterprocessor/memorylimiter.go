package memorylimiterprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/internal/memorylimiter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
)

type memoryLimiterProcessor struct {
	memlimiter *memorylimiter.MemoryLimiter
	obsrep     *obsReport
}

func newMemoryLimiterProcessor(set processor.Settings, cfg *Config) (*memoryLimiterProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *memoryLimiterProcessor) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *memoryLimiterProcessor) shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *memoryLimiterProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (p *memoryLimiterProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (p *memoryLimiterProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (p *memoryLimiterProcessor) processProfiles(ctx context.Context, td pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

package otelcol

import (
	"sync/atomic"

	"go.uber.org/zap/zapcore"
)

var _ zapcore.Core = (*collectorCore)(nil)

type collectorCore struct {
	delegate atomic.Pointer[zapcore.Core]
}

func newCollectorCore(core zapcore.Core) *collectorCore { _ = "STUB: not implemented"; return nil }

func (c *collectorCore) Enabled(l zapcore.Level) bool { _ = "STUB: not implemented"; return false }

func (c *collectorCore) With(f []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (c *collectorCore) Check(e zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectorCore) Write(e zapcore.Entry, f []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collectorCore) Sync() error { _ = "STUB: not implemented"; return nil }

func (c *collectorCore) SetCore(core zapcore.Core) { _ = "STUB: not implemented"; return }

func (c *collectorCore) loadDelegate() zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

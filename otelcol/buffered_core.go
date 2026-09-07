package otelcol

import (
	"sync"

	"go.uber.org/zap/zapcore"
)

type loggedEntry struct {
	zapcore.Entry
	Context []zapcore.Field
}

func newBufferedCore(enab zapcore.LevelEnabler) *bufferedCore {
	_ = "STUB: not implemented"
	return nil
}

var _ zapcore.Core = (*bufferedCore)(nil)

type bufferedCore struct {
	zapcore.LevelEnabler
	mu        sync.Mutex
	logs      []loggedEntry
	context   []zapcore.Field
	logsTaken bool
}

func (bc *bufferedCore) Level() zapcore.Level {
	_ = "STUB: not implemented"
	return *new(zapcore.Level)
}

func (bc *bufferedCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bufferedCore) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (bc *bufferedCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bufferedCore) Sync() error { _ = "STUB: not implemented"; return nil }

func (bc *bufferedCore) TakeLogs() []loggedEntry { _ = "STUB: not implemented"; return nil }

package otelconftelemetry

import (
	"go.opentelemetry.io/otel/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapCoreProvider struct {
	sourceCore zapcore.Core
	lp         log.LoggerProvider
	scopeName  string
}

func (zcp *zapCoreProvider) newCore() zapCore { _ = "STUB: not implemented"; return *new(zapCore) }

type zapCore struct {
	sourceCore zapcore.Core
	otelCore   zapcore.Core
	provider   *zapCoreProvider
	withFields []zap.Field
}

var _ zapcore.Core = zapCore{}

func (zc zapCore) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (zc zapCore) Enabled(level zapcore.Level) bool { _ = "STUB: not implemented"; return false }

func (zc zapCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (zc zapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (zc zapCore) Sync() error { _ = "STUB: not implemented"; return nil }

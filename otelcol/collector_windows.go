//go:build windows

package otelcol

import (
	"flag"

	"go.uber.org/zap/zapcore"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
)

type windowsService struct {
	settings CollectorSettings
	col      *Collector
	flags    *flag.FlagSet
}

func NewSvcHandler(set CollectorSettings) svc.Handler {
	_ = "STUB: not implemented"
	return *new(svc.Handler)
}

func (s *windowsService) Execute(args []string, requests <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	_ = "STUB: not implemented"
	return false, 0
}

func (s *windowsService) start(elog *eventlog.Log, colErrorChannel chan error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *windowsService) stop(colErrorChannel chan error) error {
	_ = "STUB: not implemented"
	return nil
}

func openEventLog(serviceName string) (*eventlog.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ zapcore.Core = (*windowsEventLogCore)(nil)

type windowsEventLogCore struct {
	core    zapcore.Core
	elog    *eventlog.Log
	encoder zapcore.Encoder
}

func (w windowsEventLogCore) Enabled(level zapcore.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (w windowsEventLogCore) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (w windowsEventLogCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (w windowsEventLogCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

func (w windowsEventLogCore) Sync() error { _ = "STUB: not implemented"; return nil }

func withWindowsCore(elog *eventlog.Log) func(zapcore.Core) zapcore.Core {
	_ = "STUB: not implemented"
	return nil
}

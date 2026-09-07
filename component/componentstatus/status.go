package componentstatus

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Reporter interface {
	Report(*Event)
}

type Watcher interface {
	ComponentStatusChanged(source *InstanceID, event *Event)
}

type Status int32

const (
	StatusNone Status = iota

	StatusStarting

	StatusOK

	StatusRecoverableError

	StatusPermanentError

	StatusFatalError

	StatusStopping

	StatusStopped
)

func (s Status) String() string { _ = "STUB: not implemented"; return "" }

type Event struct {
	attributes pcommon.Map

	status    Status
	err       error
	timestamp time.Time
}

func (ev *Event) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (ev *Event) Err() error { _ = "STUB: not implemented"; return nil }

func (ev *Event) Attributes() pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

func (ev *Event) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type EventBuilderOption interface {
	applyOption(*Event)
}

type eventOptionFunc func(*Event)

func (f eventOptionFunc) applyOption(event *Event) { _ = "STUB: not implemented"; return }

func NewEvent(st Status, opts ...EventBuilderOption) *Event { _ = "STUB: not implemented"; return nil }

func WithAttributes(attributes pcommon.Map) EventBuilderOption {
	_ = "STUB: not implemented"
	return *new(EventBuilderOption)
}

func WithError(err error) EventBuilderOption {
	_ = "STUB: not implemented"
	return *new(EventBuilderOption)
}

func NewRecoverableErrorEvent(err error) *Event { _ = "STUB: not implemented"; return nil }

func NewPermanentErrorEvent(err error) *Event { _ = "STUB: not implemented"; return nil }

func NewFatalErrorEvent(err error) *Event { _ = "STUB: not implemented"; return nil }

func StatusIsError(status Status) bool { _ = "STUB: not implemented"; return false }

func ReportStatus(host component.Host, e *Event) { _ = "STUB: not implemented"; return }

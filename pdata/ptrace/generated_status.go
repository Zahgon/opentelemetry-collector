package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Status struct {
	orig  *internal.Status
	state *internal.State
}

func newStatus(orig *internal.Status, state *internal.State) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

func NewStatus() Status { _ = "STUB: not implemented"; return *new(Status) }

func (ms Status) MoveTo(dest Status) { _ = "STUB: not implemented"; return }

func (ms Status) Message() string { _ = "STUB: not implemented"; return "" }

func (ms Status) SetMessage(v string) { _ = "STUB: not implemented"; return }

func (ms Status) Code() StatusCode { _ = "STUB: not implemented"; return *new(StatusCode) }

func (ms Status) SetCode(v StatusCode) { _ = "STUB: not implemented"; return }

func (ms Status) CopyTo(dest Status) { _ = "STUB: not implemented"; return }

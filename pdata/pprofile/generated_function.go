package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Function struct {
	orig  *internal.Function
	state *internal.State
}

func newFunction(orig *internal.Function, state *internal.State) Function {
	_ = "STUB: not implemented"
	return *new(Function)
}

func NewFunction() Function { _ = "STUB: not implemented"; return *new(Function) }

func (ms Function) MoveTo(dest Function) { _ = "STUB: not implemented"; return }

func (ms Function) NameStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Function) SetNameStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms Function) SystemNameStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Function) SetSystemNameStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms Function) FilenameStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Function) SetFilenameStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms Function) StartLine() int64 { _ = "STUB: not implemented"; return 0 }

func (ms Function) SetStartLine(v int64) { _ = "STUB: not implemented"; return }

func (ms Function) CopyTo(dest Function) { _ = "STUB: not implemented"; return }

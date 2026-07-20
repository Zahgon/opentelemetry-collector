package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Line struct {
	orig  *internal.Line
	state *internal.State
}

func newLine(orig *internal.Line, state *internal.State) Line {
	_ = "STUB: not implemented"
	return *new(Line)
}

func NewLine() Line { _ = "STUB: not implemented"; return *new(Line) }

func (ms Line) MoveTo(dest Line) { _ = "STUB: not implemented"; return }

func (ms Line) FunctionIndex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Line) SetFunctionIndex(v int32) { _ = "STUB: not implemented"; return }

func (ms Line) Line() int64 { _ = "STUB: not implemented"; return 0 }

func (ms Line) SetLine(v int64) { _ = "STUB: not implemented"; return }

func (ms Line) Column() int64 { _ = "STUB: not implemented"; return 0 }

func (ms Line) SetColumn(v int64) { _ = "STUB: not implemented"; return }

func (ms Line) CopyTo(dest Line) { _ = "STUB: not implemented"; return }

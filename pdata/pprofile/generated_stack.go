package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Stack struct {
	orig  *internal.Stack
	state *internal.State
}

func newStack(orig *internal.Stack, state *internal.State) Stack {
	_ = "STUB: not implemented"
	return *new(Stack)
}

func NewStack() Stack { _ = "STUB: not implemented"; return *new(Stack) }

func (ms Stack) MoveTo(dest Stack) { _ = "STUB: not implemented"; return }

func (ms Stack) LocationIndices() pcommon.Int32Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Int32Slice)
}

func (ms Stack) CopyTo(dest Stack) { _ = "STUB: not implemented"; return }

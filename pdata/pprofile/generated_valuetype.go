package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type ValueType struct {
	orig  *internal.ValueType
	state *internal.State
}

func newValueType(orig *internal.ValueType, state *internal.State) ValueType {
	_ = "STUB: not implemented"
	return *new(ValueType)
}

func NewValueType() ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

func (ms ValueType) MoveTo(dest ValueType) { _ = "STUB: not implemented"; return }

func (ms ValueType) TypeStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms ValueType) SetTypeStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms ValueType) UnitStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms ValueType) SetUnitStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms ValueType) CopyTo(dest ValueType) { _ = "STUB: not implemented"; return }

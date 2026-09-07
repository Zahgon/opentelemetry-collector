package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type KeyValueAndUnit struct {
	orig  *internal.KeyValueAndUnit
	state *internal.State
}

func newKeyValueAndUnit(orig *internal.KeyValueAndUnit, state *internal.State) KeyValueAndUnit {
	_ = "STUB: not implemented"
	return *new(KeyValueAndUnit)
}

func NewKeyValueAndUnit() KeyValueAndUnit { _ = "STUB: not implemented"; return *new(KeyValueAndUnit) }

func (ms KeyValueAndUnit) MoveTo(dest KeyValueAndUnit) { _ = "STUB: not implemented"; return }

func (ms KeyValueAndUnit) KeyStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms KeyValueAndUnit) SetKeyStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms KeyValueAndUnit) Value() pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}

func (ms KeyValueAndUnit) UnitStrindex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms KeyValueAndUnit) SetUnitStrindex(v int32) { _ = "STUB: not implemented"; return }

func (ms KeyValueAndUnit) CopyTo(dest KeyValueAndUnit) { _ = "STUB: not implemented"; return }

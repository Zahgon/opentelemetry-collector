package entity

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type EntityRef internal.EntityRefWrapper

func newEntityRef(orig *internal.EntityRef, state *internal.State) EntityRef {
	_ = "STUB: not implemented"
	return *new(EntityRef)
}

func NewEntityRef() EntityRef { _ = "STUB: not implemented"; return *new(EntityRef) }

func (ms EntityRef) MoveTo(dest EntityRef) { _ = "STUB: not implemented"; return }

func (ms EntityRef) SchemaUrl() string { _ = "STUB: not implemented"; return "" }

func (ms EntityRef) SetSchemaUrl(v string) { _ = "STUB: not implemented"; return }

func (ms EntityRef) Type() string { _ = "STUB: not implemented"; return "" }

func (ms EntityRef) SetType(v string) { _ = "STUB: not implemented"; return }

func (ms EntityRef) IdKeys() pcommon.StringSlice {
	_ = "STUB: not implemented"
	return *new(pcommon.StringSlice)
}

func (ms EntityRef) DescriptionKeys() pcommon.StringSlice {
	_ = "STUB: not implemented"
	return *new(pcommon.StringSlice)
}

func (ms EntityRef) CopyTo(dest EntityRef) { _ = "STUB: not implemented"; return }

func (ms EntityRef) getOrig() *internal.EntityRef { _ = "STUB: not implemented"; return nil }

func (ms EntityRef) getState() *internal.State { _ = "STUB: not implemented"; return nil }

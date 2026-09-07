package xpdata

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type MapBuilder struct {
	state internal.State
	pairs []internal.KeyValue
}

func (mb *MapBuilder) EnsureCapacity(capacity int) { _ = "STUB: not implemented"; return }

func (mb *MapBuilder) getValue(i int) pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}

func (mb *MapBuilder) AppendEmpty(k string) pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}

func (mb *MapBuilder) UnsafeIntoMap(m pcommon.Map) { _ = "STUB: not implemented"; return }

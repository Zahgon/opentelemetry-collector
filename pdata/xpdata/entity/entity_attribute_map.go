package entity

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type EntityAttributeMap struct {
	keys       pcommon.StringSlice
	attributes pcommon.Map
}

func (m EntityAttributeMap) Get(key string) (pcommon.Value, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), false
}

func (m EntityAttributeMap) CanPut(key string) bool { _ = "STUB: not implemented"; return false }

func (m EntityAttributeMap) PutEmpty(k string) pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}

func (m EntityAttributeMap) PutStr(k, v string) { _ = "STUB: not implemented"; return }

func (m EntityAttributeMap) Remove(key string) bool { _ = "STUB: not implemented"; return false }

func (m EntityAttributeMap) containsKey(key string) bool { _ = "STUB: not implemented"; return false }

func (m EntityAttributeMap) All() iter.Seq2[string, pcommon.Value] {
	_ = "STUB: not implemented"
	return nil
}

package entity

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type EntityMap struct {
	refs       EntityRefSlice
	attributes pcommon.Map
}

func NewEntityMap() EntityMap { _ = "STUB: not implemented"; return *new(EntityMap) }

func (em EntityMap) Len() int { _ = "STUB: not implemented"; return 0 }

func (em EntityMap) EnsureCapacity(newCap int) { _ = "STUB: not implemented"; return }

func (em EntityMap) Get(entityType string) (Entity, bool) {
	_ = "STUB: not implemented"
	return *new(Entity), false
}

func (em EntityMap) All() iter.Seq2[string, Entity] { _ = "STUB: not implemented"; return nil }

func (em EntityMap) Remove(entityType string) bool { _ = "STUB: not implemented"; return false }

func (em EntityMap) PutEmpty(entityType string) Entity {
	_ = "STUB: not implemented"
	return *new(Entity)
}

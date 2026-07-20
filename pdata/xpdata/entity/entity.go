package entity

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Entity struct {
	ref        EntityRef
	attributes pcommon.Map
}

func NewEntity(t string) Entity { _ = "STUB: not implemented"; return *new(Entity) }

func (e Entity) Type() string { _ = "STUB: not implemented"; return "" }

func (e Entity) SchemaURL() string { _ = "STUB: not implemented"; return "" }

func (e Entity) SetSchemaURL(schemaURL string) { _ = "STUB: not implemented"; return }

func (e Entity) IdentifyingAttributes() EntityAttributeMap {
	_ = "STUB: not implemented"
	return *new(EntityAttributeMap)
}

func (e Entity) DescriptiveAttributes() EntityAttributeMap {
	_ = "STUB: not implemented"
	return *new(EntityAttributeMap)
}

func (e Entity) CopyToResource(res pcommon.Resource) { _ = "STUB: not implemented"; return }

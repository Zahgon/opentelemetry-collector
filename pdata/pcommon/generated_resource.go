package pcommon

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Resource internal.ResourceWrapper

func newResource(orig *internal.Resource, state *internal.State) Resource {
	_ = "STUB: not implemented"
	return *new(Resource)
}

func NewResource() Resource { _ = "STUB: not implemented"; return *new(Resource) }

func (ms Resource) MoveTo(dest Resource) { _ = "STUB: not implemented"; return }

func (ms Resource) Attributes() Map { _ = "STUB: not implemented"; return *new(Map) }

func (ms Resource) DroppedAttributesCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms Resource) SetDroppedAttributesCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms Resource) CopyTo(dest Resource) { _ = "STUB: not implemented"; return }

func (ms Resource) getOrig() *internal.Resource { _ = "STUB: not implemented"; return nil }

func (ms Resource) getState() *internal.State { _ = "STUB: not implemented"; return nil }

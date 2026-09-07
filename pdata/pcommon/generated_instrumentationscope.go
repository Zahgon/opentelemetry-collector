package pcommon

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type InstrumentationScope internal.InstrumentationScopeWrapper

func newInstrumentationScope(orig *internal.InstrumentationScope, state *internal.State) InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(InstrumentationScope)
}

func NewInstrumentationScope() InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(InstrumentationScope)
}

func (ms InstrumentationScope) MoveTo(dest InstrumentationScope) { _ = "STUB: not implemented"; return }

func (ms InstrumentationScope) Name() string { _ = "STUB: not implemented"; return "" }

func (ms InstrumentationScope) SetName(v string) { _ = "STUB: not implemented"; return }

func (ms InstrumentationScope) Version() string { _ = "STUB: not implemented"; return "" }

func (ms InstrumentationScope) SetVersion(v string) { _ = "STUB: not implemented"; return }

func (ms InstrumentationScope) Attributes() Map { _ = "STUB: not implemented"; return *new(Map) }

func (ms InstrumentationScope) DroppedAttributesCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms InstrumentationScope) SetDroppedAttributesCount(v uint32) {
	_ = "STUB: not implemented"
	return
}

func (ms InstrumentationScope) CopyTo(dest InstrumentationScope) { _ = "STUB: not implemented"; return }

func (ms InstrumentationScope) getOrig() *internal.InstrumentationScope {
	_ = "STUB: not implemented"
	return nil
}

func (ms InstrumentationScope) getState() *internal.State { _ = "STUB: not implemented"; return nil }

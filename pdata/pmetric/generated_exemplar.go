package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Exemplar struct {
	orig  *internal.Exemplar
	state *internal.State
}

func newExemplar(orig *internal.Exemplar, state *internal.State) Exemplar {
	_ = "STUB: not implemented"
	return *new(Exemplar)
}

func NewExemplar() Exemplar { _ = "STUB: not implemented"; return *new(Exemplar) }

func (ms Exemplar) MoveTo(dest Exemplar) { _ = "STUB: not implemented"; return }

func (ms Exemplar) FilteredAttributes() pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func (ms Exemplar) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms Exemplar) SetTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms Exemplar) ValueType() ExemplarValueType {
	_ = "STUB: not implemented"
	return *new(ExemplarValueType)
}

func (ms Exemplar) DoubleValue() float64 { _ = "STUB: not implemented"; return 0 }

func (ms Exemplar) SetDoubleValue(v float64) { _ = "STUB: not implemented"; return }

func (ms Exemplar) IntValue() int64 { _ = "STUB: not implemented"; return 0 }

func (ms Exemplar) SetIntValue(v int64) { _ = "STUB: not implemented"; return }

func (ms Exemplar) TraceID() pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

func (ms Exemplar) SetTraceID(v pcommon.TraceID) { _ = "STUB: not implemented"; return }

func (ms Exemplar) SpanID() pcommon.SpanID { _ = "STUB: not implemented"; return *new(pcommon.SpanID) }

func (ms Exemplar) SetSpanID(v pcommon.SpanID) { _ = "STUB: not implemented"; return }

func (ms Exemplar) CopyTo(dest Exemplar) { _ = "STUB: not implemented"; return }

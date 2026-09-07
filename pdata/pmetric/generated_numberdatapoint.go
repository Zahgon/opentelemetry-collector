package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type NumberDataPoint struct {
	orig  *internal.NumberDataPoint
	state *internal.State
}

func newNumberDataPoint(orig *internal.NumberDataPoint, state *internal.State) NumberDataPoint {
	_ = "STUB: not implemented"
	return *new(NumberDataPoint)
}

func NewNumberDataPoint() NumberDataPoint { _ = "STUB: not implemented"; return *new(NumberDataPoint) }

func (ms NumberDataPoint) MoveTo(dest NumberDataPoint) { _ = "STUB: not implemented"; return }

func (ms NumberDataPoint) Attributes() pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func (ms NumberDataPoint) StartTimestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms NumberDataPoint) SetStartTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms NumberDataPoint) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms NumberDataPoint) SetTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms NumberDataPoint) ValueType() NumberDataPointValueType {
	_ = "STUB: not implemented"
	return *new(NumberDataPointValueType)
}

func (ms NumberDataPoint) DoubleValue() float64 { _ = "STUB: not implemented"; return 0 }

func (ms NumberDataPoint) SetDoubleValue(v float64) { _ = "STUB: not implemented"; return }

func (ms NumberDataPoint) IntValue() int64 { _ = "STUB: not implemented"; return 0 }

func (ms NumberDataPoint) SetIntValue(v int64) { _ = "STUB: not implemented"; return }

func (ms NumberDataPoint) Exemplars() ExemplarSlice {
	_ = "STUB: not implemented"
	return *new(ExemplarSlice)
}

func (ms NumberDataPoint) Flags() DataPointFlags {
	_ = "STUB: not implemented"
	return *new(DataPointFlags)
}

func (ms NumberDataPoint) SetFlags(v DataPointFlags) { _ = "STUB: not implemented"; return }

func (ms NumberDataPoint) CopyTo(dest NumberDataPoint) { _ = "STUB: not implemented"; return }

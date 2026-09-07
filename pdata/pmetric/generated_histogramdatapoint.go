package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type HistogramDataPoint struct {
	orig  *internal.HistogramDataPoint
	state *internal.State
}

func newHistogramDataPoint(orig *internal.HistogramDataPoint, state *internal.State) HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(HistogramDataPoint)
}

func NewHistogramDataPoint() HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(HistogramDataPoint)
}

func (ms HistogramDataPoint) MoveTo(dest HistogramDataPoint) { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) Attributes() pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func (ms HistogramDataPoint) StartTimestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms HistogramDataPoint) SetStartTimestamp(v pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func (ms HistogramDataPoint) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms HistogramDataPoint) SetTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms HistogramDataPoint) SetCount(v uint64) { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) Sum() float64 { _ = "STUB: not implemented"; return 0 }

func (ms HistogramDataPoint) HasSum() bool { _ = "STUB: not implemented"; return false }

func (ms HistogramDataPoint) SetSum(v float64) { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) RemoveSum() { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) BucketCounts() pcommon.UInt64Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.UInt64Slice)
}

func (ms HistogramDataPoint) ExplicitBounds() pcommon.Float64Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Float64Slice)
}

func (ms HistogramDataPoint) Exemplars() ExemplarSlice {
	_ = "STUB: not implemented"
	return *new(ExemplarSlice)
}

func (ms HistogramDataPoint) Flags() DataPointFlags {
	_ = "STUB: not implemented"
	return *new(DataPointFlags)
}

func (ms HistogramDataPoint) SetFlags(v DataPointFlags) { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) Min() float64 { _ = "STUB: not implemented"; return 0 }

func (ms HistogramDataPoint) HasMin() bool { _ = "STUB: not implemented"; return false }

func (ms HistogramDataPoint) SetMin(v float64) { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) RemoveMin() { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) Max() float64 { _ = "STUB: not implemented"; return 0 }

func (ms HistogramDataPoint) HasMax() bool { _ = "STUB: not implemented"; return false }

func (ms HistogramDataPoint) SetMax(v float64) { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) RemoveMax() { _ = "STUB: not implemented"; return }

func (ms HistogramDataPoint) CopyTo(dest HistogramDataPoint) { _ = "STUB: not implemented"; return }

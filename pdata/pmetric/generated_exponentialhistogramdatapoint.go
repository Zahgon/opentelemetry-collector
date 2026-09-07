package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ExponentialHistogramDataPoint struct {
	orig  *internal.ExponentialHistogramDataPoint
	state *internal.State
}

func newExponentialHistogramDataPoint(orig *internal.ExponentialHistogramDataPoint, state *internal.State) ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPoint)
}

func NewExponentialHistogramDataPoint() ExponentialHistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPoint)
}

func (ms ExponentialHistogramDataPoint) MoveTo(dest ExponentialHistogramDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogramDataPoint) Attributes() pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func (ms ExponentialHistogramDataPoint) StartTimestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms ExponentialHistogramDataPoint) SetStartTimestamp(v pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogramDataPoint) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms ExponentialHistogramDataPoint) SetTimestamp(v pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogramDataPoint) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms ExponentialHistogramDataPoint) SetCount(v uint64) { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) Sum() float64 { _ = "STUB: not implemented"; return 0 }

func (ms ExponentialHistogramDataPoint) HasSum() bool { _ = "STUB: not implemented"; return false }

func (ms ExponentialHistogramDataPoint) SetSum(v float64) { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) RemoveSum() { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) Scale() int32 { _ = "STUB: not implemented"; return 0 }

func (ms ExponentialHistogramDataPoint) SetScale(v int32) { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) ZeroCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms ExponentialHistogramDataPoint) SetZeroCount(v uint64) { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) Positive() ExponentialHistogramDataPointBuckets {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPointBuckets)
}

func (ms ExponentialHistogramDataPoint) Negative() ExponentialHistogramDataPointBuckets {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPointBuckets)
}

func (ms ExponentialHistogramDataPoint) Flags() DataPointFlags {
	_ = "STUB: not implemented"
	return *new(DataPointFlags)
}

func (ms ExponentialHistogramDataPoint) SetFlags(v DataPointFlags) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogramDataPoint) Exemplars() ExemplarSlice {
	_ = "STUB: not implemented"
	return *new(ExemplarSlice)
}

func (ms ExponentialHistogramDataPoint) Min() float64 { _ = "STUB: not implemented"; return 0 }

func (ms ExponentialHistogramDataPoint) HasMin() bool { _ = "STUB: not implemented"; return false }

func (ms ExponentialHistogramDataPoint) SetMin(v float64) { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) RemoveMin() { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) Max() float64 { _ = "STUB: not implemented"; return 0 }

func (ms ExponentialHistogramDataPoint) HasMax() bool { _ = "STUB: not implemented"; return false }

func (ms ExponentialHistogramDataPoint) SetMax(v float64) { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) RemoveMax() { _ = "STUB: not implemented"; return }

func (ms ExponentialHistogramDataPoint) ZeroThreshold() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (ms ExponentialHistogramDataPoint) SetZeroThreshold(v float64) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogramDataPoint) CopyTo(dest ExponentialHistogramDataPoint) {
	_ = "STUB: not implemented"
	return
}

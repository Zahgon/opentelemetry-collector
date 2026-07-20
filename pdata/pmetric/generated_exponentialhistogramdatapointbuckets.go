package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type ExponentialHistogramDataPointBuckets struct {
	orig  *internal.ExponentialHistogramDataPointBuckets
	state *internal.State
}

func newExponentialHistogramDataPointBuckets(orig *internal.ExponentialHistogramDataPointBuckets, state *internal.State) ExponentialHistogramDataPointBuckets {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPointBuckets)
}

func NewExponentialHistogramDataPointBuckets() ExponentialHistogramDataPointBuckets {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogramDataPointBuckets)
}

func (ms ExponentialHistogramDataPointBuckets) MoveTo(dest ExponentialHistogramDataPointBuckets) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogramDataPointBuckets) Offset() int32 { _ = "STUB: not implemented"; return 0 }

func (ms ExponentialHistogramDataPointBuckets) SetOffset(v int32) {
	_ = "STUB: not implemented"
	return
}

func (ms ExponentialHistogramDataPointBuckets) BucketCounts() pcommon.UInt64Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.UInt64Slice)
}

func (ms ExponentialHistogramDataPointBuckets) CopyTo(dest ExponentialHistogramDataPointBuckets) {
	_ = "STUB: not implemented"
	return
}

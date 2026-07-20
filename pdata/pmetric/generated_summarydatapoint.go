package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type SummaryDataPoint struct {
	orig  *internal.SummaryDataPoint
	state *internal.State
}

func newSummaryDataPoint(orig *internal.SummaryDataPoint, state *internal.State) SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(SummaryDataPoint)
}

func NewSummaryDataPoint() SummaryDataPoint {
	_ = "STUB: not implemented"
	return *new(SummaryDataPoint)
}

func (ms SummaryDataPoint) MoveTo(dest SummaryDataPoint) { _ = "STUB: not implemented"; return }

func (ms SummaryDataPoint) Attributes() pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func (ms SummaryDataPoint) StartTimestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms SummaryDataPoint) SetStartTimestamp(v pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func (ms SummaryDataPoint) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms SummaryDataPoint) SetTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms SummaryDataPoint) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms SummaryDataPoint) SetCount(v uint64) { _ = "STUB: not implemented"; return }

func (ms SummaryDataPoint) Sum() float64 { _ = "STUB: not implemented"; return 0 }

func (ms SummaryDataPoint) SetSum(v float64) { _ = "STUB: not implemented"; return }

func (ms SummaryDataPoint) QuantileValues() SummaryDataPointValueAtQuantileSlice {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointValueAtQuantileSlice)
}

func (ms SummaryDataPoint) Flags() DataPointFlags {
	_ = "STUB: not implemented"
	return *new(DataPointFlags)
}

func (ms SummaryDataPoint) SetFlags(v DataPointFlags) { _ = "STUB: not implemented"; return }

func (ms SummaryDataPoint) CopyTo(dest SummaryDataPoint) { _ = "STUB: not implemented"; return }

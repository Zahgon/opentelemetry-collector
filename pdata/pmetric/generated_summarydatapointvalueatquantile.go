package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type SummaryDataPointValueAtQuantile struct {
	orig  *internal.SummaryDataPointValueAtQuantile
	state *internal.State
}

func newSummaryDataPointValueAtQuantile(orig *internal.SummaryDataPointValueAtQuantile, state *internal.State) SummaryDataPointValueAtQuantile {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointValueAtQuantile)
}

func NewSummaryDataPointValueAtQuantile() SummaryDataPointValueAtQuantile {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointValueAtQuantile)
}

func (ms SummaryDataPointValueAtQuantile) MoveTo(dest SummaryDataPointValueAtQuantile) {
	_ = "STUB: not implemented"
	return
}

func (ms SummaryDataPointValueAtQuantile) Quantile() float64 { _ = "STUB: not implemented"; return 0 }

func (ms SummaryDataPointValueAtQuantile) SetQuantile(v float64) { _ = "STUB: not implemented"; return }

func (ms SummaryDataPointValueAtQuantile) Value() float64 { _ = "STUB: not implemented"; return 0 }

func (ms SummaryDataPointValueAtQuantile) SetValue(v float64) { _ = "STUB: not implemented"; return }

func (ms SummaryDataPointValueAtQuantile) CopyTo(dest SummaryDataPointValueAtQuantile) {
	_ = "STUB: not implemented"
	return
}

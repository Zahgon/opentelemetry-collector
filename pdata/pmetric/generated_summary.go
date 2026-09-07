package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type Summary struct {
	orig  *internal.Summary
	state *internal.State
}

func newSummary(orig *internal.Summary, state *internal.State) Summary {
	_ = "STUB: not implemented"
	return *new(Summary)
}

func NewSummary() Summary { _ = "STUB: not implemented"; return *new(Summary) }

func (ms Summary) MoveTo(dest Summary) { _ = "STUB: not implemented"; return }

func (ms Summary) DataPoints() SummaryDataPointSlice {
	_ = "STUB: not implemented"
	return *new(SummaryDataPointSlice)
}

func (ms Summary) CopyTo(dest Summary) { _ = "STUB: not implemented"; return }

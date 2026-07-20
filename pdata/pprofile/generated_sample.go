package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Sample struct {
	orig  *internal.Sample
	state *internal.State
}

func newSample(orig *internal.Sample, state *internal.State) Sample {
	_ = "STUB: not implemented"
	return *new(Sample)
}

func NewSample() Sample { _ = "STUB: not implemented"; return *new(Sample) }

func (ms Sample) MoveTo(dest Sample) { _ = "STUB: not implemented"; return }

func (ms Sample) StackIndex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Sample) SetStackIndex(v int32) { _ = "STUB: not implemented"; return }

func (ms Sample) AttributeIndices() pcommon.Int32Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Int32Slice)
}

func (ms Sample) LinkIndex() int32 { _ = "STUB: not implemented"; return 0 }

func (ms Sample) SetLinkIndex(v int32) { _ = "STUB: not implemented"; return }

func (ms Sample) Values() pcommon.Int64Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Int64Slice)
}

func (ms Sample) TimestampsUnixNano() pcommon.UInt64Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.UInt64Slice)
}

func (ms Sample) CopyTo(dest Sample) { _ = "STUB: not implemented"; return }

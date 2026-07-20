package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Profile struct {
	orig  *internal.Profile
	state *internal.State
}

func newProfile(orig *internal.Profile, state *internal.State) Profile {
	_ = "STUB: not implemented"
	return *new(Profile)
}

func NewProfile() Profile { _ = "STUB: not implemented"; return *new(Profile) }

func (ms Profile) MoveTo(dest Profile) { _ = "STUB: not implemented"; return }

func (ms Profile) SampleType() ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

func (ms Profile) Samples() SampleSlice { _ = "STUB: not implemented"; return *new(SampleSlice) }

func (ms Profile) Time() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms Profile) SetTime(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms Profile) DurationNano() uint64 { _ = "STUB: not implemented"; return 0 }

func (ms Profile) SetDurationNano(v uint64) { _ = "STUB: not implemented"; return }

func (ms Profile) PeriodType() ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

func (ms Profile) Period() int64 { _ = "STUB: not implemented"; return 0 }

func (ms Profile) SetPeriod(v int64) { _ = "STUB: not implemented"; return }

func (ms Profile) ProfileID() ProfileID { _ = "STUB: not implemented"; return *new(ProfileID) }

func (ms Profile) SetProfileID(v ProfileID) { _ = "STUB: not implemented"; return }

func (ms Profile) DroppedAttributesCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms Profile) SetDroppedAttributesCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms Profile) OriginalPayloadFormat() string { _ = "STUB: not implemented"; return "" }

func (ms Profile) SetOriginalPayloadFormat(v string) { _ = "STUB: not implemented"; return }

func (ms Profile) OriginalPayload() pcommon.ByteSlice {
	_ = "STUB: not implemented"
	return *new(pcommon.ByteSlice)
}

func (ms Profile) AttributeIndices() pcommon.Int32Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Int32Slice)
}

func (ms Profile) CopyTo(dest Profile) { _ = "STUB: not implemented"; return }

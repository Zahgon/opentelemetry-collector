package pmetric

const noRecordValueMask = uint32(1)

var DefaultDataPointFlags = DataPointFlags(0)

type DataPointFlags uint32

func (ms DataPointFlags) NoRecordedValue() bool { _ = "STUB: not implemented"; return false }

func (ms DataPointFlags) WithNoRecordedValue(b bool) DataPointFlags {
	_ = "STUB: not implemented"
	return *new(DataPointFlags)
}

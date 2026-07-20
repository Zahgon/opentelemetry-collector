package plog

const isSampledMask = uint32(1)

var DefaultLogRecordFlags = LogRecordFlags(0)

type LogRecordFlags uint32

func (ms LogRecordFlags) IsSampled() bool { _ = "STUB: not implemented"; return false }

func (ms LogRecordFlags) WithIsSampled(b bool) LogRecordFlags {
	_ = "STUB: not implemented"
	return *new(LogRecordFlags)
}

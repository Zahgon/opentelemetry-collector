package plog

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type LogRecord struct {
	orig  *internal.LogRecord
	state *internal.State
}

func newLogRecord(orig *internal.LogRecord, state *internal.State) LogRecord {
	_ = "STUB: not implemented"
	return *new(LogRecord)
}

func NewLogRecord() LogRecord { _ = "STUB: not implemented"; return *new(LogRecord) }

func (ms LogRecord) MoveTo(dest LogRecord) { _ = "STUB: not implemented"; return }

func (ms LogRecord) Timestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms LogRecord) SetTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms LogRecord) ObservedTimestamp() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms LogRecord) SetObservedTimestamp(v pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (ms LogRecord) SeverityNumber() SeverityNumber {
	_ = "STUB: not implemented"
	return *new(SeverityNumber)
}

func (ms LogRecord) SetSeverityNumber(v SeverityNumber) { _ = "STUB: not implemented"; return }

func (ms LogRecord) SeverityText() string { _ = "STUB: not implemented"; return "" }

func (ms LogRecord) SetSeverityText(v string) { _ = "STUB: not implemented"; return }

func (ms LogRecord) Body() pcommon.Value { _ = "STUB: not implemented"; return *new(pcommon.Value) }

func (ms LogRecord) Attributes() pcommon.Map { _ = "STUB: not implemented"; return *new(pcommon.Map) }

func (ms LogRecord) DroppedAttributesCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (ms LogRecord) SetDroppedAttributesCount(v uint32) { _ = "STUB: not implemented"; return }

func (ms LogRecord) Flags() LogRecordFlags { _ = "STUB: not implemented"; return *new(LogRecordFlags) }

func (ms LogRecord) SetFlags(v LogRecordFlags) { _ = "STUB: not implemented"; return }

func (ms LogRecord) TraceID() pcommon.TraceID {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID)
}

func (ms LogRecord) SetTraceID(v pcommon.TraceID) { _ = "STUB: not implemented"; return }

func (ms LogRecord) SpanID() pcommon.SpanID { _ = "STUB: not implemented"; return *new(pcommon.SpanID) }

func (ms LogRecord) SetSpanID(v pcommon.SpanID) { _ = "STUB: not implemented"; return }

func (ms LogRecord) EventName() string { _ = "STUB: not implemented"; return "" }

func (ms LogRecord) SetEventName(v string) { _ = "STUB: not implemented"; return }

func (ms LogRecord) CopyTo(dest LogRecord) { _ = "STUB: not implemented"; return }

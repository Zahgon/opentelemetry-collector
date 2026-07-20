package plog

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalLogs(ld Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ProtoMarshaler) LogsSize(ld Logs) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) ResourceLogsSize(ld ResourceLogs) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) ScopeLogsSize(ld ScopeLogs) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) LogRecordSize(ld LogRecord) int { _ = "STUB: not implemented"; return 0 }

var _ Unmarshaler = (*ProtoUnmarshaler)(nil)

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalLogs(buf []byte) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

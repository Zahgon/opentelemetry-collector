package plog

type MarshalSizer interface {
	Marshaler
	Sizer
}

type Marshaler interface {
	MarshalLogs(ld Logs) ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalLogs(buf []byte) (Logs, error)
}

type Sizer interface {
	LogsSize(ld Logs) int
}

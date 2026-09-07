package ptrace

type MarshalSizer interface {
	Marshaler
	Sizer
}

type Marshaler interface {
	MarshalTraces(td Traces) ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalTraces(buf []byte) (Traces, error)
}

type Sizer interface {
	TracesSize(td Traces) int
}

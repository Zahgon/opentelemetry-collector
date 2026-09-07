package ptrace

type JSONMarshaler struct{}

func (*JSONMarshaler) MarshalTraces(td Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type JSONUnmarshaler struct {
	_ struct{}

	DisallowUnknownFields bool
}

func (u *JSONUnmarshaler) UnmarshalTraces(buf []byte) (Traces, error) {
	_ = "STUB: not implemented"
	return *new(Traces), nil
}

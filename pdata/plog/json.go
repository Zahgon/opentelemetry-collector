package plog

type JSONMarshaler struct{}

func (*JSONMarshaler) MarshalLogs(ld Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ Unmarshaler = (*JSONUnmarshaler)(nil)

type JSONUnmarshaler struct {
	_ struct{}

	DisallowUnknownFields bool
}

func (u *JSONUnmarshaler) UnmarshalLogs(buf []byte) (Logs, error) {
	_ = "STUB: not implemented"
	return *new(Logs), nil
}

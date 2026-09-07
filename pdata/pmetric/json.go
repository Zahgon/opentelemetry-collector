package pmetric

var _ Marshaler = (*JSONMarshaler)(nil)

type JSONMarshaler struct{}

func (*JSONMarshaler) MarshalMetrics(md Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type JSONUnmarshaler struct {
	_ struct{}

	DisallowUnknownFields bool
}

func (u *JSONUnmarshaler) UnmarshalMetrics(buf []byte) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

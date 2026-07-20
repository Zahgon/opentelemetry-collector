package pprofile

type JSONMarshaler struct{}

func (*JSONMarshaler) MarshalProfiles(pd Profiles) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type JSONUnmarshaler struct {
	DisallowUnknownFields bool

	_ struct{}
}

func (u *JSONUnmarshaler) UnmarshalProfiles(buf []byte) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

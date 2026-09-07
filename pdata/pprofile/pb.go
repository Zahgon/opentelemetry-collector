package pprofile

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalProfiles(pd Profiles) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ProtoMarshaler) ProfilesSize(pd Profiles) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) ResourceProfilesSize(pd ResourceProfiles) int {
	_ = "STUB: not implemented"
	return 0
}

func (e *ProtoMarshaler) ScopeProfilesSize(pd ScopeProfiles) int {
	_ = "STUB: not implemented"
	return 0
}

func (e *ProtoMarshaler) ProfileSize(pd Profile) int { _ = "STUB: not implemented"; return 0 }

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalProfiles(buf []byte) (Profiles, error) {
	_ = "STUB: not implemented"
	return *new(Profiles), nil
}

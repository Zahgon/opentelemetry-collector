package pprofile

type MarshalSizer interface {
	Marshaler
	Sizer
}

type Marshaler interface {
	MarshalProfiles(td Profiles) ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalProfiles(buf []byte) (Profiles, error)
}

type Sizer interface {
	ProfilesSize(td Profiles) int
}

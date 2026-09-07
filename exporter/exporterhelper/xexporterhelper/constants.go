package xexporterhelper

import (
	"errors"
)

var (
	errNilConfig = errors.New("nil config")

	errNilLogger = errors.New("nil logger")

	errNilConsumeRequest = errors.New("nil RequestConsumeFunc")

	errNilPushProfileData = errors.New("nil PushProfiles")

	errNilProfilesConverter = errors.New("nil RequestFromProfilesFunc")
)

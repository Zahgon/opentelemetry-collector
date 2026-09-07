package xconsumererror

import (
	"go.opentelemetry.io/collector/consumer/consumererror/internal"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

type Profiles struct {
	internal.Retryable[pprofile.Profiles]
}

func NewProfiles(err error, data pprofile.Profiles) error { _ = "STUB: not implemented"; return nil }

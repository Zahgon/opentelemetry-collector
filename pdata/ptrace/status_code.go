package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type StatusCode int32

const (
	StatusCodeUnset = StatusCode(internal.StatusCode_STATUS_CODE_UNSET)
	StatusCodeOk    = StatusCode(internal.StatusCode_STATUS_CODE_OK)
	StatusCodeError = StatusCode(internal.StatusCode_STATUS_CODE_ERROR)
)

func (sc StatusCode) String() string { _ = "STUB: not implemented"; return "" }

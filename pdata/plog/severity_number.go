package plog

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type SeverityNumber int32

const (
	SeverityNumberUnspecified = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED)
	SeverityNumberTrace       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_TRACE)
	SeverityNumberTrace2      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_TRACE2)
	SeverityNumberTrace3      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_TRACE3)
	SeverityNumberTrace4      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_TRACE4)
	SeverityNumberDebug       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_DEBUG)
	SeverityNumberDebug2      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_DEBUG2)
	SeverityNumberDebug3      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_DEBUG3)
	SeverityNumberDebug4      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_DEBUG4)
	SeverityNumberInfo        = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_INFO)
	SeverityNumberInfo2       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_INFO2)
	SeverityNumberInfo3       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_INFO3)
	SeverityNumberInfo4       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_INFO4)
	SeverityNumberWarn        = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_WARN)
	SeverityNumberWarn2       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_WARN2)
	SeverityNumberWarn3       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_WARN3)
	SeverityNumberWarn4       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_WARN4)
	SeverityNumberError       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_ERROR)
	SeverityNumberError2      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_ERROR2)
	SeverityNumberError3      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_ERROR3)
	SeverityNumberError4      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_ERROR4)
	SeverityNumberFatal       = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_FATAL)
	SeverityNumberFatal2      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_FATAL2)
	SeverityNumberFatal3      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_FATAL3)
	SeverityNumberFatal4      = SeverityNumber(internal.SeverityNumber_SEVERITY_NUMBER_FATAL4)
)

func (sn SeverityNumber) String() string { _ = "STUB: not implemented"; return "" }

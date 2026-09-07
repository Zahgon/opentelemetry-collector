package pprofile

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func (ms Profile) switchDictionary(src, dst ProfilesDictionary) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms Profile) Duration() pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func (ms Profile) SetDuration(_ pcommon.Timestamp) { _ = "STUB: not implemented"; return }

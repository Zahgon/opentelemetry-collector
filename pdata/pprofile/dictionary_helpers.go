package pprofile

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func mapKeyValues(m pcommon.Map) []internal.KeyValue { _ = "STUB: not implemented"; return nil }

func resolveProfilesReferences(profiles Profiles) { _ = "STUB: not implemented"; return }

func resolveKeyValueReferences(dict ProfilesDictionary, kvs []internal.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func resolveAnyValueReference(dict ProfilesDictionary, anyValue *internal.AnyValue) {
	_ = "STUB: not implemented"
	return
}

func convertProfilesToReferences(profiles Profiles) { _ = "STUB: not implemented"; return }

func convertKeyValueToReferences(getStringIndex func(string) int32, kvs []internal.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func convertAnyValueToReference(getStringIndex func(string) int32, anyValue *internal.AnyValue) {
	_ = "STUB: not implemented"
	return
}

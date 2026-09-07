package xconfmap

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/internal"
)

type ExpandedValue = internal.ExpandedValue

func ToStringMapRaw(conf *confmap.Conf) map[string]any { _ = "STUB: not implemented"; return nil }

func WithUnredacted() confmap.MarshalOption {
	_ = "STUB: not implemented"
	return *new(confmap.MarshalOption)
}

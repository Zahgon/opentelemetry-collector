package xconfmap

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/internal"
)

type Validator = internal.Validator

func Validate(cfg any) error { _ = "STUB: not implemented"; return nil }

func WithForceUnmarshaler() confmap.UnmarshalOption {
	_ = "STUB: not implemented"
	return *new(confmap.UnmarshalOption)
}

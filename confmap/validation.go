package confmap

import (
	"go.opentelemetry.io/collector/confmap/internal"
)

type Validator = internal.Validator

func Validate(cfg any) error { _ = "STUB: not implemented"; return nil }

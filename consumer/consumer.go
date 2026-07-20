package consumer

import (
	"errors"

	"go.opentelemetry.io/collector/consumer/internal"
)

type Capabilities = internal.Capabilities

var errNilFunc = errors.New("nil consumer func")

type Option = internal.Option

func WithCapabilities(capabilities Capabilities) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

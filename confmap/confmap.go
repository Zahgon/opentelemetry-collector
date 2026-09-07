//go:generate mdatagen metadata.yaml

package confmap

import (
	"go.opentelemetry.io/collector/confmap/internal"
)

var KeyDelimiter = internal.KeyDelimiter

var MapstructureTag = internal.MapstructureTag

func New() *Conf { _ = "STUB: not implemented"; return nil }

func NewFromStringMap(data map[string]any) *Conf { _ = "STUB: not implemented"; return nil }

type Conf = internal.Conf

type UnmarshalOption = internal.UnmarshalOption

func WithIgnoreUnused() UnmarshalOption { _ = "STUB: not implemented"; return *new(UnmarshalOption) }

func WithForceUnmarshaler() UnmarshalOption {
	_ = "STUB: not implemented"
	return *new(UnmarshalOption)
}

type MarshalOption = internal.MarshalOption

type Unmarshaler = internal.Unmarshaler

type Marshaler = internal.Marshaler

type ScalarValue = internal.ScalarValue

type ScalarUnmarshaler = internal.ScalarUnmarshaler

type ScalarMarshaler = internal.ScalarMarshaler

var ErrValueNotApplicable = internal.ErrValueNotApplicable

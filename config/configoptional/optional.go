package configoptional

//go:generate mdatagen metadata.yaml

import (
	"reflect"

	"go.opentelemetry.io/collector/confmap"
)

type flavor int

const (
	noneFlavor    flavor = 0
	defaultFlavor flavor = 1
	someFlavor    flavor = 2
)

type Optional[T any] struct {
	value T

	flavor flavor
}

func deref(t reflect.Type) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func assertNoEnabledField[T any]() error { _ = "STUB: not implemented"; return nil }

func Some[T any](value T) Optional[T] { _ = "STUB: not implemented"; return nil }

func Default[T any](value T) Optional[T] { _ = "STUB: not implemented"; return nil }

func None[T any]() Optional[T] { _ = "STUB: not implemented"; return nil }

func (o Optional[T]) HasValue() bool { _ = "STUB: not implemented"; return false }

func (o *Optional[T]) Get() *T { _ = "STUB: not implemented"; return nil }

func (o *Optional[T]) GetOrInsertDefault() *T { _ = "STUB: not implemented"; return nil }

var (
	_ confmap.Unmarshaler       = (*Optional[any])(nil)
	_ confmap.ScalarUnmarshaler = (*Optional[any])(nil)
)

func (o *Optional[T]) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

func (o *Optional[T]) UnmarshalScalar(scalarValue confmap.ScalarValue) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	_ confmap.Marshaler       = (*Optional[any])(nil)
	_ confmap.ScalarMarshaler = (*Optional[any])(nil)
)

func (o Optional[T]) Marshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

func (o Optional[T]) MarshalScalar(scalarValue confmap.ScalarValue) error {
	_ = "STUB: not implemented"
	return nil
}

var _ confmap.Validator = (*Optional[any])(nil)

func (o *Optional[T]) Validate() error { _ = "STUB: not implemented"; return nil }

package confmap

import (
	"reflect"
)

var configValidatorType = reflect.TypeFor[Validator]()

type Validator interface {
	Validate() error
}

func Validate(cfg any) error { _ = "STUB: not implemented"; return nil }

type pathError struct {
	err  error
	path []string
}

func (pe pathError) Error() string { _ = "STUB: not implemented"; return "" }

func (pe pathError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func validate(v reflect.Value) []pathError { _ = "STUB: not implemented"; return nil }

func callValidateIfPossible(v reflect.Value) error { _ = "STUB: not implemented"; return nil }

func fieldName(field reflect.StructField) string { _ = "STUB: not implemented"; return "" }

func stringifyMapKey(val reflect.Value) string { _ = "STUB: not implemented"; return "" }

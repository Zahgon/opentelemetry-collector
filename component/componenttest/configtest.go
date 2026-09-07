package componenttest

import (
	"reflect"
	"regexp"
)

var configFieldTagRegExp = regexp.MustCompile("^[a-z0-9][a-z0-9_]*(/[a-z0-9][a-z0-9_]*)*$")

func CheckConfigStruct(config any) error { _ = "STUB: not implemented"; return nil }

func validateConfigDataType(t reflect.Type) error { _ = "STUB: not implemented"; return nil }

func checkStructFieldTags(f reflect.StructField) error { _ = "STUB: not implemented"; return nil }

package pprofile

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type attributable interface {
	AttributeIndices() pcommon.Int32Slice
}

func FromAttributeIndices(table KeyValueAndUnitSlice, record attributable, dic ProfilesDictionary) (pcommon.Map, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Map), nil
}

var errTooManyAttributeTableEntries = errors.New("too many entries in AttributeTable")

func SetAttribute(table KeyValueAndUnitSlice, attr KeyValueAndUnit) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

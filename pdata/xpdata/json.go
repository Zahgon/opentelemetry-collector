package xpdata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type JSONMarshaler struct{}

func (*JSONMarshaler) MarshalValue(value pcommon.Value) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type JSONUnmarshaler struct {
	DisallowUnknownFields bool

	_ struct{}
}

func (u *JSONUnmarshaler) UnmarshalValue(buf []byte) (pcommon.Value, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), nil
}

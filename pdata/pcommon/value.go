package pcommon

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type ValueType int32

const (
	ValueTypeEmpty ValueType = iota
	ValueTypeStr
	ValueTypeInt
	ValueTypeDouble
	ValueTypeBool
	ValueTypeMap
	ValueTypeSlice
	ValueTypeBytes
)

func (avt ValueType) String() string { _ = "STUB: not implemented"; return "" }

type Value internal.ValueWrapper

func NewValueEmpty() Value { _ = "STUB: not implemented"; return *new(Value) }

func NewValueStr(v string) Value { _ = "STUB: not implemented"; return *new(Value) }

func NewValueInt(v int64) Value { _ = "STUB: not implemented"; return *new(Value) }

func NewValueDouble(v float64) Value { _ = "STUB: not implemented"; return *new(Value) }

func NewValueBool(v bool) Value { _ = "STUB: not implemented"; return *new(Value) }

func NewValueMap() Value { _ = "STUB: not implemented"; return *new(Value) }

func NewValueSlice() Value { _ = "STUB: not implemented"; return *new(Value) }

func NewValueBytes() Value { _ = "STUB: not implemented"; return *new(Value) }

func newValue(orig *internal.AnyValue, state *internal.State) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (v Value) getOrig() *internal.AnyValue { _ = "STUB: not implemented"; return nil }

func (v Value) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func (v Value) FromRaw(iv any) error { _ = "STUB: not implemented"; return nil }

func (v Value) Type() ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

func (v Value) Str() string { _ = "STUB: not implemented"; return "" }

func (v Value) Int() int64 { _ = "STUB: not implemented"; return 0 }

func (v Value) Double() float64 { _ = "STUB: not implemented"; return 0 }

func (v Value) Bool() bool { _ = "STUB: not implemented"; return false }

func (v Value) Map() Map { _ = "STUB: not implemented"; return *new(Map) }

func (v Value) Slice() Slice { _ = "STUB: not implemented"; return *new(Slice) }

func (v Value) Bytes() ByteSlice { _ = "STUB: not implemented"; return *new(ByteSlice) }

func (v Value) SetStr(sv string) { _ = "STUB: not implemented"; return }

func (v Value) SetInt(iv int64) { _ = "STUB: not implemented"; return }

func (v Value) SetDouble(dv float64) { _ = "STUB: not implemented"; return }

func (v Value) SetBool(bv bool) { _ = "STUB: not implemented"; return }

func (v Value) SetEmptyBytes() ByteSlice { _ = "STUB: not implemented"; return *new(ByteSlice) }

func (v Value) SetEmptyMap() Map { _ = "STUB: not implemented"; return *new(Map) }

func (v Value) SetEmptySlice() Slice { _ = "STUB: not implemented"; return *new(Slice) }

func (v Value) MoveTo(dest Value) { _ = "STUB: not implemented"; return }

func (v Value) CopyTo(dest Value) { _ = "STUB: not implemented"; return }

func (v Value) AsString() string { _ = "STUB: not implemented"; return "" }

func marshalJSONNoHTMLEscape(v any) string { _ = "STUB: not implemented"; return "" }

func float64AsString(f float64) string { _ = "STUB: not implemented"; return "" }

func (v Value) AsRaw() any { _ = "STUB: not implemented"; return *new(any) }

func (v Value) Equal(c Value) bool { _ = "STUB: not implemented"; return false }

package pcommon

import (
	"iter"

	"go.opentelemetry.io/collector/pdata/internal"
)

type Map internal.MapWrapper

func NewMap() Map { _ = "STUB: not implemented"; return *new(Map) }

func (m Map) getOrig() *[]internal.KeyValue { _ = "STUB: not implemented"; return nil }

func (m Map) getState() *internal.State { _ = "STUB: not implemented"; return nil }

func newMap(orig *[]internal.KeyValue, state *internal.State) Map {
	_ = "STUB: not implemented"
	return *new(Map)
}

func (m Map) Clear() { _ = "STUB: not implemented"; return }

func (m Map) EnsureCapacity(capacity int) { _ = "STUB: not implemented"; return }

func (m Map) Get(key string) (Value, bool) { _ = "STUB: not implemented"; return *new(Value), false }

func (m Map) Remove(key string) bool { _ = "STUB: not implemented"; return false }

func (m Map) RemoveIf(f func(string, Value) bool) { _ = "STUB: not implemented"; return }

func (m Map) PutEmpty(k string) Value { _ = "STUB: not implemented"; return *new(Value) }

func (m Map) GetOrPutEmpty(k string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func (m Map) PutStr(k, v string) { _ = "STUB: not implemented"; return }

func (m Map) PutInt(k string, v int64) { _ = "STUB: not implemented"; return }

func (m Map) PutDouble(k string, v float64) { _ = "STUB: not implemented"; return }

func (m Map) PutBool(k string, v bool) { _ = "STUB: not implemented"; return }

func (m Map) PutEmptyBytes(k string) ByteSlice { _ = "STUB: not implemented"; return *new(ByteSlice) }

func (m Map) PutEmptyMap(k string) Map { _ = "STUB: not implemented"; return *new(Map) }

func (m Map) PutEmptySlice(k string) Slice { _ = "STUB: not implemented"; return *new(Slice) }

func (m Map) Len() int { _ = "STUB: not implemented"; return 0 }

func (m Map) Range(f func(k string, v Value) bool) { _ = "STUB: not implemented"; return }

func (m Map) All() iter.Seq2[string, Value] { _ = "STUB: not implemented"; return nil }

func (m Map) MoveTo(dest Map) { _ = "STUB: not implemented"; return }

func (m Map) CopyTo(dest Map) { _ = "STUB: not implemented"; return }

func (m Map) AsRaw() map[string]any { _ = "STUB: not implemented"; return nil }

func (m Map) FromRaw(rawMap map[string]any) error { _ = "STUB: not implemented"; return nil }

func (m Map) Equal(val Map) bool { _ = "STUB: not implemented"; return false }

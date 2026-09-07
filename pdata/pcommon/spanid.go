package pcommon

var emptySpanID = SpanID([8]byte{})

type SpanID [8]byte

func NewSpanIDEmpty() SpanID { _ = "STUB: not implemented"; return *new(SpanID) }

func (ms SpanID) String() string { _ = "STUB: not implemented"; return "" }

func (ms SpanID) IsEmpty() bool { _ = "STUB: not implemented"; return false }

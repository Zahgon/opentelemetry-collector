package pcommon

var emptyTraceID = TraceID([16]byte{})

type TraceID [16]byte

func NewTraceIDEmpty() TraceID { _ = "STUB: not implemented"; return *new(TraceID) }

func (ms TraceID) String() string { _ = "STUB: not implemented"; return "" }

func (ms TraceID) IsEmpty() bool { _ = "STUB: not implemented"; return false }

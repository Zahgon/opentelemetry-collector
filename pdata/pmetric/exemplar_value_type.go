package pmetric

type ExemplarValueType int32

const (
	ExemplarValueTypeEmpty ExemplarValueType = iota
	ExemplarValueTypeInt
	ExemplarValueTypeDouble
)

func (nt ExemplarValueType) String() string { _ = "STUB: not implemented"; return "" }

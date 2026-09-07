package pmetric

type NumberDataPointValueType int32

const (
	NumberDataPointValueTypeEmpty NumberDataPointValueType = iota
	NumberDataPointValueTypeInt
	NumberDataPointValueTypeDouble
)

func (nt NumberDataPointValueType) String() string { _ = "STUB: not implemented"; return "" }

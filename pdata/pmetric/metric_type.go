package pmetric

type MetricType int32

const (
	MetricTypeEmpty MetricType = iota
	MetricTypeGauge
	MetricTypeSum
	MetricTypeHistogram
	MetricTypeExponentialHistogram
	MetricTypeSummary
)

func (mdt MetricType) String() string { _ = "STUB: not implemented"; return "" }

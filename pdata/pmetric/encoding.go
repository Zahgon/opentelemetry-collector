package pmetric

type MarshalSizer interface {
	Marshaler
	Sizer
}

type Marshaler interface {
	MarshalMetrics(md Metrics) ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalMetrics(buf []byte) (Metrics, error)
}

type Sizer interface {
	MetricsSize(md Metrics) int
}

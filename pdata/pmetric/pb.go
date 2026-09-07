package pmetric

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalMetrics(md Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ProtoMarshaler) MetricsSize(md Metrics) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) ResourceMetricsSize(md ResourceMetrics) int {
	_ = "STUB: not implemented"
	return 0
}

func (e *ProtoMarshaler) ScopeMetricsSize(md ScopeMetrics) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) MetricSize(md Metric) int { _ = "STUB: not implemented"; return 0 }

func (e *ProtoMarshaler) NumberDataPointSize(md NumberDataPoint) int {
	_ = "STUB: not implemented"
	return 0
}

func (e *ProtoMarshaler) SummaryDataPointSize(md SummaryDataPoint) int {
	_ = "STUB: not implemented"
	return 0
}

func (e *ProtoMarshaler) HistogramDataPointSize(md HistogramDataPoint) int {
	_ = "STUB: not implemented"
	return 0
}

func (e *ProtoMarshaler) ExponentialHistogramDataPointSize(md ExponentialHistogramDataPoint) int {
	_ = "STUB: not implemented"
	return 0
}

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalMetrics(buf []byte) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

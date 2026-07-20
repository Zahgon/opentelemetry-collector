package batchprocessor

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func splitMetrics(size int, src pmetric.Metrics) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func resourceMetricsDPC(rs pmetric.ResourceMetrics) int { _ = "STUB: not implemented"; return 0 }

func scopeMetricsDPC(ilm pmetric.ScopeMetrics) int { _ = "STUB: not implemented"; return 0 }

func metricDPC(ms pmetric.Metric) int { _ = "STUB: not implemented"; return 0 }

func splitMetric(ms, dest pmetric.Metric, size int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func splitNumberDataPoints(src, dst pmetric.NumberDataPointSlice, size int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func splitHistogramDataPoints(src, dst pmetric.HistogramDataPointSlice, size int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func splitExponentialHistogramDataPoints(src, dst pmetric.ExponentialHistogramDataPointSlice, size int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func splitSummaryDataPoints(src, dst pmetric.SummaryDataPointSlice, size int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

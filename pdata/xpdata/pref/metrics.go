package pref

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func MarkPipelineOwnedMetrics(md pmetric.Metrics) bool { _ = "STUB: not implemented"; return false }

func RefMetrics(md pmetric.Metrics) { _ = "STUB: not implemented"; return }

func UnrefMetrics(md pmetric.Metrics) { _ = "STUB: not implemented"; return }

func EqualMetrics(md1, md2 pmetric.Metrics) bool { _ = "STUB: not implemented"; return false }

package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
)

type AggregationTemporality int32

const (
	AggregationTemporalityUnspecified = AggregationTemporality(internal.AggregationTemporality_AGGREGATION_TEMPORALITY_UNSPECIFIED)

	AggregationTemporalityDelta = AggregationTemporality(internal.AggregationTemporality_AGGREGATION_TEMPORALITY_DELTA)

	AggregationTemporalityCumulative = AggregationTemporality(internal.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE)
)

func (at AggregationTemporality) String() string { _ = "STUB: not implemented"; return "" }

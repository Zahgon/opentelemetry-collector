package extensions

import (
	"gonum.org/v1/gonum/graph"

	"go.opentelemetry.io/collector/component"
)

type node struct {
	nodeID int64
	extID  component.ID
}

func (n node) ID() int64 { _ = "STUB: not implemented"; return 0 }

func computeOrder(exts *Extensions) ([]component.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cycleErr(err error, cycles [][]graph.Node) error { _ = "STUB: not implemented"; return nil }

package featuregate

import (
	"sync/atomic"

	"github.com/hashicorp/go-version"
)

type Gate struct {
	id           string
	description  string
	referenceURL string
	fromVersion  *version.Version
	toVersion    *version.Version
	stage        Stage
	enabled      *atomic.Bool
}

func (g *Gate) ID() string { _ = "STUB: not implemented"; return "" }

func (g *Gate) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (g *Gate) Description() string { _ = "STUB: not implemented"; return "" }

func (g *Gate) Stage() Stage { _ = "STUB: not implemented"; return *new(Stage) }

func (g *Gate) ReferenceURL() string { _ = "STUB: not implemented"; return "" }

func (g *Gate) FromVersion() string { _ = "STUB: not implemented"; return "" }

func (g *Gate) ToVersion() string { _ = "STUB: not implemented"; return "" }

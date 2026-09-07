//go:build js && wasm

package otelcol

import (
	"syscall"
)

const (
	SIGHUP  = syscall.Signal(-1)
	SIGTERM = syscall.Signal(-1)
)

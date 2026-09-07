//go:build !(js && wasm)

package otelcol

import (
	"syscall"
)

const (
	SIGHUP  = syscall.SIGHUP
	SIGTERM = syscall.SIGTERM
)

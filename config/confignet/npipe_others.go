//go:build !windows

package confignet

import (
	"context"
	"errors"
	"net"
	"time"
)

var errNpipeUnsupported = errors.New("npipe transport is only supported on Windows")

func dialNpipe(_ context.Context, _ string, _ time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func listenNpipe(_ string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

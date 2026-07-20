//go:build windows

package confignet

import (
	"context"
	"net"
	"time"
)

func dialNpipe(ctx context.Context, endpoint string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func listenNpipe(endpoint string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

package confignet

import (
	"context"
	"net"
)

const (
	TransportTypeTCP        TransportType = "tcp"
	TransportTypeTCP4       TransportType = "tcp4"
	TransportTypeTCP6       TransportType = "tcp6"
	TransportTypeUDP        TransportType = "udp"
	TransportTypeUDP4       TransportType = "udp4"
	TransportTypeUDP6       TransportType = "udp6"
	TransportTypeIP         TransportType = "ip"
	TransportTypeIP4        TransportType = "ip4"
	TransportTypeIP6        TransportType = "ip6"
	TransportTypeUnix       TransportType = "unix"
	TransportTypeUnixgram   TransportType = "unixgram"
	TransportTypeUnixPacket TransportType = "unixpacket"
	TransportTypeNpipe      TransportType = "npipe"
	transportTypeEmpty      TransportType = ""
)

func (tt *TransportType) UnmarshalText(in []byte) error { _ = "STUB: not implemented"; return nil }

func (na *AddrConfig) Dial(ctx context.Context) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (na *AddrConfig) Listen(ctx context.Context) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (na *AddrConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func validateNpipePath(endpoint string) error { _ = "STUB: not implemented"; return nil }

func (na *TCPAddrConfig) Dial(ctx context.Context) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (na *TCPAddrConfig) Listen(ctx context.Context) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

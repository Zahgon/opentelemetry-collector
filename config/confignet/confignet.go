package confignet

import (
	"context"
	"net"
	"time"
)

type TransportType string

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

type DialerConfig struct {
	Timeout time.Duration `mapstructure:"timeout,omitempty"`

	_ struct{}
}

func NewDefaultDialerConfig() DialerConfig { _ = "STUB: not implemented"; return *new(DialerConfig) }

type AddrConfig struct {
	Endpoint string `mapstructure:"endpoint,omitempty"`

	Transport TransportType `mapstructure:"transport,omitempty"`

	DialerConfig DialerConfig `mapstructure:"dialer,omitempty"`

	_ struct{}
}

func NewDefaultAddrConfig() AddrConfig { _ = "STUB: not implemented"; return *new(AddrConfig) }

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

type TCPAddrConfig struct {
	Endpoint string `mapstructure:"endpoint,omitempty"`

	DialerConfig DialerConfig `mapstructure:"dialer,omitempty"`

	_ struct{}
}

func NewDefaultTCPAddrConfig() TCPAddrConfig { _ = "STUB: not implemented"; return *new(TCPAddrConfig) }

func (na *TCPAddrConfig) Dial(ctx context.Context) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (na *TCPAddrConfig) Listen(ctx context.Context) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

package confignet

import (
	"time"
)

type AddrConfig struct {
	DialerConfig DialerConfig `mapstructure:"dialer,omitempty"`

	Endpoint string `mapstructure:"endpoint,omitempty"`

	Transport TransportType `mapstructure:"transport,omitempty"`

	_ struct{}
}

func NewDefaultAddrConfig() AddrConfig { _ = "STUB: not implemented"; return *new(AddrConfig) }

type DialerConfig struct {
	Timeout time.Duration `mapstructure:"timeout,omitempty"`

	_ struct{}
}

func NewDefaultDialerConfig() DialerConfig { _ = "STUB: not implemented"; return *new(DialerConfig) }

type TCPAddrConfig struct {
	DialerConfig DialerConfig `mapstructure:"dialer,omitempty"`

	Endpoint string `mapstructure:"endpoint,omitempty"`

	_ struct{}
}

func NewDefaultTCPAddrConfig() TCPAddrConfig { _ = "STUB: not implemented"; return *new(TCPAddrConfig) }

type TransportType string

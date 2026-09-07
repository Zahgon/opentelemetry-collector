package confighttp

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configauth"
	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/config/configmiddleware"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap"
)

const (
	headerContentEncoding = "Content-Encoding"
)

type ClientConfig struct {
	Endpoint string `mapstructure:"endpoint,omitempty"`

	ProxyURL string `mapstructure:"proxy_url,omitempty"`

	TLS configtls.ClientConfig `mapstructure:"tls,omitempty"`

	ReadBufferSize int `mapstructure:"read_buffer_size,omitempty"`

	WriteBufferSize int `mapstructure:"write_buffer_size,omitempty"`

	Timeout time.Duration `mapstructure:"timeout,omitempty"`

	Headers configopaque.MapList `mapstructure:"headers,omitempty"`

	Auth configoptional.Optional[configauth.Config] `mapstructure:"auth,omitempty"`

	Compression configcompression.Type `mapstructure:"compression,omitempty"`

	CompressionParams configcompression.CompressionParams `mapstructure:"compression_params,omitempty"`

	MaxConnsPerHost int `mapstructure:"max_conns_per_host,omitempty"`

	HTTP2ReadIdleTimeout time.Duration `mapstructure:"http2_read_idle_timeout,omitempty"`

	HTTP2PingTimeout time.Duration `mapstructure:"http2_ping_timeout,omitempty"`

	Cookies configoptional.Optional[CookiesConfig] `mapstructure:"cookies,omitempty"`

	ForceAttemptHTTP2 bool `mapstructure:"force_attempt_http2,omitempty"`

	Middlewares []configmiddleware.Config `mapstructure:"middlewares,omitempty"`

	Keepalive configoptional.Optional[KeepaliveClientConfig] `mapstructure:"keepalive,omitempty"`

	IdleConnTimeout time.Duration `mapstructure:"idle_conn_timeout,omitempty"`

	MaxIdleConns int `mapstructure:"max_idle_conns,omitempty"`

	MaxIdleConnsPerHost int `mapstructure:"max_idle_conns_per_host,omitempty"`

	DisableKeepAlives bool `mapstructure:"disable_keep_alives,omitempty"`

	deprecationWarnings []string

	_ struct{}
}

type CookiesConfig struct {
	_ struct{}
}

type KeepaliveClientConfig struct {
	IdleConnTimeout time.Duration `mapstructure:"idle_conn_timeout"`

	MaxIdleConns int `mapstructure:"max_idle_conns"`

	MaxIdleConnsPerHost int `mapstructure:"max_idle_conns_per_host,omitempty"`

	_ struct{}
}

func NewDefaultKeepaliveClientConfig() KeepaliveClientConfig {
	_ = "STUB: not implemented"
	return *new(KeepaliveClientConfig)
}

func NewDefaultClientConfig() ClientConfig { _ = "STUB: not implemented"; return *new(ClientConfig) }

var _ confmap.Unmarshaler = (*ClientConfig)(nil)

func (cc *ClientConfig) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

func (cc *ClientConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type ToClientOption interface {
	sealed()
}

func (cc *ClientConfig) ToClient(ctx context.Context, extensions map[component.ID]component.Component, settings component.TelemetrySettings, _ ...ToClientOption) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type headerRoundTripper struct {
	transport http.RoundTripper
	headers   configopaque.MapList
}

func (interceptor *headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

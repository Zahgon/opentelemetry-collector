package confighttp

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configauth"
	"go.opentelemetry.io/collector/config/confighttp/internal"
	"go.opentelemetry.io/collector/config/configmiddleware"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension/extensionauth"
)

const defaultMaxRequestBodySize = 20 * 1024 * 1024

type ServerConfig struct {
	NetAddr confignet.AddrConfig `mapstructure:",squash"`

	TLS configoptional.Optional[configtls.ServerConfig] `mapstructure:"tls"`

	CORS configoptional.Optional[CORSConfig] `mapstructure:"cors"`

	Auth configoptional.Optional[AuthConfig] `mapstructure:"auth,omitempty"`

	MaxRequestBodySize int64 `mapstructure:"max_request_body_size,omitempty"`

	IncludeMetadata bool `mapstructure:"include_metadata,omitempty"`

	ResponseHeaders configopaque.MapList `mapstructure:"response_headers,omitempty"`

	CompressionAlgorithms []string `mapstructure:"compression_algorithms,omitempty"`

	ReadTimeout time.Duration `mapstructure:"read_timeout,omitempty"`

	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`

	WriteTimeout time.Duration `mapstructure:"write_timeout"`

	Middlewares []configmiddleware.Config `mapstructure:"middlewares,omitempty"`

	Keepalive configoptional.Optional[KeepaliveServerConfig] `mapstructure:"keepalive,omitempty"`

	IdleTimeout time.Duration `mapstructure:"idle_timeout,omitempty"`

	KeepAlivesEnabled bool `mapstructure:"keep_alives_enabled,omitempty"`

	deprecationWarnings []string

	_ struct{}
}

type KeepaliveServerConfig struct {
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`

	_ struct{}
}

func NewDefaultKeepaliveServerConfig() KeepaliveServerConfig {
	_ = "STUB: not implemented"
	return *new(KeepaliveServerConfig)
}

func NewDefaultServerConfig() ServerConfig { _ = "STUB: not implemented"; return *new(ServerConfig) }

var _ confmap.Unmarshaler = (*ServerConfig)(nil)

func (sc *ServerConfig) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

type AuthConfig struct {
	Config configauth.Config `mapstructure:",squash"`

	RequestParameters []string `mapstructure:"request_params,omitempty"`

	_ struct{}
}

func (sc *ServerConfig) ToListener(ctx context.Context) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

type toServerOptions = internal.ToServerOptions

type ToServerOption = internal.ToServerOption

func WithErrorHandler(e func(w http.ResponseWriter, r *http.Request, errorMsg string, statusCode int)) ToServerOption {
	_ = "STUB: not implemented"
	return *new(ToServerOption)
}

func WithDecoder(key string, dec func(body io.ReadCloser) (io.ReadCloser, error)) ToServerOption {
	_ = "STUB: not implemented"
	return *new(ToServerOption)
}

func (sc *ServerConfig) ToServer(ctx context.Context, extensions map[component.ID]component.Component, settings component.TelemetrySettings, handler http.Handler, opts ...ToServerOption) (*http.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func responseHeadersHandler(handler http.Handler, headers configopaque.MapList) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins,omitempty"`

	AllowedHeaders []string `mapstructure:"allowed_headers,omitempty"`

	ExposedHeaders []string `mapstructure:"exposed_headers,omitempty"`

	MaxAge int `mapstructure:"max_age,omitempty"`

	_ struct{}
}

func NewDefaultCORSConfig() CORSConfig { _ = "STUB: not implemented"; return *new(CORSConfig) }

func authInterceptor(next http.Handler, server extensionauth.Server, requestParams []string, serverOpts *internal.ToServerOptions) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func maxRequestBodySizeInterceptor(next http.Handler, maxRecvSize int64) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func standardizeHTTPMethod(method, unknown string) string { _ = "STUB: not implemented"; return "" }

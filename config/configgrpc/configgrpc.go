package configgrpc

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configauth"
	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/config/configmiddleware"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap/xconfmap"
	"go.opentelemetry.io/collector/extension/extensionauth"
)

var errMetadataNotFound = errors.New("no request metadata found")

const DefaultBalancerName = "round_robin"

type KeepaliveClientConfig struct {
	Time                time.Duration `mapstructure:"time"`
	Timeout             time.Duration `mapstructure:"timeout"`
	PermitWithoutStream bool          `mapstructure:"permit_without_stream,omitempty"`

	_ struct{}
}

func NewDefaultKeepaliveClientConfig() KeepaliveClientConfig {
	_ = "STUB: not implemented"
	return *new(KeepaliveClientConfig)
}

func BalancerName() string { _ = "STUB: not implemented"; return "" }

var _ xconfmap.Validator = (*ClientConfig)(nil)

type ClientConfig struct {
	Endpoint string `mapstructure:"endpoint,omitempty"`

	Compression configcompression.Type `mapstructure:"compression,omitempty"`

	TLS configtls.ClientConfig `mapstructure:"tls,omitempty"`

	Keepalive configoptional.Optional[KeepaliveClientConfig] `mapstructure:"keepalive,omitempty"`

	ReadBufferSize int `mapstructure:"read_buffer_size,omitempty"`

	WriteBufferSize int `mapstructure:"write_buffer_size,omitempty"`

	WaitForReady bool `mapstructure:"wait_for_ready,omitempty"`

	Headers configopaque.MapList `mapstructure:"headers,omitempty"`

	UserAgent string `mapstructure:"user_agent,omitempty"`

	BalancerName string `mapstructure:"balancer_name"`

	Authority string `mapstructure:"authority,omitempty"`

	Auth configoptional.Optional[configauth.Config] `mapstructure:"auth,omitempty"`

	Middlewares []configmiddleware.Config `mapstructure:"middlewares,omitempty"`

	_ struct{}
}

func NewDefaultClientConfig() ClientConfig { _ = "STUB: not implemented"; return *new(ClientConfig) }

type KeepaliveServerConfig struct {
	ServerParameters  configoptional.Optional[KeepaliveServerParameters]  `mapstructure:"server_parameters,omitempty"`
	EnforcementPolicy configoptional.Optional[KeepaliveEnforcementPolicy] `mapstructure:"enforcement_policy,omitempty"`

	_ struct{}
}

func NewDefaultKeepaliveServerConfig() KeepaliveServerConfig {
	_ = "STUB: not implemented"
	return *new(KeepaliveServerConfig)
}

type KeepaliveServerParameters struct {
	MaxConnectionIdle     time.Duration `mapstructure:"max_connection_idle,omitempty"`
	MaxConnectionAge      time.Duration `mapstructure:"max_connection_age,omitempty"`
	MaxConnectionAgeGrace time.Duration `mapstructure:"max_connection_age_grace,omitempty"`
	Time                  time.Duration `mapstructure:"time,omitempty"`
	Timeout               time.Duration `mapstructure:"timeout,omitempty"`

	_ struct{}
}

func NewDefaultKeepaliveServerParameters() KeepaliveServerParameters {
	_ = "STUB: not implemented"
	return *new(KeepaliveServerParameters)
}

type KeepaliveEnforcementPolicy struct {
	MinTime             time.Duration `mapstructure:"min_time,omitempty"`
	PermitWithoutStream bool          `mapstructure:"permit_without_stream,omitempty"`

	_ struct{}
}

func NewDefaultKeepaliveEnforcementPolicy() KeepaliveEnforcementPolicy {
	_ = "STUB: not implemented"
	return *new(KeepaliveEnforcementPolicy)
}

var _ xconfmap.Validator = (*ServerConfig)(nil)

type ServerConfig struct {
	NetAddr confignet.AddrConfig `mapstructure:",squash"`

	TLS configoptional.Optional[configtls.ServerConfig] `mapstructure:"tls,omitempty"`

	MaxRecvMsgSizeMiB int `mapstructure:"max_recv_msg_size_mib,omitempty"`

	MaxConcurrentStreams uint32 `mapstructure:"max_concurrent_streams,omitempty,omitempty"`

	ReadBufferSize int `mapstructure:"read_buffer_size,omitempty"`

	WriteBufferSize int `mapstructure:"write_buffer_size,omitempty"`

	Keepalive configoptional.Optional[KeepaliveServerConfig] `mapstructure:"keepalive,omitempty"`

	Auth configoptional.Optional[configauth.Config] `mapstructure:"auth,omitempty"`

	IncludeMetadata bool `mapstructure:"include_metadata,omitempty"`

	Middlewares []configmiddleware.Config `mapstructure:"middlewares,omitempty"`

	_ struct{}
}

func NewDefaultServerConfig() ServerConfig { _ = "STUB: not implemented"; return *new(ServerConfig) }

func (cc *ClientConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (cc *ClientConfig) sanitizedEndpoint() string { _ = "STUB: not implemented"; return "" }

func (cc *ClientConfig) grpcDialTarget() string { _ = "STUB: not implemented"; return "" }

func (cc *ClientConfig) isSchemeHTTP() bool { _ = "STUB: not implemented"; return false }

func (cc *ClientConfig) isSchemeHTTPS() bool { _ = "STUB: not implemented"; return false }

type ToClientConnOption interface {
	isToClientConnOption()
}

type grpcDialOptionWrapper struct {
	opt grpc.DialOption
}

func WithGrpcDialOption(opt grpc.DialOption) ToClientConnOption {
	_ = "STUB: not implemented"
	return *new(ToClientConnOption)
}

func (grpcDialOptionWrapper) isToClientConnOption() { _ = "STUB: not implemented"; return }

func (cc *ClientConfig) ToClientConn(
	ctx context.Context,
	extensions map[component.ID]component.Component,
	settings component.TelemetrySettings,
	extraOpts ...ToClientConnOption,
) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *ClientConfig) addHeadersIfAbsent(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (cc *ClientConfig) getGrpcDialOptions(
	ctx context.Context,
	extensions map[component.ID]component.Component,
	settings component.TelemetrySettings,
	extraOpts []ToClientConnOption,
) ([]grpc.DialOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *ServerConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type ToServerOption interface {
	isToServerOption()
}

type grpcServerOptionWrapper struct {
	opt grpc.ServerOption
}

func WithGrpcServerOption(opt grpc.ServerOption) ToServerOption {
	_ = "STUB: not implemented"
	return *new(ToServerOption)
}

func (grpcServerOptionWrapper) isToServerOption() { _ = "STUB: not implemented"; return }

func (sc *ServerConfig) ToServer(
	ctx context.Context,
	extensions map[component.ID]component.Component,
	settings component.TelemetrySettings,
	extraOpts ...ToServerOption,
) (*grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sc *ServerConfig) getGrpcServerOptions(
	ctx context.Context,
	extensions map[component.ID]component.Component,
	settings component.TelemetrySettings,
	extraOpts []ToServerOption,
) ([]grpc.ServerOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:contextcheck // context already handled

//nolint:contextcheck // context already handled

func getGRPCCompressionName(compressionType configcompression.Type) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func enhanceWithClientInformation(includeMetadata bool) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func enhanceStreamWithClientInformation(includeMetadata bool) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func contextWithClient(ctx context.Context, includeMetadata bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func authUnaryServerInterceptor(server extensionauth.Server) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func authStreamServerInterceptor(server extensionauth.Server) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

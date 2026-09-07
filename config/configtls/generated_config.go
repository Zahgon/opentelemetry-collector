package configtls

import (
	"time"

	"go.opentelemetry.io/collector/config/configopaque"
)

type ClientConfig struct {
	Config `mapstructure:",squash"`

	Insecure bool `mapstructure:"insecure,omitempty"`

	InsecureSkipVerify bool `mapstructure:"insecure_skip_verify,omitempty"`

	ServerName string `mapstructure:"server_name_override,omitempty"`

	_ struct{}
}

func NewDefaultClientConfig() ClientConfig { _ = "STUB: not implemented"; return *new(ClientConfig) }

type Config struct {
	CAFile string `mapstructure:"ca_file,omitempty"`

	CAPem configopaque.String `mapstructure:"ca_pem,omitempty"`

	CertFile string `mapstructure:"cert_file,omitempty"`

	CertPem configopaque.String `mapstructure:"cert_pem,omitempty"`

	CipherSuites []string `mapstructure:"cipher_suites,omitempty"`

	CurvePreferences []string `mapstructure:"curve_preferences,omitempty"`

	IncludeInsecureCipherSuites bool `mapstructure:"include_insecure_cipher_suites,omitempty"`

	IncludeSystemCACertsPool bool `mapstructure:"include_system_ca_certs_pool,omitempty"`

	KeyFile string `mapstructure:"key_file,omitempty"`

	KeyPem configopaque.String `mapstructure:"key_pem,omitempty"`

	MaxVersion string `mapstructure:"max_version,omitempty"`

	MinVersion string `mapstructure:"min_version,omitempty"`

	ReloadInterval time.Duration `mapstructure:"reload_interval,omitempty"`

	TPMConfig TPMConfig `mapstructure:"tpm,omitempty"`

	_ struct{}
}

func NewDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

type ServerConfig struct {
	Config `mapstructure:",squash"`

	ClientCAFile string `mapstructure:"client_ca_file,omitempty"`

	ReloadClientCAFile bool `mapstructure:"client_ca_file_reload,omitempty"`

	_ struct{}
}

func NewDefaultServerConfig() ServerConfig { _ = "STUB: not implemented"; return *new(ServerConfig) }

type TPMConfig struct {
	Auth string `mapstructure:"auth,omitempty"`

	Enabled bool `mapstructure:"enabled,omitempty"`

	OwnerAuth string `mapstructure:"owner_auth,omitempty"`

	Path string `mapstructure:"path,omitempty"`

	_ struct{}
}

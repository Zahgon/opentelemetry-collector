package configtls

import (
	"crypto/tls"

	"github.com/google/go-tpm/tpm2/transport"
)

type TPMConfig struct {
	Enabled bool `mapstructure:"enabled"`

	Path      string `mapstructure:"path"`
	OwnerAuth string `mapstructure:"owner_auth"`
	Auth      string `mapstructure:"auth"`

	_ struct{}
}

func (c TPMConfig) tpmCertificate(keyPem, certPem []byte, openTPM func() (transport.TPMCloser, error)) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

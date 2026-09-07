package configtls

import (
	"crypto/tls"

	"github.com/google/go-tpm/tpm2/transport"
)

func (c TPMConfig) tpmCertificate(keyPem, certPem []byte, openTPM func() (transport.TPMCloser, error)) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

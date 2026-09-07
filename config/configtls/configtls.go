package configtls

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"sync"
	"time"
)

const defaultMinTLSVersion = tls.VersionTLS12

const defaultMaxTLSVersion = 0

var systemCertPool = x509.SystemCertPool

type certReloader struct {
	nextReload time.Time
	cert       *tls.Certificate
	lock       sync.RWMutex
	tls        Config
}

func (c Config) newCertReloader() (*certReloader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *certReloader) GetCertificate() (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (c ServerConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c Config) loadTLSConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func convertCipherSuites(cipherSuites []string, includeInsecure bool) ([]uint16, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) loadCACertPool() (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) loadCertFile(certPath string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) loadCertPem(certPem []byte) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) loadCertificate() (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

func (c Config) loadCert(caPath string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c ClientConfig) LoadTLSConfig(_ context.Context) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c ServerConfig) LoadTLSConfig(_ context.Context) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c ServerConfig) loadClientCAFile() (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) hasCA() bool   { _ = "STUB: not implemented"; return false }
func (c Config) hasCert() bool { _ = "STUB: not implemented"; return false }
func (c Config) hasKey() bool  { _ = "STUB: not implemented"; return false }

func (c Config) hasCAFile() bool { _ = "STUB: not implemented"; return false }
func (c Config) hasCAPem() bool  { _ = "STUB: not implemented"; return false }

func (c Config) hasCertFile() bool { _ = "STUB: not implemented"; return false }
func (c Config) hasCertPem() bool  { _ = "STUB: not implemented"; return false }

func (c Config) hasKeyFile() bool { _ = "STUB: not implemented"; return false }
func (c Config) hasKeyPem() bool  { _ = "STUB: not implemented"; return false }

func convertVersion(v string, defaultVersion uint16) (uint16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var tlsVersions = map[string]uint16{
	"1.0": tls.VersionTLS10,
	"1.1": tls.VersionTLS11,
	"1.2": tls.VersionTLS12,
	"1.3": tls.VersionTLS13,
}

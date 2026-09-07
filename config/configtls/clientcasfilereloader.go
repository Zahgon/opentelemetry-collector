package configtls

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"sync"
	"time"
)

const defaultClientCAsReloadInterval = time.Second

type clientCAsFileReloader struct {
	clientCAsFile string
	loader        clientCAsFileLoader

	reloadInterval time.Duration

	lock            sync.Mutex
	certPool        *x509.CertPool
	lastReloadError error

	lastCheck time.Time

	fileID fileIdentity
}

type fileIdentity struct {
	exists bool
	sum    [sha256.Size]byte
}

func identifyFile(path string) fileIdentity { _ = "STUB: not implemented"; return *new(fileIdentity) }

type clientCAsFileLoader interface {
	loadClientCAFile() (*x509.CertPool, error)
}

func newClientCAsReloader(clientCAsFile string, loader clientCAsFileLoader) (*clientCAsFileReloader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *clientCAsFileReloader) getClientConfig(original *tls.Config) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *clientCAsFileReloader) reloadIfModified(now time.Time) { _ = "STUB: not implemented"; return }

func (r *clientCAsFileReloader) getLastError() error { _ = "STUB: not implemented"; return nil }

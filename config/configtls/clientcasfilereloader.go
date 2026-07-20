package configtls

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type clientCAsFileReloader struct {
	clientCAsFile   string
	certPool        *x509.CertPool
	lastReloadError error
	lock            sync.RWMutex
	loader          clientCAsFileLoader
	watcher         *fsnotify.Watcher
	shutdownCH      chan bool
}

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

func (r *clientCAsFileReloader) reload() { _ = "STUB: not implemented"; return }

func (r *clientCAsFileReloader) getLastError() error { _ = "STUB: not implemented"; return nil }

func (r *clientCAsFileReloader) startWatching() error { _ = "STUB: not implemented"; return nil }

func (r *clientCAsFileReloader) handleWatcherEvents() { _ = "STUB: not implemented"; return }

func (r *clientCAsFileReloader) shutdown() error { _ = "STUB: not implemented"; return nil }

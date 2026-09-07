//go:build linux

package configtls

import (
	"github.com/google/go-tpm/tpm2/transport"
)

var tpmSimulator transport.TPMCloser

func openTPM(path string) func() (transport.TPMCloser, error) {
	_ = "STUB: not implemented"
	return nil
}

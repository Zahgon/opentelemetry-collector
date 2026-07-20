//go:build windows

package configtls

import (
	"github.com/google/go-tpm/tpm2/transport"
)

func openTPM(_ string) func() (transport.TPMCloser, error) { _ = "STUB: not implemented"; return nil }

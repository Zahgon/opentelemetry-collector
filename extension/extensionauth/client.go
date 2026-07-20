package extensionauth

import (
	"net/http"

	"google.golang.org/grpc/credentials"
)

type HTTPClient interface {
	RoundTripper(base http.RoundTripper) (http.RoundTripper, error)
}

type GRPCClient interface {
	PerRPCCredentials() (credentials.PerRPCCredentials, error)
}

var _ HTTPClient = (*ClientRoundTripperFunc)(nil)

type ClientRoundTripperFunc func(base http.RoundTripper) (http.RoundTripper, error)

func (f ClientRoundTripperFunc) RoundTripper(base http.RoundTripper) (http.RoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper), nil
}

var _ GRPCClient = (*ClientPerRPCCredentialsFunc)(nil)

type ClientPerRPCCredentialsFunc func() (credentials.PerRPCCredentials, error)

func (f ClientPerRPCCredentialsFunc) PerRPCCredentials() (credentials.PerRPCCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.PerRPCCredentials), nil
}

package extensionmiddlewaretest

import (
	"net/http"

	"go.opentelemetry.io/collector/extension"
)

func NewNop() extension.Extension { _ = "STUB: not implemented"; return *new(extension.Extension) }

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

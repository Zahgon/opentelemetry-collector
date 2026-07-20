package confighttp

import (
	"context"
	"net"
	"net/http"
)

type clientInfoHandler struct {
	next http.Handler

	includeMetadata bool
}

func (h *clientInfoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

//nolint:contextcheck //context already handled through contextWithClient

func contextWithClient(req *http.Request, includeMetadata bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func parseIP(source string) *net.IPAddr { _ = "STUB: not implemented"; return nil }

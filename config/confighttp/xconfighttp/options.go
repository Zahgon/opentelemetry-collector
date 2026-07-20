package xconfighttp

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"go.opentelemetry.io/collector/config/confighttp"
)

func WithOtelHTTPOptions(httpopts ...otelhttp.Option) confighttp.ToServerOption {
	_ = "STUB: not implemented"
	return *new(confighttp.ToServerOption)
}

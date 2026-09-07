package exporterhelper

import (
	"errors"
)

var (
	errNilConfig = errors.New("nil config")

	errNilLogger = errors.New("nil logger")

	errNilPushTraces = errors.New("nil PushTraces")

	errNilPushMetrics = errors.New("nil PushMetrics")

	errNilPushLogs = errors.New("nil PushLogs")
)

package otlpreceiver

import (
	"net/http"

	"google.golang.org/grpc/status"

	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/logs"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/metrics"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/profiles"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/trace"
)

var fallbackMsg = []byte(`{"code": 13, "message": "failed to marshal error message"}`)

const fallbackContentType = "application/json"

func handleTraces(resp http.ResponseWriter, req *http.Request, tracesReceiver *trace.Receiver) {
	_ = "STUB: not implemented"
	return
}

func handleMetrics(resp http.ResponseWriter, req *http.Request, metricsReceiver *metrics.Receiver) {
	_ = "STUB: not implemented"
	return
}

func handleLogs(resp http.ResponseWriter, req *http.Request, logsReceiver *logs.Receiver) {
	_ = "STUB: not implemented"
	return
}

func handleProfiles(resp http.ResponseWriter, req *http.Request, profilesReceiver *profiles.Receiver) {
	_ = "STUB: not implemented"
	return
}

func readContentType(resp http.ResponseWriter, req *http.Request) (encoder, bool) {
	_ = "STUB: not implemented"
	return *new(encoder), false
}

func readAndCloseBody(resp http.ResponseWriter, req *http.Request, enc encoder) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func writeError(w http.ResponseWriter, encoder encoder, err error, statusCode int) {
	_ = "STUB: not implemented"
	return
}

func errorHandler(w http.ResponseWriter, r *http.Request, errMsg string, statusCode int) {
	_ = "STUB: not implemented"
	return
}

func writeStatusResponse(w http.ResponseWriter, enc encoder, statusCode int, st *status.Status) {
	_ = "STUB: not implemented"
	return
}

func writeResponse(w http.ResponseWriter, contentType string, statusCode int, msg []byte) {
	_ = "STUB: not implemented"
	return
}

func getMimeTypeFromContentType(contentType string) string { _ = "STUB: not implemented"; return "" }

func handleUnmatchedMethod(resp http.ResponseWriter) { _ = "STUB: not implemented"; return }

func handleUnmatchedContentType(resp http.ResponseWriter) { _ = "STUB: not implemented"; return }

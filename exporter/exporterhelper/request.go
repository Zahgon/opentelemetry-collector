package exporterhelper

import (
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/request"
)

type RequestSizerType = request.SizerType

var (
	RequestSizerTypeBytes    = request.SizerTypeBytes
	RequestSizerTypeItems    = request.SizerTypeItems
	RequestSizerTypeRequests = request.SizerTypeRequests
)

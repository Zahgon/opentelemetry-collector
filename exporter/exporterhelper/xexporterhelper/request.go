package xexporterhelper

import "go.opentelemetry.io/collector/exporter/exporterhelper/internal/request"

type Request = request.Request

type RequestErrorHandler = request.ErrorHandler

type RequestConverterFunc[T any] = request.RequestConverterFunc[T]

type RequestConsumeFunc = request.RequestConsumeFunc

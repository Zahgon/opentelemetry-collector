package otlpreceiver

import (
	spb "google.golang.org/genproto/googleapis/rpc/status"

	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/pprofile/pprofileotlp"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
)

const (
	pbContentType   = "application/x-protobuf"
	jsonContentType = "application/json"
)

var (
	pbEncoder = &protoEncoder{}
	jsEncoder = &jsonEncoder{}
)

type encoder interface {
	unmarshalTracesRequest(buf []byte) (ptraceotlp.ExportRequest, error)
	unmarshalMetricsRequest(buf []byte) (pmetricotlp.ExportRequest, error)
	unmarshalLogsRequest(buf []byte) (plogotlp.ExportRequest, error)
	unmarshalProfilesRequest(buf []byte) (pprofileotlp.ExportRequest, error)

	marshalTracesResponse(ptraceotlp.ExportResponse) ([]byte, error)
	marshalMetricsResponse(pmetricotlp.ExportResponse) ([]byte, error)
	marshalLogsResponse(plogotlp.ExportResponse) ([]byte, error)
	marshalProfilesResponse(pprofileotlp.ExportResponse) ([]byte, error)

	marshalStatus(rsp *spb.Status) ([]byte, error)

	contentType() string
}

type protoEncoder struct{}

func (protoEncoder) unmarshalTracesRequest(buf []byte) (ptraceotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(ptraceotlp.ExportRequest), nil
}

func (protoEncoder) unmarshalMetricsRequest(buf []byte) (pmetricotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(pmetricotlp.ExportRequest), nil
}

func (protoEncoder) unmarshalLogsRequest(buf []byte) (plogotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(plogotlp.ExportRequest), nil
}

func (protoEncoder) unmarshalProfilesRequest(buf []byte) (pprofileotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(pprofileotlp.ExportRequest), nil
}

func (protoEncoder) marshalTracesResponse(resp ptraceotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (protoEncoder) marshalMetricsResponse(resp pmetricotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (protoEncoder) marshalLogsResponse(resp plogotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (protoEncoder) marshalProfilesResponse(resp pprofileotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (protoEncoder) marshalStatus(resp *spb.Status) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (protoEncoder) contentType() string { _ = "STUB: not implemented"; return "" }

type jsonEncoder struct{}

func (jsonEncoder) unmarshalTracesRequest(buf []byte) (ptraceotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(ptraceotlp.ExportRequest), nil
}

func (jsonEncoder) unmarshalMetricsRequest(buf []byte) (pmetricotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(pmetricotlp.ExportRequest), nil
}

func (jsonEncoder) unmarshalLogsRequest(buf []byte) (plogotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(plogotlp.ExportRequest), nil
}

func (jsonEncoder) unmarshalProfilesRequest(buf []byte) (pprofileotlp.ExportRequest, error) {
	_ = "STUB: not implemented"
	return *new(pprofileotlp.ExportRequest), nil
}

func (jsonEncoder) marshalTracesResponse(resp ptraceotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jsonEncoder) marshalMetricsResponse(resp pmetricotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jsonEncoder) marshalLogsResponse(resp plogotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jsonEncoder) marshalProfilesResponse(resp pprofileotlp.ExportResponse) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jsonEncoder) marshalStatus(resp *spb.Status) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (jsonEncoder) contentType() string { _ = "STUB: not implemented"; return "" }

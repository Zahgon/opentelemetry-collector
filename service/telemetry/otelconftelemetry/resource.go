package otelconftelemetry

import (
	"context"
	"errors"

	"github.com/google/uuid"
	otelconf "go.opentelemetry.io/contrib/otelconf/v0.3.0"
	xotelconf "go.opentelemetry.io/contrib/otelconf/x"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/service/telemetry"
)

var errMissingCollectorResource = errors.New("collector resource must be initialized before creating telemetry providers")

var defaultAttributeValues = func(buildInfo component.BuildInfo) (map[string]string, error) {
	instanceUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return map[string]string{
		string(semconv.ServiceNameKey):       buildInfo.Command,
		string(semconv.ServiceVersionKey):    buildInfo.Version,
		string(semconv.ServiceInstanceIDKey): instanceUUID.String(),
	}, nil
}

var newExperimentalSDK = xotelconf.NewSDK

func createInitialResourceConfig(buildInfo component.BuildInfo, cfg *ResourceConfig) (*otelconf.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createResource(
	ctx context.Context,
	set telemetry.Settings,
	componentConfig component.Config,
) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

func putSDKAttribute(attrs pcommon.Map, key string, value attribute.Value) {
	_ = "STUB: not implemented"
	return
}

func newDetectionResourceSDK(ctx context.Context, detection *xotelconf.ExperimentalResourceDetection) (xotelconf.SDK, error) {
	_ = "STUB: not implemented"
	return *new(xotelconf.SDK), nil
}

func createFixedResourceConfig(res *pcommon.Resource, schemaURL string) (*otelconf.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

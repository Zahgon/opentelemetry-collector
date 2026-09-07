package xexporterhelper

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/queue"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/request"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/sizer"
	"go.opentelemetry.io/collector/exporter/xexporter"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

var (
	profilesMarshaler   = &pprofile.ProtoMarshaler{}
	profilesUnmarshaler = &pprofile.ProtoUnmarshaler{}
)

func NewProfilesQueueBatchSettings() QueueBatchSettings {
	_ = "STUB: not implemented"
	return *new(QueueBatchSettings)
}

var (
	_ request.Request      = (*profilesRequest)(nil)
	_ request.ErrorHandler = (*profilesRequest)(nil)
)

type profilesRequest struct {
	pd         pprofile.Profiles
	cachedSize int
}

func newProfilesRequest(pd pprofile.Profiles) Request {
	_ = "STUB: not implemented"
	return *new(Request)
}

type profilesEncoding struct{}

var _ exporterhelper.QueueBatchEncoding[request.Request] = profilesEncoding{}

func (profilesEncoding) Unmarshal(bytes []byte) (context.Context, request.Request, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(request.Request), nil
}

func (profilesEncoding) Marshal(ctx context.Context, req request.Request) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ queue.ReferenceCounter[request.Request] = profilesReferenceCounter{}

type profilesReferenceCounter struct{}

func (profilesReferenceCounter) Ref(req request.Request) { _ = "STUB: not implemented"; return }

func (profilesReferenceCounter) Unref(req request.Request) { _ = "STUB: not implemented"; return }

func (req *profilesRequest) OnError(err error) Request {
	_ = "STUB: not implemented"
	return *new(Request)
}

func (req *profilesRequest) ItemsCount() int { _ = "STUB: not implemented"; return 0 }

func (req *profilesRequest) size(sizer sizer.ProfilesSizer) int {
	_ = "STUB: not implemented"
	return 0
}

func (req *profilesRequest) setCachedSize(size int) { _ = "STUB: not implemented"; return }

func (req *profilesRequest) BytesSize() int { _ = "STUB: not implemented"; return 0 }

type profileExporter struct {
	*internal.BaseExporter
	xconsumer.Profiles
}

func NewProfiles(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
	pusher xconsumer.ConsumeProfilesFunc,
	options ...exporterhelper.Option,
) (xexporter.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xexporter.Profiles), nil
}

func requestConsumeFromProfiles(pusher xconsumer.ConsumeProfilesFunc) RequestConsumeFunc {
	_ = "STUB: not implemented"
	return *new(RequestConsumeFunc)
}

func requestFromProfiles() RequestConverterFunc[pprofile.Profiles] {
	_ = "STUB: not implemented"
	return nil
}

func NewProfilesRequest(
	_ context.Context,
	set exporter.Settings,
	converter RequestConverterFunc[pprofile.Profiles],
	pusher RequestConsumeFunc,
	options ...exporterhelper.Option,
) (xexporter.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xexporter.Profiles), nil
}

func newConsumeProfiles(converter RequestConverterFunc[pprofile.Profiles], be *internal.BaseExporter, logger *zap.Logger) xconsumer.ConsumeProfilesFunc {
	_ = "STUB: not implemented"
	return *new(xconsumer.ConsumeProfilesFunc)
}

package confighttp

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"net/http"
	"sync"

	"github.com/golang/snappy"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v4"

	"go.opentelemetry.io/collector/config/configcompression"
)

func defaultCompressionAlgorithms() []string { _ = "STUB: not implemented"; return nil }

type compressRoundTripper struct {
	rt                http.RoundTripper
	compressionType   configcompression.Type
	compressionParams configcompression.CompressionParams
	compressor        *compressor
}

var zstdReaderPool sync.Pool

type pooledZstdReadCloser struct {
	inner *zstd.Decoder
}

func (pzrc *pooledZstdReadCloser) Read(dst []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pzrc *pooledZstdReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

var availableDecoders = map[string]func(body io.ReadCloser) (io.ReadCloser, error){
	"": func(io.ReadCloser) (io.ReadCloser, error) {

		return nil, nil
	},
	"gzip": func(body io.ReadCloser) (io.ReadCloser, error) {
		gr, err := gzip.NewReader(body)
		if err != nil {
			return nil, err
		}
		return gr, nil
	},
	"zstd": func(body io.ReadCloser) (io.ReadCloser, error) {
		v := zstdReaderPool.Get()
		var zr *zstd.Decoder
		var err error
		if v == nil {

			zr, err = zstd.NewReader(body, zstd.WithDecoderConcurrency(1))
		} else {
			zr = v.(*zstd.Decoder)
			err = zr.Reset(body)
		}
		if err != nil {
			return nil, err
		}
		return &pooledZstdReadCloser{inner: zr}, nil
	},
	"zlib": func(body io.ReadCloser) (io.ReadCloser, error) {
		zr, err := zlib.NewReader(body)
		if err != nil {
			return nil, err
		}
		return zr, nil
	},
	"snappy": newSnappyHandler(0),
	//nolint:unparam // Ignoring the linter request to remove error return since it needs to match the method signature
	"lz4": func(body io.ReadCloser) (io.ReadCloser, error) {
		return &compressReadCloser{
			Reader: lz4.NewReader(body),
			orig:   body,
		}, nil
	},
	//nolint:unparam // Ignoring the linter request to remove error return since it needs to match the method signature
	"x-snappy-framed": func(body io.ReadCloser) (io.ReadCloser, error) {
		return &compressReadCloser{
			Reader: snappy.NewReader(body),
			orig:   body,
		}, nil
	},
}

var snappyFramingHeader = []byte{
	0xff, 0x06, 0x00, 0x00,
	0x73, 0x4e, 0x61, 0x50, 0x70, 0x59,
}

func newSnappyHandler(maxRequestBodySize int64) func(io.ReadCloser) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return nil
}

func newCompressionParams(level configcompression.Level) configcompression.CompressionParams {
	_ = "STUB: not implemented"
	return *new(configcompression.CompressionParams)
}

func newCompressRoundTripper(rt http.RoundTripper, compressionType configcompression.Type, compressionParams configcompression.CompressionParams) (*compressRoundTripper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *compressRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type decompressor struct {
	errHandler         func(w http.ResponseWriter, r *http.Request, errorMsg string, statusCode int)
	base               http.Handler
	decoders           map[string]func(body io.ReadCloser) (io.ReadCloser, error)
	maxRequestBodySize int64
}

func httpContentDecompressor(h http.Handler, maxRequestBodySize int64, eh func(w http.ResponseWriter, r *http.Request, errorMsg string, statusCode int), enableDecoders []string, decoders map[string]func(body io.ReadCloser) (io.ReadCloser, error)) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (d *decompressor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (d *decompressor) newBodyReader(r *http.Request) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func defaultErrorHandler(w http.ResponseWriter, _ *http.Request, errMsg string, statusCode int) {
	_ = "STUB: not implemented"
	return
}

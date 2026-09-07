package confighttp

import (
	"bytes"
	"io"
	"sync"

	"go.opentelemetry.io/collector/config/configcompression"
)

type writeCloserReset interface {
	io.WriteCloser
	Reset(w io.Writer)
}

type compressor struct {
	pool sync.Pool
}

type compressorMap map[compressionMapKey]*compressor

type compressionMapKey struct {
	compressionType   configcompression.Type
	compressionParams configcompression.CompressionParams
}

var (
	compressorPools   = make(compressorMap)
	compressorPoolsMu sync.Mutex
)

func newCompressor(compressionType configcompression.Type, compressionParams configcompression.CompressionParams) (*compressor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWriteCloserResetFunc(compressionType configcompression.Type, compressionParams configcompression.CompressionParams) (func() writeCloserReset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *compressor) compress(buf *bytes.Buffer, body io.ReadCloser) error {
	_ = "STUB: not implemented"
	return nil
}

type rawSnappyWriter struct {
	buffer bytes.Buffer
	w      io.Writer
	closed bool
}

func (w *rawSnappyWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *rawSnappyWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (w *rawSnappyWriter) Reset(newWriter io.Writer) { _ = "STUB: not implemented"; return }

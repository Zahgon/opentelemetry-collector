package confighttp

import (
	"io"
)

type compressReadCloser struct {
	io.Reader
	orig io.ReadCloser
}

var (
	_ io.Reader = (*compressReadCloser)(nil)
	_ io.Closer = (*compressReadCloser)(nil)
)

func (crc *compressReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

type panicRecoverReadCloser struct {
	inner io.ReadCloser
}

func (pr *panicRecoverReadCloser) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pr *panicRecoverReadCloser) Close() (err error) { _ = "STUB: not implemented"; return nil }

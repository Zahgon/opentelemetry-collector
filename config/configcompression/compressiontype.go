package configcompression

import (
	"compress/zlib"
)

const (
	TypeGzip                Type = "gzip"
	TypeZlib                Type = "zlib"
	TypeDeflate             Type = "deflate"
	TypeSnappy              Type = "snappy"
	TypeSnappyFramed        Type = "x-snappy-framed"
	TypeZstd                Type = "zstd"
	TypeLz4                 Type = "lz4"
	typeNone                Type = "none"
	typeEmpty               Type = ""
	DefaultCompressionLevel      = zlib.DefaultCompression
)

func (ct *Type) IsCompressed() bool { _ = "STUB: not implemented"; return false }

func (ct *Type) UnmarshalText(in []byte) error { _ = "STUB: not implemented"; return nil }

func (ct *Type) ValidateParams(p CompressionParams) error { _ = "STUB: not implemented"; return nil }

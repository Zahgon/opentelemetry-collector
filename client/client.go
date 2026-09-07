package client

import (
	"context"
	"iter"
	"net"
)

type ctxKey struct{}

type Info struct {
	Addr net.Addr

	Auth AuthData

	Metadata Metadata

	_ struct{}
}

type AuthData interface {
	GetAttribute(string) any

	GetAttributeNames() []string
}

const MetadataHostName = "Host"

func NewContext(ctx context.Context, c Info) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) Info { _ = "STUB: not implemented"; return *new(Info) }

type Metadata struct {
	data map[string][]string
}

func NewMetadata(md map[string][]string) Metadata { _ = "STUB: not implemented"; return *new(Metadata) }

func (m Metadata) Keys() iter.Seq[string] { _ = "STUB: not implemented"; return nil }

func (m Metadata) Get(key string) []string { _ = "STUB: not implemented"; return nil }

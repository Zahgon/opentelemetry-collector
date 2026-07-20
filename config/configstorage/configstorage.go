package configstorage

//go:generate mdatagen metadata.yaml

import (
	"context"
	"encoding"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/xextension/storage"
)

var (
	errNoStorageClient    = errors.New("no storage client extension found")
	errWrongExtensionType = errors.New("requested extension is not a storage extension")
)

type ID component.ID

var (
	_ encoding.TextMarshaler   = ID{}
	_ encoding.TextUnmarshaler = (*ID)(nil)
)

func (id ID) ComponentID() component.ID { _ = "STUB: not implemented"; return *new(component.ID) }

func (id ID) GetStorageClient(ctx context.Context, kind component.Kind, host component.Host, ownerID component.ID, name string) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

func (id ID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (id *ID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

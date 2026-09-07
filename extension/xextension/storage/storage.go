package storage

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

type Extension interface {
	extension.Extension

	GetClient(ctx context.Context, kind component.Kind, id component.ID, storageName string) (Client, error)
}

type Client interface {
	Get(ctx context.Context, key string) ([]byte, error)

	Set(ctx context.Context, key string, value []byte) error

	Delete(ctx context.Context, key string) error

	Batch(ctx context.Context, ops ...*Operation) error

	Close(ctx context.Context) error
}

type WalkFunc func(key string, value []byte) ([]*Operation, error)

var SkipAll = errors.New("skip everything and stop the walk") //nolint:revive,staticcheck // mimics the existing API in the standard library, see filepath

type Walker interface {
	Walk(ctx context.Context, fn WalkFunc) error
}

type OpType int

const (
	Get OpType = iota
	Set
	Delete
)

type Operation struct {
	Key string

	Value []byte

	Type OpType
}

func SetOperation(key string, value []byte) *Operation { _ = "STUB: not implemented"; return nil }

func GetOperation(key string) *Operation { _ = "STUB: not implemented"; return nil }

func DeleteOperation(key string) *Operation { _ = "STUB: not implemented"; return nil }

var ErrStorageFull = errors.New("the storage extension has run out of available space")

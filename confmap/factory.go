package confmap

type moduleFactory[T any, S any] interface {
	Create(s S) T
}

type createConfmapFunc[T any, S any] func(s S) T

type confmapModuleFactory[T any, S any] struct {
	f createConfmapFunc[T, S]
}

func (c confmapModuleFactory[T, S]) Create(s S) T { _ = "STUB: not implemented"; return *new(T) }

func newConfmapModuleFactory[T, S any](f createConfmapFunc[T, S]) moduleFactory[T, S] {
	_ = "STUB: not implemented"
	return nil
}

package featuregate

import (
	"errors"
	"regexp"
	"sync"
)

var (
	globalRegistry = NewRegistry()

	idRegexp = regexp.MustCompile(`^[0-9a-zA-Z.]*$`)
)

var ErrAlreadyRegistered = errors.New("gate is already registered")

func GlobalRegistry() *Registry { _ = "STUB: not implemented"; return nil }

type Registry struct {
	gates sync.Map
}

func NewRegistry() *Registry { _ = "STUB: not implemented"; return nil }

type RegisterOption interface {
	apply(g *Gate) error
}

type registerOptionFunc func(g *Gate) error

func (ro registerOptionFunc) apply(g *Gate) error { _ = "STUB: not implemented"; return nil }

func WithRegisterDescription(description string) RegisterOption {
	_ = "STUB: not implemented"
	return *new(RegisterOption)
}

func WithRegisterReferenceURL(referenceURL string) RegisterOption {
	_ = "STUB: not implemented"
	return *new(RegisterOption)
}

func WithRegisterFromVersion(fromVersion string) RegisterOption {
	_ = "STUB: not implemented"
	return *new(RegisterOption)
}

func WithRegisterToVersion(toVersion string) RegisterOption {
	_ = "STUB: not implemented"
	return *new(RegisterOption)
}

func (r *Registry) MustRegister(id string, stage Stage, opts ...RegisterOption) *Gate {
	_ = "STUB: not implemented"
	return nil
}

func validateID(id string) error { _ = "STUB: not implemented"; return nil }

func (r *Registry) Register(id string, stage Stage, opts ...RegisterOption) (*Gate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) Set(id string, enabled bool) error { _ = "STUB: not implemented"; return nil }

func (r *Registry) VisitAll(fn func(*Gate)) { _ = "STUB: not implemented"; return }

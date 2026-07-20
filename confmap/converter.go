package confmap

import (
	"context"

	"go.uber.org/zap"
)

type ConverterSettings struct {
	Logger *zap.Logger

	_ struct{}
}

type ConverterFactory = moduleFactory[Converter, ConverterSettings]

type CreateConverterFunc = createConfmapFunc[Converter, ConverterSettings]

func NewConverterFactory(f CreateConverterFunc) ConverterFactory {
	_ = "STUB: not implemented"
	return *new(ConverterFactory)
}

type Converter interface {
	Convert(ctx context.Context, conf *Conf) error
}

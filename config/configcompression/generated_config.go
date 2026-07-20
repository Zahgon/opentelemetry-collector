package configcompression

type CompressionParams struct {
	Level Level `mapstructure:"level"`

	_ struct{}
}

type Level int

type Type string

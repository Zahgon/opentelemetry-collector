package configcompression

type CompressionParams struct {
	Level Level `mapstructure:"level,omitempty"`

	_ struct{}
}

type Level int

type Type string

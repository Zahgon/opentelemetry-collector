package confmaptest

import (
	"regexp"

	"go.opentelemetry.io/collector/confmap"
)

func LoadConf(fileName string) (*confmap.Conf, error) { _ = "STUB: not implemented"; return nil, nil }

var schemeValidator = regexp.MustCompile("^[A-Za-z][A-Za-z0-9+.-]+$")

func ValidateProviderScheme(p confmap.Provider) error { _ = "STUB: not implemented"; return nil }

package otelcol

import (
	"regexp"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/confmap"
)

var redactedMask = func() string {
	b, _ := configopaque.String("").MarshalText()
	return string(b)
}()

func redactByMirroring(raw, redacted *confmap.Conf) *confmap.Conf {
	_ = "STUB: not implemented"
	return nil
}

func applyMask(raw, redacted any) any { _ = "STUB: not implemented"; return *new(any) }

var providerReferenceRegexp = regexp.MustCompile(`^\$\{[^}]*\}$`)

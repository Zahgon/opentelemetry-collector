package component

type BuildInfo struct {
	Command string

	Description string

	Version string

	_ struct{}
}

func NewDefaultBuildInfo() BuildInfo { _ = "STUB: not implemented"; return *new(BuildInfo) }

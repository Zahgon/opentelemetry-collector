package featuregate

type Stage int8

const (
	StageAlpha Stage = iota

	StageBeta

	StageStable

	StageDeprecated
)

func (s Stage) String() string { _ = "STUB: not implemented"; return "" }

package configtelemetry

const (
	LevelNone Level = iota - 1

	LevelBasic

	LevelNormal

	LevelDetailed

	levelNoneStr     = "None"
	levelBasicStr    = "Basic"
	levelNormalStr   = "Normal"
	levelDetailedStr = "Detailed"
)

type Level int32

func (l Level) String() string { _ = "STUB: not implemented"; return "" }

func (l Level) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (l *Level) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

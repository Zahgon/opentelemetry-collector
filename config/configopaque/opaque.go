package configopaque

type String string

const maskedString = "[REDACTED]"

func (s String) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s String) String() string { _ = "STUB: not implemented"; return "" }

func (s String) GoString() string { _ = "STUB: not implemented"; return "" }

func (s String) MarshalBinary() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

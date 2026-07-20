package consumererror

type permanent struct {
	err error
}

func NewPermanent(err error) error { _ = "STUB: not implemented"; return nil }

func (p permanent) Error() string { _ = "STUB: not implemented"; return "" }

func (p permanent) Unwrap() error { _ = "STUB: not implemented"; return nil }

func IsPermanent(err error) bool { _ = "STUB: not implemented"; return false }

package consumererror

type downstreamError struct {
	inner error
}

var _ error = downstreamError{}

func (de downstreamError) Error() string { _ = "STUB: not implemented"; return "" }

func (de downstreamError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func NewDownstream(err error) error { _ = "STUB: not implemented"; return nil }

func IsDownstream(err error) bool { _ = "STUB: not implemented"; return false }

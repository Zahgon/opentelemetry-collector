package consumererror

import (
	"google.golang.org/grpc/status"
)

type Error struct {
	error
	httpStatus  int
	grpcStatus  *status.Status
	isRetryable bool
}

var _ error = (*Error)(nil)

func NewOTLPHTTPError(origErr error, httpStatus int) error { _ = "STUB: not implemented"; return nil }

func NewOTLPGRPCError(origErr error, status *status.Status) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRetryableError(origErr error) error { _ = "STUB: not implemented"; return nil }

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *Error) IsRetryable() bool { _ = "STUB: not implemented"; return false }

func ToHTTPStatus(err error) int { _ = "STUB: not implemented"; return 0 }

func ToGRPCStatus(err error) *status.Status { _ = "STUB: not implemented"; return nil }

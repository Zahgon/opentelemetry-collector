package request

import (
	"errors"
)

const (
	protoTag1TypeByte = 0x0D

	requestFormatVersion = uint32(1)
)

var ErrInvalidFormat = errors.New("invalid request payload format")

func isRequestPayloadV1(data []byte) bool { _ = "STUB: not implemented"; return false }

package pcommon

import (
	"time"
)

type Timestamp uint64

func NewTimestampFromTime(t time.Time) Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (ts Timestamp) AsTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (ts Timestamp) String() string { _ = "STUB: not implemented"; return "" }

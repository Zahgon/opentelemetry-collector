package entity

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func ResourceEntityRefs(res pcommon.Resource) EntityRefSlice {
	_ = "STUB: not implemented"
	return *new(EntityRefSlice)
}

func ResourceEntities(res pcommon.Resource) EntityMap {
	_ = "STUB: not implemented"
	return *new(EntityMap)
}

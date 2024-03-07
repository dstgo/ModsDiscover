package dict

import (
	"github.com/dstgo/tracker/internal/types/helper"
	"github.com/dstgo/tracker/pkg/vax"
)

var (
	RuleDictName  = helper.Rules(vax.RangeLenRune(1, 30))
	RuleDictCode  = helper.Rules(vax.RangeLenRune(1, 30))
	RuleDictKey   = helper.Rules(vax.RangeLenRune(1, 30))
	RuleDictValue = helper.Rules(vax.RangeLenRune(1, 200))

	RuleDictDataType = helper.Rules(vax.In(StringType, Int64Type, Float64Type, BoolType))
)

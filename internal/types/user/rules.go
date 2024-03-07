package user

import (
	"github.com/dstgo/tracker/internal/types/helper"
	"github.com/dstgo/tracker/pkg/vax"
	"github.com/dstgo/tracker/pkg/vax/is"
)

var (
	RuleUsername = helper.Rules(is.Alphanumeric, vax.RangeLenRune(6, 20))

	RulePassword = helper.Rules(vax.RangeLenRune(10, 30))
)

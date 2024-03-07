package email

import (
	"github.com/dstgo/tracker/internal/types/helper"
	"github.com/dstgo/tracker/pkg/vax"
	"github.com/dstgo/tracker/pkg/vax/is"
)

var (
	RuleEmail = helper.Rules(is.Email)

	RuleEmailCode = helper.Rules(is.Alphanumeric, vax.EqLength(8, false))
)

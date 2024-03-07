package system

import (
	"github.com/dstgo/tracker/internal/types/helper"
	"github.com/dstgo/tracker/pkg/vax"
)

var (
	RulePing = helper.Rules(vax.RangeLenRune(1, 10), vax.In("wilson", "wendy"))
)

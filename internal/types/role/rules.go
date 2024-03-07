package role

import (
	"github.com/dstgo/tracker/internal/types/helper"
	"github.com/dstgo/tracker/pkg/vax"
	"github.com/dstgo/tracker/pkg/vax/is"
	"regexp"
)

var (
	PermObjectRule = helper.RegexRule{Regx: regexp.MustCompilePOSIX("^[a-zA-Z0-9/_-]*$"), I18n: "rule.perm.object"}
)

var (
	RulePermObj   = helper.Rules(vax.Match(PermObjectRule.Regx).Code(PermObjectRule.I18n), vax.RangeLenRune(1, 200))
	RulePermAct   = helper.Rules(vax.RangeLenRune(1, 60), is.Alpha)
	RulePermTag   = helper.Rules(vax.RangeLenRune(1, 20), is.Alphanumeric)
	RulePermGroup = helper.Rules(vax.RangeLenRune(1, 60))
	RuleRoleName  = helper.Rules(vax.RangeLenRune(1, 60))
	RuleRoleCode  = helper.Rules(vax.RangeLenRune(1, 30), is.Alphanumeric)
)

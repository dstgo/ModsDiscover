package middleware

import (
	"github.com/dstgo/tracker/internal/core/resp"
	"github.com/dstgo/tracker/pkg/ginx/bind"
	"github.com/dstgo/tracker/pkg/vax"
	"github.com/gin-gonic/gin"
)

func BindBadParamsHandler() bind.Handler {
	return func(ctx *gin.Context, bindErr error) {
		if bindErr != nil {
			switch bindErr.(type) {
			// validate internal occur error
			case vax.InternalError:
				resp.InternalFailed(ctx).MsgI18n("err.program").Error(bindErr).Send()
			default:
				resp.Fail(ctx).MsgI18n("err.badparams").Error(bindErr).Send()
			}
		}
	}
}

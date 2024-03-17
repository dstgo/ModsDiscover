package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dstgo/tracker/internal/handler"
	"github.com/dstgo/tracker/internal/types"
	"github.com/dstgo/tracker/pkg/resp"
)

type ModAPI struct {
	modHandler handler.ModHandler
}

// Search returns a list of dst workshop files
func (mod ModAPI) Search(c context.Context, ctx *app.RequestContext) {
	var queryOption types.SearchModsOption
	if err := ctx.BindAndValidate(&queryOption); err != nil {
		resp.Failed(ctx).Error(err).Do()
		return
	}

	list, err := mod.modHandler.SearchModList(c, queryOption)
	if err != nil {
		resp.Failed(ctx).Error(err).Do()
	} else {
		resp.Ok(ctx).Data(list.Response).Do()
	}
}

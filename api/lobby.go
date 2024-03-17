package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dstgo/tracker/internal/handler"
	"github.com/dstgo/tracker/internal/types"
	"github.com/dstgo/tracker/pkg/resp"
	"time"
)

type LobbyAPI struct {
	LobbyHandler handler.LobbyHandler
}

// List [GET] /lobby/list?
// returns a list of lobby servers
func (l *LobbyAPI) List(c context.Context, ctx *app.RequestContext) {
	var listOptions types.QueryLobbyServersOptions
	if err := ctx.BindAndValidate(&listOptions); err != nil {
		resp.Failed(ctx).Error(err).Do()
		return
	}

	pageResult, err := l.LobbyHandler.GetServersByPage(c, listOptions)
	if err != nil {
		resp.Failed(ctx).Error(err).Do()
	} else {
		resp.Ok(ctx).Data(pageResult).Do()
	}
}

// Details [GET] /lobby/details?rowId=KU_X19asjdla&region=ap-east-1
// returns details info for specific lobby server
func (l *LobbyAPI) Details(c context.Context, ctx *app.RequestContext) {
	var detailsOptions types.QueryLobbyServerDetailsOption
	if err := ctx.BindAndValidate(&detailsOptions); err != nil {
		resp.Failed(ctx).Error(err).Do()
		return
	}

	details, err := l.LobbyHandler.GetServerDetails(c, detailsOptions.Region, detailsOptions.RowId)
	if err != nil {
		resp.Failed(ctx).Error(err).Do()
	} else {
		resp.Ok(ctx).Data(details).Do()
	}
}

// Statistic [GET] /lobby/stat?before=xx&until=xx
// returns statistics information for dst lobby
func (l *LobbyAPI) Statistic(c context.Context, ctx *app.RequestContext) {
	var opt types.QueryLobbyStatisticOption
	if err := ctx.BindAndValidate(&opt); err != nil {
		resp.Failed(ctx).Error(err).Do()
		return
	}

	duration, err := time.ParseDuration(opt.Duration)
	if err != nil {
		resp.Failed(ctx).Error(err).Do()
		return
	}

	statisticInfo, err := l.LobbyHandler.GetStatisticInfo(c, opt.Before, opt.Until, opt.Tail, duration)
	if err != nil {
		resp.Failed(ctx).Error(err).Do()
	} else {
		resp.Ok(ctx).Data(statisticInfo).Do()
	}
}

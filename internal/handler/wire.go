//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/dstgo/tracker/internal/conf"
	"github.com/dstgo/tracker/internal/data"
	"github.com/dstgo/tracker/pkg/ginx"
	"github.com/google/wire"
)

//go:generate wire gen
func setupHandlerRouter(appConf *conf.AppConf, router *ginx.RouterGroup, datasource *data.DataSource) (Router, func(), error) {
	panic(wire.Build(HandlerProviderSet))
}

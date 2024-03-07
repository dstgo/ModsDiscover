package api

import (
	_ "github.com/dstgo/tracker/internal/api/docs"
	"github.com/dstgo/tracker/internal/api/dst"
	"github.com/dstgo/tracker/internal/api/lobby"
	"github.com/dstgo/tracker/internal/api/mod"
	"github.com/dstgo/tracker/internal/api/user"
	"github.com/dstgo/tracker/internal/conf"
	"github.com/dstgo/tracker/internal/core/authen"
	"github.com/dstgo/tracker/internal/core/log"
	"github.com/dstgo/tracker/internal/core/role"
	"github.com/dstgo/tracker/internal/data"
	"github.com/dstgo/tracker/internal/data/cache"
	"github.com/dstgo/tracker/internal/handler/middleware"
	"github.com/dstgo/tracker/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"path"
)

const (
	BasePath = "/open/v1"
	DocPath  = "/open/v1/doc"
)

var ApiProviderSet = wire.NewSet(
	user.UserRouterSet,
	wire.Struct(new(Router), "*"),
)

type Router struct {
	User   user.APIRouter
	Lobby  lobby.APIRouter
	Mod    mod.APIRouter
	Server dst.APIRouter
}

// SetupOpenAPI initializes app open api router configuration
func SetupOpenAPI(cfg *conf.AppConf, engine *gin.Engine, datasource *data.DataSource) (Router, error) {
	if !cfg.ServerConf.OpenAPI {
		return Router{}, nil
	}

	var (
		roleResolver = role.NewGormResolver(datasource.ORM())
		apikeyCache  = cache.NewAPIKeyCache(datasource.Redis())
		keyAuthor    = authen.NewAPIKeyCacheAuthor(datasource, roleResolver, apikeyCache)
	)

	openapiRouter := ginx.NewRouterGroup(engine.RouterGroup.Group(BasePath))

	openapiRouter.Attach(
		middleware.UseOpenAPIAuth(keyAuthor),
	)

	if cfg.ServerConf.Swagger {
		engine.GET(path.Join(DocPath, "*any"), ginSwagger.CustomWrapHandler(Config, swaggerFiles.NewHandler()))
		log.L().Infof("visit OpenAPI Doc on http://%s%s", cfg.ServerConf.HttpConf.Address, path.Join(DocPath, "index.html"))
	}
	router := setupOpenAPIRouter(openapiRouter, datasource)

	err := initApiRouterACL(openapiRouter, roleResolver)
	if err != nil {
		return router, err
	}

	return router, nil
}

var Config = &ginSwagger.Config{
	URL:                      "doc.json",
	DocExpansion:             "list",
	InstanceName:             "openapi",
	Title:                    "OpenAPI",
	DefaultModelsExpandDepth: 0,
	DeepLinking:              true,
	PersistAuthorization:     false,
	Oauth2DefaultClientID:    "",
}

// swagger declarative api comment

// @title		                    App Open API Documentation
// @version		                    v1.0.0
// @description                     open api documentation, to access these open api, you need to add apikey in query param named "key"
// @contact.name                    dstgo
// @contact.url                     https://github.com/dstgo
// @BasePath                        /open/v1
// @license.name                    MIT LICENSE
// @license.url                     https://mit-license.org/
// @securityDefinitions.apikey      ApiKeyAuth
// @in                              query
// @name                            key
//go:generate swag init --generatedTime --instanceName openapi -g api.go -d ./,../types,../core/resp --output ./docs && swag fmt -g api.go -d ./

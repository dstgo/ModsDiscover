package email

import (
	"github.com/dstgo/tracker/internal/types"
	"github.com/dstgo/tracker/internal/types/meta"
	"github.com/dstgo/tracker/pkg/ginx"
	"github.com/google/wire"
)

// HandlerRouter just for wire injection, no real influence
type HandlerRouter types.NopType

var EmailRouterSet = wire.NewSet(
	EmailProviderSet,
	wire.Struct(new(Handler), "*"),
	SetupRouter,
)

type Handler struct {
	Email EmailHandler
}

func SetupRouter(api *ginx.RouterGroup, Handler Handler) HandlerRouter {
	emailGroup := api.Group("email", ginx.M(meta.Group("route.email.group")))
	{
		emailGroup.GET("code", ginx.M(meta.NoAuth, meta.Name("route.email.code")), Handler.Email.SendCodeEmail)
	}
	return types.NopObj
}

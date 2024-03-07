package mod

import (
	"github.com/dstgo/steamapi"
	"github.com/dstgo/tracker/internal/types/dst"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ModApiProviderSet = wire.NewSet()

func NewHandler(steamClient *steamapi.Client) *ModsHandler {
	return &ModsHandler{}
}

type ModsLogic interface {
	Search(searchOption dst.SearchModsOption) (dst.SearchModsResult, error)
	Details() (dst.ModsDetail, error)
	Options() (any, error)
	Version() string
}

type ModsHandler struct {
	mods ModsLogic
}

// Search
// @Summary      Search
// @Description  returns result of workshop item list
// @Tags         mods
// @Accept       json
// @Produce      json
// @Param        SearchModsOption   body  types.SearchModsOption  true  "SearchModsOption"
// @Success      200  {object}  types.Response{data=types.SearchModsResult}
// @Router       /mods/search [GET]
func (m *ModsHandler) Search(ctx *gin.Context) {

}

// Details
// @Summary      Details
// @Description  returns the details for the specified work shop item
// @Tags         mods
// @Accept       json
// @Produce      json
// @Param        QueryModOption   query     types.QueryModOption true  "QueryModOption"
// @Success      200  {object}  types.Response{data=types.ModsDetail}
// @Router       /mods/details [GET]
func (m *ModsHandler) Details(ctx *gin.Context) {

}

// Options
// @Summary      Options
// @Description  returns the configuration options for the specified work shop item
// @Tags         mods
// @Accept       json
// @Produce      json
// @Param        QueryModOption   query     types.QueryModOption true  "QueryModOption"
// @Success      200  {object}  types.Response{}
// @Router       /api [GET]
func (m *ModsHandler) Options(ctx *gin.Context) {

}

// Version
// @Summary      Version
// @Description  returns the latest version for the specified work shop item
// @Tags         mods
// @Accept       json
// @Produce      json
// @Param        id   query     string  true  "work shop id"
// @Success      200  {object}  types.Response
// @Router       /api [GET]
func (m *ModsHandler) Version(ctx *gin.Context) {

}

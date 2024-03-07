package lobby

import (
	lobby "github.com/dstgo/lobbyapi"
	"github.com/dstgo/tracker/internal/data"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var LobbyApiProvider = wire.NewSet()

func NewHandler(lobbyClient *lobby.Client, ds *data.DataSource) *QueryAPI {
	return &QueryAPI{}
}

type QueryAPI struct {
}

// GetServerList
// @Summary      GetServerList
// @Description  returns a list of rooms in lobby
// @Tags         lobby
// @Accept       json
// @Produce      json
// @Param        QueryLobbyServersOptions  path   types.QueryLobbyServersOptions  true  "QueryLobbyServersOptions"
// @Success      200  {object}  types.Response
// @Router       /lobby/servers [GET]
func (l *QueryAPI) GetServerList(ctx *gin.Context) {

}

// GetServerDetails
// @Summary      GetServerDetails
// @Description  returns the details for specific room in from lobby
// @Tags         lobby
// @Accept       json
// @Produce      json
// @Param        rowId   path   string  true  "rowId"
// @Success      200  {object}  types.Response
// @Router       /lobby/details [GET]
func (l *QueryAPI) GetServerDetails(ctx *gin.Context) {

}

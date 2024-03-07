package dst

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ServerApiProviderSet = wire.NewSet()

type ServerHandler struct {
}

// GetDstVersion
// @Summary      GetDstVersion
// @Description  returns the latest version of the dst server
// @Tags         dst
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response{data=string}
// @Router       /dst/version [GET]
func (s *ServerHandler) GetDstVersion(ctx *gin.Context) {

}

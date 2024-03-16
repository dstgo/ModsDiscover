package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"time"
)

type SystemAPI struct {
}

// Ts [GET] /ts
// returns the timestamp of now
func (s *SystemAPI) Ts(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, map[string]any{
		"ts": time.Now().UnixNano(),
	})
}

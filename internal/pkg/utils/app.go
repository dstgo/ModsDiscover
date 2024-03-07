package utils

import (
	"github.com/dstgo/tracker/internal/core/log"
	"github.com/dstgo/tracker/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func IsDebugMode() bool {
	return log.L().Level >= logrus.DebugLevel || gin.Mode() == gin.DebugMode
}

// PrintRouters
// use for debugging, print all the route has benn register
func PrintRouters(root *ginx.RouterGroup, printGroup bool) error {
	return root.Walk(func(info ginx.WalkRouteInfo) error {
		if !printGroup && info.IsGroup {
			return nil
		}
		log.L().Debugf("Method:%s\tPath:%-20s\tMeta:%+v", info.Method, info.FullPath, info.Meta)
		return nil
	})
}

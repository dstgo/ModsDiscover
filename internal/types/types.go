package types

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dstgo/steamapi"
	"github.com/dstgo/tracker/pkg/lobbyapi"
	"github.com/oschwald/geoip2-golang"
	"github.com/qiniu/qmgo"
	"time"
)

var TimeZone *time.Location

func init() {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	TimeZone = location
}

type PageResult[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}

type Env struct {
	MongoDB  *qmgo.QmgoClient
	LobbyCLI *lobbyapi.Client
	SteamCLI *steamapi.Client
	GeoIpDB  *geoip2.Reader
	Logger   hlog.FullLogger
}

const (
	DstAppID = 322330
)

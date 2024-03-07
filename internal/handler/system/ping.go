package system

import (
	"fmt"
	"github.com/dstgo/tracker/internal/conf"
	"github.com/dstgo/tracker/internal/types"
	"github.com/dstgo/tracker/internal/types/system"
	"time"
)

type PingApp struct {
	conf *conf.AppConf
}

func NewPingLogic(conf *conf.AppConf) PingApp {
	return PingApp{
		conf: conf,
	}
}

func (p PingApp) Ping(name string) system.PingReply {
	return system.PingReply{Reply: fmt.Sprintf("hello %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}

func (p PingApp) Pong(name string) system.PingReply {
	return system.PingReply{Reply: fmt.Sprintf("goodbye %s! Now is %s.", name, time.Now().Format(types.DateTimeFormat))}
}

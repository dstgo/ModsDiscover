package jobs

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/internal/handler"
	"github.com/robfig/cron/v3"
	"time"
)

// LobbyCollector collects server information from klei lobby
type LobbyCollector struct {
	conf    conf.LobbyConf
	handler handler.LobbyHandler
	cron    *cron.Cron
}

// Collect collects server information from klei lobby
func (l LobbyCollector) Collect() {
	start := time.Now()
	// max cost time duration
	ctx, cancelFunc := context.WithTimeout(context.Background(), l.conf.Timeout)
	defer cancelFunc()

	// collect
	collected, err := l.handler.SyncLocalServers(ctx, 20)
	if err != nil {
		hlog.Errorf("LOBBY_COLLECTOR: error=%v", err)
		return
	}

	// log events
	cost := time.Now().Sub(start).String()
	hlog.Infof("LOBBY_COLLECTOR: cost=%s collected=%d", cost, collected)
}

// Clear clears expired data
func (l LobbyCollector) Clear() {
	start := time.Now()

	// clear expired
	deleted, total, err := l.handler.ClearExpiredServers(context.Background(), l.conf.TTL)
	if err != nil {
		hlog.Errorf("LOBBY_COLLECTOR: error=%v", err)
		return
	}

	cost := time.Now().Sub(start).String()
	hlog.Infof("LOBBY_COLLECTOR: cost=%s deleted=%d remained=%d", cost, deleted, total)
}

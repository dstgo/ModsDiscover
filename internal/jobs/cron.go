package jobs

import (
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/internal/handler"
	"github.com/dstgo/tracker/internal/types"
	"github.com/robfig/cron/v3"
	"log/slog"
)

// cronLogger adapter
type cronLogger struct {
	logger *slog.Logger
	prefab string
}

func (c cronLogger) Info(msg string, keysAndValues ...interface{}) {
	c.logger.Info(c.prefab+": "+msg, keysAndValues...)
}

func (c cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	c.logger.Error(c.prefab+": "+msg, append([]any{"err", err}, keysAndValues...)...)
}

func LoadCronJobs(dstConf conf.DstConf, lobbyHandler handler.LobbyHandler) (*cron.Cron, error) {
	cronJob := cron.New(
		cron.WithLogger(cronLogger{logger: slog.Default(), prefab: "CRON"}),
		cron.WithLocation(types.TimeZone),
	)

	// lobby collector
	lobbyCollector := LobbyCollector{dstConf.Lobby, lobbyHandler, cronJob}
	if _, err := cronJob.AddFunc(dstConf.Lobby.CollectCron, lobbyCollector.Collect); err != nil {
		return nil, err
	}
	if _, err := cronJob.AddFunc(dstConf.Lobby.ClearCron, lobbyCollector.Clear); err != nil {
		return nil, err
	}

	return cronJob, nil
}

package server

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dstgo/tracker/conf"
	hertzslog "github.com/hertz-contrib/logger/slog"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

// NewLogger use hertz FullLogger adapter
func NewLogger(logConf conf.LogConf) (hlog.FullLogger, io.Closer, error) {
	dir := filepath.Dir(logConf.File)
	if dir != "." && len(dir) > 0 {
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			return nil, nil, err
		}
	}

	// open log file
	logFile, err := os.OpenFile(logConf.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}

	writer := io.MultiWriter(os.Stdout, logFile)
	// set level
	var level = new(slog.LevelVar)
	if err := level.UnmarshalText([]byte(logConf.Level)); err != nil {
		return nil, nil, err
	}

	logger := hertzslog.NewLogger(
		hertzslog.WithLevel(level),
		hertzslog.WithOutput(writer),
	)

	slog.SetDefault(logger.Logger())

	// set logger
	hlog.SetLogger(logger)

	return logger, logFile, nil
}

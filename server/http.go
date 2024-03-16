package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/dstgo/tracker/conf"
	"github.com/dstgo/tracker/internal/types"
	"github.com/dstgo/tracker/pkg/resp"
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/requestid"
	"net/http"
)

// returns a new hertz http server
func newHttpServer(httpConf conf.HttpConf) (*server.Hertz, error) {
	hertz := server.New(
		server.WithHostPorts(httpConf.Listen),
		server.WithReadTimeout(httpConf.ReadTimeout),
		server.WithIdleTimeout(httpConf.IdleTimeout),
		server.WithBasePath(httpConf.BasePath),
	)

	logHandler, err := accessLogHandler()
	if err != nil {
		return nil, err
	}

	hertz.Use(
		// recovery handler
		recoveryHandler(),
		// request limiter
		limiterHandler(),
		// X-Request-ID
		requestid.New(),
		// log handler
		logHandler,
		// cache handler
		cacheHandler(httpConf),
	)

	return hertz, nil
}

func accessLogHandler() (app.HandlerFunc, error) {
	accesslog.Tags["requestId"] = func(output accesslog.Buffer, c *app.RequestContext, data *accesslog.Data, extraParam string) (int, error) {
		requestId := c.Response.Header.Get("X-Request-ID")
		return output.WriteString(requestId)
	}

	// format string
	format := "${status}   -${latency} ${method} ${url} ${ip} ${queryParams} ${resHeaders}"

	return accesslog.New(
		accesslog.WithTimeZoneLocation(types.TimeZone),
		accesslog.WithFormat(format),
	), nil
}

func recoveryHandler() app.HandlerFunc {
	return recovery.Recovery(
		recovery.WithRecoveryHandler(func(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
			hlog.DefaultLogger().CtxErrorf(c, "[Recovery] err=%v\nstack=%s", err, stack)
			hlog.DefaultLogger().Infof("Client: %s", ctx.Request.Header.UserAgent())
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}),
	)
}

func cacheHandler(httpConf conf.HttpConf) app.HandlerFunc {
	store := persist.NewMemoryStore(httpConf.CacheTTL)
	cacheH := cache.NewCacheByRequestURIWithIgnoreQueryOrder(store, httpConf.CacheTTL, cache.WithPrefixKey("tracker-cache-"))
	return cacheH
}

func limiterHandler() app.HandlerFunc {
	limiter := bbr.NewLimiter()
	return func(c context.Context, ctx *app.RequestContext) {
		done, err := limiter.Allow()
		if err != nil {
			resp.New(ctx).Status(consts.StatusTooManyRequests).Msg("too many requests").Do()
			ctx.Abort()
		} else {
			ctx.Next(c)
			done(ratelimit.DoneInfo{})
		}
	}
}

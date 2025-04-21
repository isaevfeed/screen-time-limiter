package app

import (
	"context"
	"fmt"
	"net/http"
	"screen-time-limiter/cmd/screen-time-limiter/service_provider"
	"screen-time-limiter/internal/app/handlers"
	"screen-time-limiter/internal/config"
	"screen-time-limiter/internal/utils/timer"
)

type (
	App struct {
		config *config.Config
		server *http.Server

		sp *service_provider.Provider
	}
)

func Init(ctx context.Context, config *config.Config, sp *service_provider.Provider) *App {
	app := &App{
		config: config,
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%s", config.Service.Host, config.Service.Port),
		},
	}

	logMid := sp.GetLogMiddleware().ServeHTTP(bootstrapHandlers(ctx, sp))
	app.server.Handler = logMid

	return app
}

func (a *App) ListenAndServe() error {
	return a.server.ListenAndServe()
}

func bootstrapHandlers(ctx context.Context, sp *service_provider.Provider) http.Handler {
	mx := http.NewServeMux()

	mx.Handle("POST /v1/user", handlers.NewAddUserHandler(sp.GetUserRepo(ctx)))
	mx.Handle("POST /v1/limit", handlers.NewAddTimeLimitHandler(sp.GetLimitRepo(ctx)))
	mx.Handle("POST /v1/limit/history", handlers.NewPushCurrentTimeHandler(sp.GetLimitHistoryRepo(ctx), sp.GetLimitRepo(ctx), timer.New()))
	mx.Handle("GET /v1/limit/history/total", handlers.NewGetTotalHistoryTimesHandler(sp.GetLimitHistoryRepo(ctx), sp.GetLimitRepo(ctx), timer.New()))

	return mx
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

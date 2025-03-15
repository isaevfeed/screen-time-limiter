package app

import (
	"context"
	"fmt"
	"net/http"
	"screen-time-limiter/internal/app/handlers"
	"screen-time-limiter/internal/config"
)

type (
	App struct {
		config *config.Config
		server *http.Server
	}
)

func Init(config *config.Config) *App {
	app := &App{
		config: config,
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%s", config.Service.Host, config.Service.Port),
		},
	}

	app.server.Handler = bootstrapHandlers()

	return app
}

func (a *App) ListenAndServe() error {
	return a.server.ListenAndServe()
}

func bootstrapHandlers() http.Handler {
	mx := http.NewServeMux()

	mx.Handle("POST /v1/user", handlers.NewAddUserHandler())

	return mx
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

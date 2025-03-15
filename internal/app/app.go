package app

import (
	"context"
	"fmt"
	"net/http"
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

	return app
}

func (a *App) ListenAndServe() error {
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"screen-time-limiter/cmd/screen-time-limiter/service_provider"
	"screen-time-limiter/internal/app"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	p := service_provider.New()
	ctx := context.Background()
	cfg := p.MustGetConfig()
	log := p.GetLogger(cfg.Service.Env)
	appInit := app.Init(
		ctx,
		cfg,
		p,
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info(fmt.Sprintf("Server is started on %s:%s", cfg.Service.Host, cfg.Service.Port))

		if err := appInit.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("app.ListenAndServe: ", "err", err)
		}
	}()

	<-stop

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, cfg.Service.Timeout)
	defer cancel()

	if err := appInit.Shutdown(ctx); err != nil {
		log.Error("app.Shutdown: ", "err", err)
	}

	log.Info("Server stopped")
}

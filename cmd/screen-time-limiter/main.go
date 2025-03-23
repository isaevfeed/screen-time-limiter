package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"screen-time-limiter/cmd/screen-time-limiter/service_provider"
	"screen-time-limiter/internal/app"
	"screen-time-limiter/internal/config"
	"syscall"
)

func main() {
	cfg, err := config.Load(os.Getenv("CONFIG_FILE"))
	if err != nil {
		panic(err)
	}
	p := service_provider.New()
	log := p.GetLogger(cfg.Service.Env)
	app := app.Init(
		cfg,
		p.GetLogMiddleware(cfg.Service.Env),
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info(fmt.Sprintf("Server is started on %s:%s", cfg.Service.Host, cfg.Service.Port))

		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("app.ListenAndServe: ", err)
		}
	}()

	<-stop

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Service.Timeout)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Error("app.Shutdown: ", err)
	}

	log.Info("Server stopped")
}

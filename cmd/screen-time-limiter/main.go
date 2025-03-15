package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"screen-time-limiter/internal/app"
	"screen-time-limiter/internal/config"
	"syscall"
)

const (
	prod = "prod"
)

func main() {
	cfg, err := config.Load(os.Getenv("CONFIG_FILE"))
	if err != nil {
		panic(err)
	}
	app := app.Init(cfg)
	log := initLogger(cfg.Service.Env)

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

func initLogger(env string) *slog.Logger {
	switch env {
	case prod:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
}

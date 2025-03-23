package service_provider

import (
	"log/slog"
	"os"
)

const (
	prod = "prod"
)

func (p *Provider) GetLogger(env string) *slog.Logger {
	if p.logger == nil {
		p.logger = initLogger(env)
	}

	return p.logger
}

func initLogger(env string) *slog.Logger {
	switch env {
	case prod:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
}

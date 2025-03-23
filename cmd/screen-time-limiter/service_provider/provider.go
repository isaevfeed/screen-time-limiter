package service_provider

import (
	"log/slog"
	"screen-time-limiter/internal/app/middlewares"
)

type Provider struct {
	logger        *slog.Logger
	logMiddleware *middlewares.LogQueries
}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) GetLogMiddleware(env string) *middlewares.LogQueries {
	if p.logMiddleware == nil {
		p.logMiddleware = middlewares.NewLogQueries(
			p.GetLogger(env),
		)
	}

	return p.logMiddleware
}

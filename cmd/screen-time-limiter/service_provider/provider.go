package service_provider

import (
	"context"
	"log/slog"
	"screen-time-limiter/internal/app/middlewares"
	"screen-time-limiter/internal/config"
	"screen-time-limiter/internal/repos/limit"
	"screen-time-limiter/internal/repos/limit_history"
	"screen-time-limiter/internal/repos/user"
)

type Provider struct {
	cfg           *config.Config
	logger        *slog.Logger
	logMiddleware *middlewares.LogQueries

	userRepo         *user.Repo
	limitRepo        *limit.Repo
	limitHistoryRepo *limit_history.Repo
}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) GetLogMiddleware() *middlewares.LogQueries {
	if p.logMiddleware == nil {
		p.logMiddleware = middlewares.NewLogQueries(
			p.GetLogger(p.MustGetConfig().Service.Env),
		)
	}

	return p.logMiddleware
}

func (p *Provider) GetUserRepo(ctx context.Context) *user.Repo {
	if p.userRepo == nil {
		p.userRepo = user.NewRepo(mustInitDatabase(ctx, p.MustGetConfig()))
	}

	return p.userRepo
}

func (p *Provider) GetLimitRepo(ctx context.Context) *limit.Repo {
	if p.limitRepo == nil {
		p.limitRepo = limit.NewRepo(mustInitDatabase(ctx, p.MustGetConfig()))
	}

	return p.limitRepo
}

func (p *Provider) GetLimitHistoryRepo(ctx context.Context) *limit_history.Repo {
	if p.limitHistoryRepo == nil {
		p.limitHistoryRepo = limit_history.NewRepo(mustInitDatabase(ctx, p.MustGetConfig()))
	}

	return p.limitHistoryRepo
}

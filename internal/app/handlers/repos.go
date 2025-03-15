package handlers

import (
	"context"
	"screen-time-limiter/internal/domain/model"
	"time"
)

type (
	limitRepo interface {
		Create(ctx context.Context, limit model.Limit) error
		GetByLimitID(ctx context.Context, limitID int32) (*model.Limit, error)
	}

	userRepo interface {
		Add(ctx context.Context, user model.User) error
	}

	limitHistoryRepo interface {
		Push(ctx context.Context, history model.LimitHistory) error
		Sum(ctx context.Context, limitID int32, date time.Time) (int, error)
	}
)

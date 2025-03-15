package user

import (
	"context"
	"screen-time-limiter/internal/domain/model"
)

type (
	UseCase struct {
		repo repo
	}

	repo interface {
		Add(ctx context.Context, user model.User) error
		List(ctx context.Context) ([]model.User, error)
	}
)

func NewUseCase(repo repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

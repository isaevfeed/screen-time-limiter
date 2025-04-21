package user

import (
	"context"
	"screen-time-limiter/internal/domain/model"
)

func (u *UseCase) Add(ctx context.Context, user model.User) error {
	return u.repo.Add(ctx, user)
}

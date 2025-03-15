package user

import (
	"context"
	"fmt"
	"screen-time-limiter/internal/domain/model"
)

func (r *Repo) Add(ctx context.Context, user model.User) error {
	query, args, err := psql.Insert(tableName).
		Columns("first_name", "last_name").
		Values(user.FirstName, user.LastName).
		Suffix("ON CONFLICT DO NOTHING").
		ToSql()
	if err != nil {
		return fmt.Errorf("psql.Insert: %w", err)
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("pool.Exec: %w", err)
	}

	return nil
}

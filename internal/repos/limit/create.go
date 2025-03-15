package limit

import (
	"context"
	"fmt"
	"screen-time-limiter/internal/domain/model"
)

func (r *Repo) Create(ctx context.Context, limit model.Limit) error {
	query, args, err := psql.Insert(tableName).Columns("user_id", "time_limit").
		Values(limit.UserID, limit.Amount).ToSql()
	if err != nil {
		return fmt.Errorf("psql.Insert: %w", err)
	}
	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("pool.Exec: %w", err)
	}

	return nil
}

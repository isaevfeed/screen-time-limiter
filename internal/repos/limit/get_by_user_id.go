package limit

import (
	"context"
	"fmt"
	"screen-time-limiter/internal/domain/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repo) GetByUserID(ctx context.Context, userID int32) (model.Limit, error) {
	query, args, err := psql.Select("id, user_id, time_limit").From(tableName).
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return model.Limit{}, fmt.Errorf("sq.Select: %w", err)
	}

	var limit dao
	_ = r.pool.QueryRow(ctx, query, args...).
		Scan(&limit.ID, &limit.UserID, &limit.Amount)

	return model.Limit{
		ID:     limit.ID,
		UserID: limit.UserID,
		Amount: limit.Amount,
	}, nil
}

package limit_history

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *Repo) Sum(ctx context.Context, limitID int32, date time.Time) (int, error) {
	var limitSum int
	err := pgx.BeginTxFunc(ctx, r.pool, pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	}, func(tx pgx.Tx) error {
		query, args, err := psql.Select("SUM(time_amount)").From(tableName).
			Where(squirrel.Eq{"limit_id": limitID}).
			Where(squirrel.Eq{"limit_date": date}).
			GroupBy("limit_date").
			ToSql()
		if err != nil {
			return fmt.Errorf("squirrel.Select: %w", err)
		}

		_ = tx.QueryRow(ctx, query, args...).Scan(&limitSum)

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("pgx.BeginTxFunc: %w", err)
	}

	return limitSum, nil
}

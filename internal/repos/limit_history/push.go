package limit_history

import (
	"context"
	"fmt"
	"screen-time-limiter/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) Push(ctx context.Context, history model.LimitHistory) error {
	err := pgx.BeginTxFunc(ctx, r.pool, pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	}, func(tx pgx.Tx) error {
		year, month, day := history.LimitDate.Date()

		query, args, err := psql.Insert(tableName).
			Columns("limit_id", "time_amount", "sent_at", "limit_date").
			Values(history.LimitID, history.TimeAmount, history.SentAt, fmt.Sprintf("%d-%d-%d", year, month, day)).
			ToSql()
		if err != nil {
			return fmt.Errorf("squirrel.Insert: %w", err)
		}

		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("tx.Exec: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("pgx.BeginTxFunc: %w", err)
	}

	return nil
}

package user

import (
	"context"
	"fmt"
	"screen-time-limiter/internal/domain/model"
)

func (r *Repo) List(ctx context.Context) ([]model.User, error) {
	query, _, err := psql.Select("id, first_name, last_name").From(tableName).ToSql()
	if err != nil {
		return nil, fmt.Errorf("psql.Select: %w", err)
	}

	var rowsDao []dao
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("db.Select: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var row dao

		if err = rows.Scan(&row.ID, &row.FirstName, &row.LastName); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		rowsDao = append(rowsDao, row)
	}

	users := make([]model.User, len(rowsDao))
	for i, row := range rowsDao {
		users[i] = model.User{
			ID:        row.ID,
			FirstName: row.FirstName,
			LastName:  row.LastName,
		}
	}

	return users, nil
}

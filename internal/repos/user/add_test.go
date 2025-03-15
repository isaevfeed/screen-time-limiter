package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"screen-time-limiter/internal/domain/model"
	"testing"
)

func TestRepo_Add(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	a := assert.New(t)
	repo, db := makeRepo(ctx)

	tx, err := db.Begin(ctx)
	defer tx.Rollback(ctx)

	a.NoError(err)

	user := model.User{
		FirstName: "Михаил",
		LastName:  "Исаев",
	}
	err = repo.Add(context.Background(), user)
	a.NoError(err)

	type dao struct {
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
	}

	var row dao
	_ = db.QueryRow(ctx, "SELECT first_name, last_name FROM users WHERE first_name=$1 AND last_name=$2", user.FirstName, user.LastName).
		Scan(&row.FirstName, &row.LastName)

	a.NotNil(row)
	a.Equal(row.FirstName, user.FirstName)
	a.Equal(row.LastName, user.LastName)

	err = tx.Rollback(ctx)
	a.NoError(err)
}

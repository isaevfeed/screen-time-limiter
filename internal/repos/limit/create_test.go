package limit

import (
	"context"
	"math"
	"screen-time-limiter/internal/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepo_Create(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	a := assert.New(t)
	repo, db := makeRepo(ctx)

	tx, err := db.Begin(ctx)
	defer tx.Rollback(ctx)

	a.NoError(err)

	limitAmount := time.Hour * 2
	limit := model.Limit{
		UserID: 1,
		Amount: int32(math.Round(limitAmount.Seconds())),
	}
	err = repo.Create(context.Background(), limit)
	a.NoError(err)

	var row dao
	_ = db.QueryRow(ctx, "SELECT user_id, time_limit FROM limits WHERE user_id=$1", limit.UserID).
		Scan(&row.UserID, &row.Amount)

	a.NotNil(row)
	a.Equal(row.Amount, limit.Amount)

	err = tx.Rollback(ctx)
	a.NoError(err)
}

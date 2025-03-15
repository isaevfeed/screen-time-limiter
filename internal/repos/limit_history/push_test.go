package limit_history

import (
	"context"
	"fmt"
	"screen-time-limiter/internal/domain/model"
	"screen-time-limiter/internal/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepo_Push(t *testing.T) {
	t.Parallel()

	a := assert.New(t)
	ctx := context.Background()
	repo, db := makeRepo(ctx)

	year, month, day := time.Now().Date()
	date, err := time.Parse(time.DateOnly, fmt.Sprintf("%d-%s-%s", year, utils.FixDate(int(month)), utils.FixDate(day)))
	a.NoError(err)

	history := model.LimitHistory{
		LimitID:    1,
		TimeAmount: 100,
		SentAt:     time.Now(),
		LimitDate:  date,
	}
	err = repo.Push(context.Background(), history)
	a.NoError(err)

	type dao struct {
		LimitID    int64     `db:"limit_id"`
		TimeAmount int32     `db:"time_amount"`
		SentAt     time.Time `db:"sent_at"`
		LimitDate  time.Time `db:"limit_date"`
	}

	var row dao
	err = db.QueryRow(ctx, "SELECT limit_id, time_amount, limit_date FROM limit_histories WHERE limit_id=$1 AND limit_date=$2", history.LimitID, date).
		Scan(&row.LimitID, &row.TimeAmount, &row.LimitDate)
	a.NoError(err)
	a.Equal(history.TimeAmount, row.TimeAmount)
	a.Equal(history.LimitDate.Format(time.DateOnly), row.LimitDate.Format(time.DateOnly))
}

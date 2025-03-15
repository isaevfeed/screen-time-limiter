package limit_history

import (
	"context"
	"fmt"
	"math/rand/v2"
	"screen-time-limiter/internal/domain/model"
	"screen-time-limiter/internal/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepo_Sum(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	a := assert.New(t)
	repo, db := makeRepo(ctx)

	limitID := rand.Int32()
	now := time.Now()
	year, month, day := now.Date()
	dateStr := fmt.Sprintf("%d-%s-%s", year, utils.FixDate(int(month)), utils.FixDate(day))
	date, _ := time.Parse(time.DateOnly, dateStr)
	limitModel := model.LimitHistory{
		LimitID:    limitID,
		TimeAmount: 100,
		SentAt:     now,
		LimitDate:  date,
	}
	err := repo.Push(ctx, limitModel)
	a.NoError(err)
	limitModel = model.LimitHistory{
		LimitID:    limitID,
		TimeAmount: 100,
		SentAt:     now,
		LimitDate:  date,
	}
	err = repo.Push(ctx, limitModel)
	a.NoError(err)
	limitModel = model.LimitHistory{
		LimitID:    limitID,
		TimeAmount: 100,
		SentAt:     now,
		LimitDate:  date,
	}
	err = repo.Push(ctx, limitModel)
	a.NoError(err)
	totalSum := 300

	limitSum, err := repo.Sum(ctx, limitID, date)
	a.NoError(err)
	a.Greater(limitSum, 0)
	a.Equal(limitSum, totalSum)

	_, _ = db.Exec(ctx, "DELETE FROM limit_histories WHERE limit_id = $1", limitID)
}

package limit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"screen-time-limiter/internal/domain/model"
	"testing"
	"time"
)

func TestRepo_GetByLimitID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	a := assert.New(t)
	repo, db := makeRepo(ctx)

	limitModel := model.Limit{
		UserID: rand.Int32(),
		Amount: int32((time.Hour * 2).Seconds()),
	}
	err := repo.Create(ctx, limitModel)
	a.NoError(err)
	defer func() {
		_, _ = db.Exec(ctx, "DELETE FROM limits WHERE user_id=$1", limitModel.ID)
	}()

	var limitDao dao
	_ = db.QueryRow(ctx, "SELECT id, user_id, time_limit FROM limits WHERE user_id=$1", limitModel.UserID).
		Scan(&limitDao.ID, &limitDao.UserID, &limitDao.Amount)

	limit, err := repo.GetByLimitID(ctx, limitDao.ID)
	a.NoError(err)
	a.NotNil(limit)
	a.Equal(limitModel.Amount, limit.Amount)
}

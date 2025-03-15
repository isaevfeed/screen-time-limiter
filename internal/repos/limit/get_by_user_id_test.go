package limit

import (
	"context"
	"math/rand/v2"
	"screen-time-limiter/internal/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepo_GetByUserID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	a := assert.New(t)
	repo, _ := makeRepo(ctx)

	limitModel := model.Limit{
		UserID: rand.Int32(),
		Amount: int32((time.Hour * 2).Seconds()),
	}
	err := repo.Create(ctx, limitModel)
	a.NoError(err)

	limit, err := repo.GetByUserID(ctx, limitModel.UserID)
	a.NoError(err)
	a.NotNil(limit)
	a.Equal(limit.Amount, limitModel.Amount)
}

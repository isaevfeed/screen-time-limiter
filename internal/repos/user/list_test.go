package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"screen-time-limiter/internal/domain/model"
	"strconv"
	"testing"
)

func TestRepo_List(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	a := assert.New(t)
	repo, _ := makeRepo(ctx)

	_ = repo.Add(ctx, model.User{
		FirstName: strconv.Itoa(rand.Int() * 10000),
		LastName:  strconv.Itoa(rand.Int() * 10000),
	})

	users, err := repo.List(ctx)
	a.Nil(err)
	a.NotNil(users)
	a.Greater(len(users), 0)
}

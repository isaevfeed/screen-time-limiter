package limit

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const tableName = "limits"

type (
	Repo struct {
		pool *pgxpool.Pool
	}
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{
		pool: pool,
	}
}

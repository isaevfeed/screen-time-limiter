package limit_history

import (
	"context"
	"fmt"
	"os"
	"screen-time-limiter/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func makeRepo(ctx context.Context) (*Repo, *pgxpool.Pool) {
	cfg, _ := config.Load(os.Getenv("CONFIG_FILE"))

	// Строка подключения к базе данных
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbTest.Host, cfg.DbTest.Port, cfg.DbTest.Username, cfg.DbTest.Password, cfg.DbTest.Database)

	// Открываем соединение с базой данных
	db, _ := pgxpool.New(ctx, psqlInfo)

	return NewRepo(db), db
}

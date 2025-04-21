package service_provider

import (
	"context"
	"fmt"
	"screen-time-limiter/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func mustInitDatabase(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	pool, err := initDatabase(ctx, cfg)
	if err != nil {
		panic(err)
	}

	return pool
}

func initDatabase(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	// Строка подключения к базе данных
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.Username, cfg.Db.Password, cfg.Db.Database)

	pool, err := pgxpool.New(ctx, psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("pool.Ping: %w", err)
	}

	return pool, nil
}

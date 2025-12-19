package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seeques/task-api/internal/config"
)

type PostgresStorage struct {
	pool *pgxpool.Pool
}

func NewPostgresStorage(pool *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{pool: pool}
}

func CreatePool() (*pgxpool.Pool, error) {
	cfg := config.LoadConfig()

	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (s *PostgresStorage) ClosePool() {
	s.pool.Close()
}

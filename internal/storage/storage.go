package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seeques/task-api/internal/config"
)

type PostgresStorage struct {
	pool *pgxpool.Pool
}

type Storage interface {
	SaveUser(ctx context.Context, usr *User) error
}

func (s *PostgresStorage) SaveUser(ctx context.Context, usr *User) error {
	sql := `INSERT INTO users (email, password_hash)
	VALUES ($1, $2)
	RETURNING id, created_at`

	_, err := s.pool.QueryRow(ctx, sql, usr.Email, usr.PasswordHash).Scan(&usr.ID, &usr.CreatedAt)
	return err
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

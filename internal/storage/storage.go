package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx"
	"github.com/seeques/task-api/internal/config"
)

type PostgresStorage struct {
	pool *pgxpool.Pool
}

type Storage interface {
	SaveUser(ctx context.Context, usr *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

func (s *PostgresStorage) SaveUser(ctx context.Context, usr *User) error {
	sql := `INSERT INTO users (email, password_hash)
	VALUES ($1, $2)
	RETURNING id, created_at`

	err := s.pool.QueryRow(ctx, sql, usr.Email, usr.PasswordHash).Scan(&usr.ID, &usr.CreatedAt)
	return err
}

func (s *PostgresStorage) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	sql := `SELECT id, password_hash, created_at FROM users WHERE email = $1`

	usr := &User{
		Email: email,
	}

	err := s.pool.QueryRow(ctx, sql, email).Scan(usr.ID, usr.PasswordHash, usr.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return usr, nil
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

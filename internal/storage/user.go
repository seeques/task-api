package storage

import (
	"time"
	"net/http"
	"context"
	"errors"

	"github.com/jackc/pgx"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt`
}

// define new type for user_id key to prevent collisions with other packages
// which also might use user_id as key
type contextKey string

const UserIDKey contextKey = "user_id"

// helper to get the userID
func GetUserID(r *http.Request) int {
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		return 0
	}

	return userID
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
	query := `SELECT id, password_hash, created_at FROM users WHERE email = $1`

	usr := &User{
		Email: email,
	}

	err := s.pool.QueryRow(ctx, query, email).Scan(&usr.ID, &usr.PasswordHash, &usr.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return usr, nil
} 
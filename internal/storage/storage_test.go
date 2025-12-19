package storage

import (
    "context"
    "testing"
    
    "github.com/jackc/pgx/v5/pgxpool"
)

const testDatabaseURL = "postgres://taskapi:taskapi@localhost:5433/task_api_test?sslmode=disable"

func setUpTestStorage (t *testing.T) *PostgresStorage {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, testDatabaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}

	cleanTables(t, pool)

	return NewPostgresStorage(pool)
}

func cleanTables(t *testing.T, pool *pgxpool.Pool) {
	ctx := context.Background()

	_, err := pool.Exec(ctx, "DELETE FROM tasks")
	if err != nil {
		t.Fatalf("failed to clean tasks: %v", err)
	}

	_, err = pool.Exec(ctx, "DELETE FROM users")
	if err != nil {
		t.Fatalf("failed to clean users: %v", err)
	}
} 
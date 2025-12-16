package storage

import (
	"context"
	"time"

)

type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskStorage interface {
	CreateTask(ctx context.Context, task *Task) error
}

func (s *PostgresStorage) CreateTask(ctx context.Context, task Task) error {
	newTaskSql := `INSERT INTO tasks (user_id, title, description) 
	VALUES ($1, $2, $3)
	RETURNING id, completed, created_at, updated_at`

	err := s.pool.QueryRow(ctx, newTaskSql, task.UserID, task.Title, task.Description).Scan(
		&task.ID, 
		&task.Completed, 
		&task.CreatedAt, 
		&task.UpdatedAt,
	)
	return err
}
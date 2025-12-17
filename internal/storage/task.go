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
	GetTask(ctx context.Context, taskID int, userID int)
}

func (s *PostgresStorage) CreateTask(ctx context.Context, task *Task) error {
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

func (s *PostgresStorage) GetTask(ctx context.Context, taskID int, userID int) (*Task, error) {
	query := `SELECT id, user_id, title, description, completed, created_at, updated_at 
    FROM tasks WHERE id = $1 AND user_id = $2`

	task := Task{}

    err := s.pool.QueryRow(ctx, query, taskID, userID).Scan(
        &task.ID,
        &task.UserID,
        &task.Title, 
        &task.Description, 
        &task.Completed, 
        &task.CreatedAt, 
        &task.UpdatedAt,
    )
	return &task, err
}

func (s *PostgresStorage) ListTasks(ctx context.Context, userID int) ([]Task, error) {
	query := `SELECT id, user_id, title, description, completed, created_at, updated_at 
    FROM tasks WHERE user_id = $1`

	rows, err := s.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
		&task.ID,
        &task.UserID,
        &task.Title, 
        &task.Description, 
        &task.Completed, 
        &task.CreatedAt, 
        &task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
        return nil, err
    }

	return tasks, nil
}
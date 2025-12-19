package storage

import (
    "context"
    "testing"
)

func createTestUser(t *testing.T, s *PostgresStorage) *User {
    user := &User{
        Email:        "user@example.com",
        PasswordHash: "hash",
    }
    err := s.SaveUser(context.Background(), user)
    if err != nil {
        t.Fatalf("failed to create test user: %v", err)
    }
    return user
}

func TestCreateTask(t *testing.T) {
	s := setUpTestStorage(t)
    ctx := context.Background()
    
    user := createTestUser(t, s)

	task := &Task{
		UserID:      user.ID,
        Title:       "Test Task",
        Description: "Test Description",
	}

	err := s.CreateTask(ctx, task)
	if err != nil {
		t.Fatalf("failed to create task, %v", err)
	}

	if task.ID == 0 {
		t.Fatal("expected taks ID")
	}
}
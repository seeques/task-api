package storage

import (
    "context"
    "testing"
)

func TestSaveUser(t *testing.T) {
    s := setUpTestStorage(t)
    ctx := context.Background()
    
    user := &User{
        Email:        "test@example.com",
        PasswordHash: "hashed-password",
    }
    
    err := s.SaveUser(ctx, user)
    if err != nil {
        t.Fatalf("failed to save user: %v", err)
    }
    
    if user.ID == 0 {
        t.Error("expected user ID to be set")
    }
    
    if user.CreatedAt.IsZero() {
        t.Error("expected CreatedAt to be set")
    }
}

func TestGetUserByEmail(t *testing.T) {
	s := setUpTestStorage(t)
    ctx := context.Background()

	user := &User{
        Email:        "test@example.com",
        PasswordHash: "hashed-password",
    }

	err := s.SaveUser(ctx, user)
    if err != nil {
        t.Fatalf("failed to save user: %v", err)
    }

	found, err := s.GetUserByEmail(ctx, user.Email)

	if found.Email != "test@example.com" {
        t.Fatalf("expected email %s, got %s", "test@example.com", found.Email)
    }
}
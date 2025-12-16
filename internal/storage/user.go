package storage

import (
	"time"
	"net/http"
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
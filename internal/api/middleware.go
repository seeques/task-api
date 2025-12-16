package api

import (
	"net/http"
	"strings"
	"context"

	"github.com/seeques/task-api/internal/config"
	"github.com/seeques/task-api/internal/response"
	"github.com/seeques/task-api/internal/auth"
)

// define new type for user_id key to prevent collisions with other packages
// which also might use user_id as key
type contextKey string

const UserIDKey contextKey = "user_id"

type Middleware struct {
	cfg *config.Config
}

// constructor for Middleware struct
func NewMiddleware(cfg *config.Config) *Middleware {
	return &Middleware{cfg: cfg}
}

// helper to get the userID
func GetUserID(r *http.Request) int {
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		return 0
	}

	return userID
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.RespondError(w, http.StatusUnauthorized, "missig authorization header")
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			response.RespondError(w, http.StatusUnauthorized, "invalid authorization format")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, prefix)

		userID, err := auth.ValidateToken(tokenString, m.cfg.JWTsecret)
		if err != nil {
			response.RespondError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
package api

import (
	"net/http"
	"strings"
	"context"

	"github.com/seeques/task-api/internal/config"
	"github.com/seeques/task-api/internal/response"
	"github.com/seeques/task-api/internal/auth"
	"github.com/seeques/task-api/internal/storage"
)

type Middleware struct {
	cfg *config.Config
}

// constructor for Middleware struct
func NewMiddleware(cfg *config.Config) *Middleware {
	return &Middleware{cfg: cfg}
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

		ctx := context.WithValue(r.Context(), storage.UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
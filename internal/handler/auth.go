package handler

import (
	"net/http"
	"encoding/json"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/seeques/task-api/internal/response"
	"github.com/seeques/task-api/internal/auth"
	"github.com/seeques/task-api/internal/storage"
)

type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	storage *storage.PostgresStorage
}

func NewHandler(storage *storage.PostgresStorage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if !strings.Contains(req.Email, "@") {
		response.RespondError(w, http.StatusBadRequest, "invalid email")
		return
	}

	if strings.Count(req.Password, "") < 8 {
		response.RespondError(w, http.StatusBadRequest, "password needs at least 8 characters")
		return
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "password hashing failed")
		return
	}

	usr := &storage.User{
		Email: req.Email,
		PasswordHash: hash,
	}

	err = h.storage.SaveUser(r.Context(), usr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to save user's data")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usr)
}
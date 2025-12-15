package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/seeques/task-api/internal/auth"
	"github.com/seeques/task-api/internal/response"
	"github.com/seeques/task-api/internal/storage"
	"github.com/seeques/task-api/internal/config"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:password`
}

type Handler struct {
	storage *storage.PostgresStorage
	cfg *config.Config
}

func NewHandler(storage *storage.PostgresStorage, cfg *config.Config) *Handler {
	return &Handler{
		storage: storage,
		cfg: cfg,
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
		Email:        req.Email,
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	usr, err := h.storage.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "database error")
		return
	}
	if usr == nil {
		response.RespondError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	if err := auth.CheckPassword(req.Password, usr.PasswordHash); err != nil {
    	response.RespondError(w, http.StatusUnauthorized, "invalid credentials")
    	return
	}

	token, err := auth.GenerateToken(usr.ID, h.cfg.JWTsecret)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "token creation error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
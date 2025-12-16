package handler

import (
	"encoding/json"
	"net/http"

	"github.com/seeques/task-api/internal/storage"
	"github.com/seeques/task-api/internal/response"
)

type CreateTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID := storage.GetUserID(r)

	var req CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if req.Title == "" {
		response.RespondError(w, http.StatusBadRequest, "title is required")
		return
	}

	task := &storage.Task{
		UserID: userID,
		Title: req.Title,
		Description: req.Description,
	}

	err = h.storage.CreateTask(r.Context(), task)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "task creation failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

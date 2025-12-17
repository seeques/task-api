package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"errors"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/seeques/task-api/internal/response"
	"github.com/seeques/task-api/internal/storage"
)

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
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
		UserID:      userID,
		Title:       req.Title,
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

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	userID := storage.GetUserID(r)

	idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, "invalid task id")
		return
	}

	if id == 0 {
		response.RespondError(w, http.StatusBadRequest, "task id is required")
		return
	}

	task, err := h.storage.GetTask(r.Context(), id, userID)
	if errors.Is(err, pgx.ErrNoRows) {
        response.RespondError(w, http.StatusNotFound, "task not found")
        return
    }
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "update failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	userID := storage.GetUserID(r)

	tasks, err := h.storage.ListTasks(r.Context(), userID)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "unable to get tasks")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID := storage.GetUserID(r)

	idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, "invalid task id")
		return
	}

	var req UpdateTaskRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
        response.RespondError(w, http.StatusBadRequest, "invalid JSON")
        return
    }

	if req.Title == "" {
		response.RespondError(w, http.StatusBadRequest, "title is required")
		return
	}

	task := &storage.Task{
		ID: id,
		UserID: userID,
		Title: req.Title,
		Description: req.Description,
		Completed: req.Completed,
	}

	err = h.storage.UpdateTask(r.Context(), task)

	if errors.Is(err, pgx.ErrNoRows) {
        response.RespondError(w, http.StatusNotFound, "task not found")
        return
    }
	if err != nil {
        response.RespondError(w, http.StatusInternalServerError, "update failed")
        return
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(task)
}
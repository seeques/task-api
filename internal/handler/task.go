package handler

import (
	"net/http"
	"encoding/json"

	"github.com/seeques/task-api/internal/storage"
)

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
    userID := storage.GetUserID(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"userID": userID})
}
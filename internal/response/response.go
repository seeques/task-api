package response

import (
	"net/http"
	"encoding/json"
)

func RespondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Contetnt-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
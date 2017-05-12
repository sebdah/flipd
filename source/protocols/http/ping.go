package http

import (
	"encoding/json"
	"net/http"
)

// pingResponse defines what the ping response structure should look like.
type pingResponse struct {
	Message string `json:"message"`
}

// Ping responds to ping requests.
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&pingResponse{
		Message: "Pong",
	})
}

package http

import (
	"encoding/json"
	"net/http"

	"github.com/sebdah/flipd/source/types"
)

// getFeaturesResponse defines the response body of the GetFeatures endpoint.
type getFeaturesResponse struct {
	Features []*types.Feature `json:"features"`
}

// GetFeatures returns a list of all registered features.
func (h *Handler) GetFeatures(w http.ResponseWriter, r *http.Request) {
	features, err := h.storage.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&getFeaturesResponse{
		Features: features,
	})

	return
}

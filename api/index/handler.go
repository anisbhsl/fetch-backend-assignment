package index

import (
	"encoding/json"
	"net/http"
)

// Index is http handler for "/" endpoint
func (s service) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "Welcome to Fetch Receipt Processor API",
		})
	}
}

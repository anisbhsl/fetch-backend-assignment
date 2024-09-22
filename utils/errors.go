package utils

import (
	"encoding/json"
	"net/http"
)

// SendErrorResponse is a wrraper to send error response
func SendErrorResponse(w http.ResponseWriter, errorMessage interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   errorMessage,
	})
}

// SendSuccessResponse is a wrapper to send 200 response
func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

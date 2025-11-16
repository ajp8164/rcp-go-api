package handlers

import (
	"encoding/json"
	"net/http"
	"rcp-go-api/api"
)

// Returns a welcome message.
func Welcome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var response = api.WelcomeResponse{
		Code: http.StatusOK,
		Message: "Welcome to the RCP API!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

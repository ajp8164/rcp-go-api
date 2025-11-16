package api

import (
	"encoding/json"
	"net/http"
	"rcp-go-api/internal/tools"
)

// GET welcome
type WelcomeResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
}

// GET clubs
type ClubsResponse struct {
	Code    int         `json:"code"`
	Clubs []tools.Club  `json:"clubs"`
}

// Error handling helpers.
type Error struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
}

func WriteError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)

package handlers

import (
	"encoding/json"
	"net/http"
	"rcp-go-api/api"
	"rcp-go-api/internal/tools"
)

// Returns all clubs.
func GetClubs(w http.ResponseWriter, r *http.Request) {
	var err error
	var database *tools.DatabaseInterface

	database, err = tools.NewDatabase()
	if (err != nil) {
		api.InternalErrorHandler(w, err)
		return 
	}

	clubs := (*database).GetClubs()

	var response = api.ClubsResponse{
		Code: http.StatusOK,
		Clubs: clubs,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if (err != nil) {
		api.InternalErrorHandler(w, err)
		return 
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"rcp-go-api/internal/handlers"
	"rcp-go-api/internal/middleware"
)

func main() {
	 // Create a new ServeMux (not using default).
	mux := http.NewServeMux()

	// Request handlers.
	// Middlewae ensures only correct method is matched.
	mux.Handle("/welcome", middleware.GET(http.HandlerFunc(handlers.Welcome)))
	mux.Handle("/clubs", middleware.GET(http.HandlerFunc(handlers.GetClubs)))

	// Apply middleware to request handlers,
	// Middleware functions are executed bottom to top.
	wrapped := middleware.Chain(
		mux,
    middleware.StripTrailingSlashes,
    middleware.Authorization,
	)

	// Start the server on port localhost:8080
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	err := http.ListenAndServe(port, wrapped)

	if (err != nil) {
		log.Fatal(err)
	}
}

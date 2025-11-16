package middleware

import (
	"net/http"
	"strings"
)

// Removes trailing slashes from the URL path.
func StripTrailingSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimRight(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

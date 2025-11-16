package middleware

import (
	"errors"
	"log"
	"net/http"
	"rcp-go-api/api"
	"strings"
)

var ErrUnauthorizedError = errors.New("invalid authorization token")

// Checks for a valid authorization token.
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string = r.Header.Get("Authorization")

		// Must be a Basic authg token.
		isBasic := strings.HasPrefix(token, "Basic")
		token = strings.Trim(strings.ReplaceAll(token, "Basic", ""), " ")

		if (!isBasic || token == "") {
			log.Println(ErrUnauthorizedError)
			api.RequestErrorHandler(w, ErrUnauthorizedError)
			return
		}

		// Token value can be anything for now..

		next.ServeHTTP(w, r)
	})
}

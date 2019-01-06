package middleware

import (
	"net/http"
)

// A middleware for injecting tokens into the Authorization header when forwarding to services
var TokenMiddlewareFactory = func(token string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
		// Override the Authorization token if specified,
		// otherwise assume the Authorization header is to be passed through
		if token != "" {
			r.Header.Set("Authorization", token)
		}
	})
}

package proxy

import (
	"net/http"
)

// A middleware which performs basic preprocessing including sanitization and authorization
var PreprocessingMiddleware http.Handler

// A middleware for injecting tokens into the Authorization header when forwarding to services
var TokenMiddlewareFactory func(token string) http.Handler

func init() {
	PreprocessingMiddleware = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
		err := requestPreprocessing(w, r)

		if err != nil {
			w.WriteHeader(500)
		}
	})

	TokenMiddlewareFactory = func(token string) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
			// Override the Authorization token if specified,
			// otherwise assume the Authorization header is to be passed through
			if token != "" {
				r.Header.Set("Authorization", token)
			}
		})
	}
}

package proxy

import (
	"net/http"
)

var PreprocessingMiddleware http.Handler

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
			if token != "" {
				r.Header.Set("Authorization", token)
			}
		})
	}
}

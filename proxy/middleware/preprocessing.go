package middleware

import (
	"net/http"
)

// A middleware which performs basic preprocessing including sanitization and authorization
var PreprocessingMiddleware = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
	err := requestPreprocessing(w, r)

	if err != nil {
		w.WriteHeader(500)
	}
})

package middleware

import (
	"net/http"

	"github.com/arbor-dev/arbor/proxy/constants"
)
// A middleware for handling CORS
var CORSMiddleware = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", r.Method)
	w.Header().Set("Access-Control-Allow-Headers", constants.AccessControlAllowHeaders)
})

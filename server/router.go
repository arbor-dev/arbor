/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package server

import (
	"io"
	"net/http"
	"time"

	"github.com/arbor-dev/arbor/services"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	logRequest(r.Method, r.RequestURI, "UNKNOWN", http.StatusNotFound, time.Duration(0))
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "404 Not Found")
}

func corsPreflight(w http.ResponseWriter, r *http.Request) {
	logRequest(r.Method, r.RequestURI, "CORS Preflight", http.StatusOK, time.Duration(0))
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, CONNECT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin")
	w.WriteHeader(http.StatusOK)
}

func NewRouter(routes services.RouteCollection) *mux.Router {

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(notFound)
	for _, route := range routes {
		var handler http.Handler

		handler = route.Handler
		//Log request
		handler = httpLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		// Handle CORS preflight requests
		router.
			Methods("OPTIONS").
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(corsPreflight)
	}
	return router
}

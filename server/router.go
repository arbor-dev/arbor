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
	logRequest(r, "Unknown", 404, time.Duration(0))
	w.WriteHeader(404)
	io.WriteString(w, "404 Not Found")
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

	}
	return router
}

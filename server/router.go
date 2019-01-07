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
	"strings"
	"time"

	"github.com/arbor-dev/arbor/services"
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/proxy/constants"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	logRequest(r.Method, r.RequestURI, "UNKNOWN", http.StatusNotFound, time.Duration(0))
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "404 Not Found")
}

func corsPreflight(methods []string) http.HandlerFunc {
	allowedMethods := strings.Join(methods, ", ")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", proxy.AccessControlPolicy)
		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", constants.AccessControlAllowHeaders)
		w.WriteHeader(http.StatusOK)
	}
}

func buildPreflightRoutes(routes services.RouteCollection) services.RouteCollection {
	allowedMethods := make(map[string][]string)

	for _, route := range routes {
		_, exists := allowedMethods[route.Pattern]

		if !exists {
			allowedMethods[route.Pattern] = []string{"OPTIONS", "CONNECT"}
		}

		allowedMethods[route.Pattern] = append(allowedMethods[route.Pattern], route.Method)
	}

	var preflightRoutes []services.Route

	for pattern, methods := range allowedMethods {
		preflightRoutes = append(preflightRoutes, services.Route {
			"Preflight",
			"OPTIONS",
			pattern,
			corsPreflight(methods),
		})
	}

	return preflightRoutes
}

func NewRouter(routes services.RouteCollection) *mux.Router {

	routes = append(routes, buildPreflightRoutes(routes)...)

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

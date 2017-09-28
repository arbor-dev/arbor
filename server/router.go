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
	"net/http"

	"github.com/acm-uiuc/arbor/services"
	"github.com/gorilla/mux"
)

func NewRouter(routes services.RouteCollection) *mux.Router {

	router := mux.NewRouter()
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
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	return router
}

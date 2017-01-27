/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.  
* 
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package main

import (
	// "log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/acm-uiuc/groot/services"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	for _, route := range services.Routes {
		var handler http.Handler

		handler = route.HandlerFunc
		//Log request
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	return router
}

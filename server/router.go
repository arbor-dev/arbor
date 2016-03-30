package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/acm-uiuc/groot/services"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	for _, route := range services.Routes {
		var handler http.Handler

		handler = route.HandlerFunc
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

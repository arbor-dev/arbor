package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAPIs() {
	routes = append(routes, todoRoutes...)
}

func NewRouter() *mux.Router {

	router := mux.NewRouter()
	for _, route := range routes {
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

package services

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type RouteCollection []Route

var Routes = RouteCollection{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func RegisterAPIs() {
	Routes = append(Routes, TodoRoutes...)
	Routes = append(Routes, GrandpaRoutes...)
	Routes = append(Routes, BearsRoutes...)
	Routes = append(Routes, TestRoutes...)
	Routes = append(Routes, AuthRoutes...)
    Routes = append(Routes, UsersRoutes...)
	Routes = append(Routes, QuotesRoutes...)
}

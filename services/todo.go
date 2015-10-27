package services

import (
	"net/http"

	"github.com/acm-uiuc/groot/proxy"
)

//Location
var TodoURL string = "http://localhost:8080"

//API Interface
var TodoRoutes = RouteCollection{
	Route{
		"TodoIndex",
		"GET",
		"/todo",
		todoIndex,
	},
	Route{
		"TodoAll",
		"GET",
		"/todos",
		TodoAll,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}

//Route handler
func todoIndex(w http.ResponseWriter, r *http.Request) {
	//TODO: THIS API CALL IS A SPECIAL CASE
	proxy.GETHandler(w, TodoURL, r)
}

func TodoAll(w http.ResponseWriter, r *http.Request) {
	proxy.GETHandler(w, TodoURL+r.URL.String(), r)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	proxy.GETHandler(w, TodoURL+r.URL.String(), r)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	//proxy.POSTHandler()
}

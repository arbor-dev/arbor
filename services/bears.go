package services

import (
	"net/http"

	"github.com/acm-uiuc/groot/proxy"
)

//Location
var BearsURL string = "http://localhost:8080/api"

//API Interface
var BearsRoutes = RouteCollection{
	Route{
		"GetAllBears",
		"GET",
		"/bears",
		GetAllBears,
	},
	Route{
		"CreateBear",
		"POST",
		"/bears",
		CreateBear,
	},
	Route{
		"GetBear",
		"GET",
		"/bears/{bear_id}",
		GetBear,
	},
	Route{
		"UpdateBear",
		"PUT",
		"/bears/{bear_id}",
		UpdateBear,
	},
	Route{
		"DeleteBear",
		"DELETE",
		"/bears/{bear_id}",
		DeleteBear,
	},
}

//Route handler
func GetAllBears(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, GrandpaURL+r.URL.String(), r)
}

func CreateBear(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, GrandpaURL+r.URL.String(), r)
}
func GetBear(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, GrandpaURL+r.URL.String(), r)
}
func UpdateBear(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, GrandpaURL+r.URL.String(), r)
}
func DeleteBear(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, GrandpaURL+r.URL.String(), r)
}

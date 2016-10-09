package services

import (
	"github.com/acm-uiuc/groot/proxy"
	"net/http"
)

//Location
const HardwareURL string = "http://localhost:4523/api/v1.0"

//Client Data Type
const HardwareFormat string = "JSON"

//API Interface
var HardwareRoutes = RouteCollection {
	Route{
		"NewItem",
		"POST",
		"/items/{id}",
		NewSession,
	},
	Route{
		"UpdateItem",
		"PUT",
		"/items/{id}",
		EndUsersSessions,
	},
	Route{
		"DeleteItem",
		"DELETE",
		"/item/{id}",
		GetAuthenticatedUser,
	},
}

//Route handler
func NewItem(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, AuthURL+r.URL.String(), AuthFormat, r)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, AuthURL+r.URL.String(), AuthFormat, r)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, AuthURL+r.URL.String(), AuthFormat, r)
}

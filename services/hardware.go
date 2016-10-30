package services

import (
	"github.com/acm-uiuc/groot/proxy"
	"net/http"
)

//Location
const HardwareURL string = "http://localhost:4523/api/v1.0"

//Service Data Type
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
	proxy.POST(w, HardwareURL+r.URL.String(), HardwareFormat, "", r)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, HardwareURL+r.URL.String(), HardwareFormat, "", r)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, HardwareURL+r.URL.String(), HardwareFormat, "", r)
}

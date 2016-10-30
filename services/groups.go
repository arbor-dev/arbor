package services

import (
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
)

//Location
const GroupsURL string = "http://localhost:9001"

//Service Data Type
const GroupsFormat string = "JSON"

//API Interface
var GroupsRoutes = RouteCollection {
	Route{
		"GetGroups",
		"GET",
		"/items/{id}",
		GetGroups,
	},
}

//Route handler
func GetGroups(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

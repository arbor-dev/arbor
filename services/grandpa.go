package services

import (
	"fmt"
	"net/http"

	"github.com/acm-uiuc/groot/proxy"
)

//Location
var GrandpaURL string = "http://grandpa.cs.illinois.edu:8080"

//API Interface
var GrandpaRoutes = RouteCollection{
	Route{
		"grandpaCorpora",
		"GET",
		"/classification",
		grandpaCorpora,
	},
}

//Route handler
func grandpaCorpora(w http.ResponseWriter, r *http.Request) {
	fmt.Println(GrandpaURL + r.URL.String())
	proxy.GET(w, GrandpaURL+r.URL.String(), r)
}

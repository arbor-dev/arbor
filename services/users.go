package services

import (
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
    "github.com/acm-uiuc/groot/secrets"
)

//Location
const UsersURL string = secrets.CrowdURL

const UserToken string = secrets.CrowdToken

//Service Data Type
const UserFormat string = "JSON"

//API Interface
var UsersRoutes = RouteCollection{
    Route{
        "GetUser",
        "GET",
        "/user",
        GetUser,
    },
}

//Route handler
func GetUser(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, UsersURL+r.URL.String(), UserFormat, UserToken, r)
}

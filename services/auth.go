package services

import (
	// "log"
	"github.com/acm-uiuc/groot/proxy"
	"github.com/acm-uiuc/groot/secrets"
	"net/http"
)

//Location
const AuthURL string = secrets.CrowdURL;

//token
const AuthToken string = secrets.CrowdToken;

//Service Data Type
const AuthFormat string = "JSON"

//API Interface
var AuthRoutes = RouteCollection{
	Route{
		"NewSession",
		"POST",
		"/session",
		NewSession,
	},
	Route{
		"EndUsersSessions",
		"DELETE",
		"/session?username={username}",
		EndUsersSessions,
	},
	Route{
		"GetAuthenticatedUser",
		"GET",
		"/session/{token}",
		GetAuthenticatedUser,
	},
	Route{
		"ValidateSession",
		"POST",
		"/session/{token}",
		ValidateSession,
	},
	Route{
		"EndSession",
		"DELETE",
		"/session/{token}",
		EndSession,
	},
}

//Route handler
// w = writer, r = reader
func NewSession(w http.ResponseWriter, r *http.Request) {
	//Auth: THIS API CALL IS A SPECIAL CASE
	// log.Printf("new session called")
	proxy.POST(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func EndUsersSessions(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func ValidateSession(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

func EndSession(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, AuthURL+r.URL.String(), AuthFormat, AuthToken, r)
}

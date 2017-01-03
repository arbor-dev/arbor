/**
* Copyright © 2016, ACM@UIUC
*
* This file is part of the Groot Project.  
* 
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package services

import (
	"github.com/acm-uiuc/groot/proxy"
	"github.com/acm-uiuc/groot/config"
	"net/http"
)

//Location
const AuthURL string = config.CrowdURL;

//token
const AuthToken string = config.AuthPrefix + config.CrowdToken;

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

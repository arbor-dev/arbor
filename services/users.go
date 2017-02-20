/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package services

import (
	"net/http"

	"github.com/acm-uiuc/groot/proxy"
)

//Location
const UsersURL string = "http://localhost:8001"

//Service Data Type
const UsersFormat string = "JSON"

//API Interface
var UsersRoutes = RouteCollection{
	Route{
		"GetUsers",
		"GET",
		"/users",
		GetUsers,
	},
	Route{
		"IsUser",
		"POST",
		"/users/{netid}/isUser",
		IsUser,
	},
	Route{
		"NewUser",
		"POST",
		"/users",
		NewUser,
	},
	Route{
		"ConfirmUser",
		"PUT",
		"/users/{user_id}/paid",
		ConfirmUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/users/{user_id}",
		DeleteUser,
	},
	Route{
		"UsersLogin",
		"POST",
		"/users/login",
		UsersLogin,
	},
	Route{
		"UsersLogout",
		"POST",
		"/users/logout",
		UsersLogout,
	},
}

//Route handler
func GetUsers(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func IsUser(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func ConfirmUser(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func UsersLogin(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

func UsersLogout(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, UsersURL+r.URL.String(), UsersFormat, "", r)
}

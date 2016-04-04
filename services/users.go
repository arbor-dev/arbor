package services

import (
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
)

//Location
var UsersURL string = "http://localhost:4567"

//Client Data Type
var UserFormat string = "XML"

//API Interface
var UsersRoutes = RouteCollection{
	Route{
		"GetAllUsers",
		"GET",
		"/users",
		GetAllUsers,
	},
    Route{
		"DeleteUser",
		"DELETE",
		"/users/{netid}",
		DeleteUser,
	},
    Route{
		"GetUser",
		"GET",
		"/users/{netid}",
		GetUser,
	},
    Route{
		"CreateUser",
		"POST",
		"/users",
		CreateUser,
	},
    Route{
		"UpdateUser",
		"PUT",
		"/users/{netid}",
		UpdateUser,
	},
}

//Route handler
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, UsersURL+r.URL.String(), UserFormat, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	proxy.DELETE(w, UsersURL+r.URL.String(), UserFormat, r)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, UsersURL+r.URL.String(), UserFormat, r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, UsersURL+r.URL.String(), UserFormat, r)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	proxy.PUT(w, UsersURL+r.URL.String(), UserFormat, r)
}

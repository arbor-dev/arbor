package services

import (
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
    "github.com/acm-uiuc/groot/secrets"
)

//Location
const MembersURL string = "http://localhost:8001"

//Service Data Type
const MembersFormat string = "JSON"

//API Interface
var MemberRoutes = RouteCollection{
    Route{
        "PreMembers",
        "POST",
        "/users/pre",
        PreMembers,
    },
    Route{
        "CurrentMembers",
        "POST",
        "/users/current",
        CurrentMembers,
    },
    Route{
        "MemberInfo",
        "POST",
        "/users/{netid}",
        MemberInfo,
    },
    Route{
        "IsMember",
        "POST",
        "/users/{netid}/isMember",
        IsMember,
    },
    Route{
        "NewMember",
        "POST",
        "/newUser",
        NewMember,
    },
    Route{
        "ConfirmMember",
        "POST",
        "/users/paid",
        ConfirmMember,
    },
}

//Route handler
func PreMembers(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, UsersURL+r.URL.String(), UserFormat, "", r)
}

func CurrentMembers(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, UsersURL+r.URL.String(), UserFormat, "", r)
}

func MemberInfo(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, UsersURL+r.URL.String(), UserFormat, "", r)
}

func IsMember(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, UsersURL+r.URL.String(), UserFormat, "", r)
}

func NewMember(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, UsersURL+r.URL.String(), UserFormat, "", r)
}

func ConfirmMember(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, UsersURL+r.URL.String(), UserFormat, "", r)
}

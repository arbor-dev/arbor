package services

import (
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
)

//Location
const MembersURL string = "http://localhost:8001"

//Service Data Type
const MembersFormat string = "JSON"

//API Interface
var MembersRoutes = RouteCollection{
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
	proxy.POST(w, MembersURL+r.URL.String(), MembersFormat, "", r)
}

func CurrentMembers(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, MembersURL+r.URL.String(), MembersFormat, "", r)
}

func MemberInfo(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, MembersURL+r.URL.String(), MembersFormat, "", r)
}

func IsMember(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, MembersURL+r.URL.String(), MembersFormat, "", r)
}

func NewMember(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, MembersURL+r.URL.String(), MembersFormat, "", r)
}

func ConfirmMember(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, MembersURL+r.URL.String(), MembersFormat, "", r)
}

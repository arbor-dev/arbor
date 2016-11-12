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
		"GetGroupTypes",
		"GET",
		"/groups",
		GetGroupsTypes,
	},
	Route{
		"GetGroups",
		"GET",
		"/groups/:groupType",
		GetGroups,
	},
	Route{
		"GetGroup",
		"GET",
		"/groups/{groupType}/{groupName}",
		GetGroup,
	},
	Route{
		"IsMember",
		"GET",
		"groups/{groupType}/{groupName}?isMember={netid}",
		IsMember,
	},
}

//Route handler
func GetGroupTypes(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

func IsMember(w http.ResponseWriter, r *http.Request) {
	proxy.GET(w, AuthURL+r.URL.String(), AuthFormat, "", r)
}

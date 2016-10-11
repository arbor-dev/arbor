package services

import (
	"net/http"
	"github.com/acm-uiuc/groot/proxy"
)

//Location
const ResumeURL string = "http://localhost:4567"

const ResumeToken string = ""

//Service Data Type
const ResumeFormat string = "JSON"

//API Interface
var ResumeRoutes = RouteCollection{
    Route{
        "PostResume",
        "POST",
        "/resume",
        PostResume,
    },
}

//Route handler
func PostResume(w http.ResponseWriter, r *http.Request) {
	proxy.POST(w, ResumeURL+r.URL.String(), ResumeFormat, ResumeToken, r)
}

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

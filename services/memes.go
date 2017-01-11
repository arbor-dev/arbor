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
const MemeURL string = "http://localhost:42069"

//Service Data Type
const MemeFormat string = "JSON"

//API Interface
var MemeRoutes = RouteCollection{
    Route{
        "ListMemes",
        "GET",
        "/memes",
        ListMemes,
    },
    Route{
        "NewMeme",
        "POST",
        "/memes",
        NewMeme,
    },
    Route{
        "MemeInfo",
        "GET",
        "/memes/{meme_id}",
        MemeInfo,
    },
}

//Route handler
func ListMemes(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func NewMeme(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

func MemeInfo(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, MemeURL+r.URL.String(), MemeFormat, "", r)
}

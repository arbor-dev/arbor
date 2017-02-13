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
    "github.com/acm-uiuc/groot/proxy"
    "net/http"
)

//Location
const CreditsURL string = "http://localhost:8765"

//Service Data Type
const CreditsFormat string = "JSON"

//API Interface
var CreditsRoutes = RouteCollection {
    Route{
        "NewPayment",
        "POST",
        "/payment",
        NewPayment,
    },
    Route{
        "GetCreditsUser",
        "GET",
        "/credits/users/{netid}",
        GetCreditsUser,
    },
    Route{
        "GetCreditsUserMultiple",
        "GET",
        "/credits/users",
        GetCreditsUserMultiple,
    },
    Route{
        "GetTransaction",
        "GET",
        "/credits/transactions/{id}",
        GetTransaction,
    },
    Route{
        "CreateTransaction",
        "POST",
        "/credits/transactions",
        CreateTransaction,
    },
    Route{
        "DeleteTransaction",
        "DELETE",
        "/credits/transactions",
        DeleteTransaction,
    },
}

//Route handler
func NewPayment(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetCreditsUser(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetCreditsUserMultiple(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
    proxy.GET(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
    proxy.POST(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
    proxy.DELETE(w, CreditsURL+r.URL.String(), CreditsFormat, "", r)
}

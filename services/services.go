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
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type RouteCollection []Route

var Routes = RouteCollection{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func RegisterAPIs() {
	Routes = append(Routes, AuthRoutes...)
	Routes = append(Routes, GroupsRoutes...)
	Routes = append(Routes, MembersRoutes...)
	Routes = append(Routes, QuotesRoutes...)
	Routes = append(Routes, RecruitersRoutes...)
	Routes = append(Routes, HardwareRoutes...)
	Routes = append(Routes, MemeRoutes...)
	Routes = append(Routes, EventsRoutes...)
}

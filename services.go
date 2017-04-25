/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package arbor

import "net/http"

// Route is a struct that defines a route for a microservice
// Name: Name of the route
// Method: The type of request (GET, POST, DELETE, etc.)
// Pattern: The exposed url pattern for clients to hit, allows for url encoded variables to be specified with {VARIABLE}
// HandlerFunc: The function to handle the request, this basicically should just be the proxy call, but it allows you to specify more specific things
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// RouteCollection is a slice of routes that is used to represent a service (may change name here )
// Usage: The recomendation is to create a RouteCollection variable for all of you services and for each service create a specific one then in a registration function append all the service collections into the single master collection.
type RouteCollection []Route

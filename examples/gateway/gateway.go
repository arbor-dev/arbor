package gateway

import (
	"fmt"
	"net/http"
	"github.com/arbor-dev/arbor"
	"github.com/arbor-dev/arbor/logger"
	"github.com/arbor-dev/arbor/proxy"
	"github.com/arbor-dev/arbor/security"
)

//The Global Collection of Routes, not necessary but good for organization
var routes = arbor.RouteCollection{
	//Route on the Arbor Server (No backing service)
	arbor.Route{
		Name:    "Index",
		Method:  "GET",
		Pattern: "/",
		Handler: index,
	},
}

//Handler for the Index Route
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

//Create the full collection of Routes
func RegisterRoutes() arbor.RouteCollection {
	routes = append(routes, productServiceRoutes...)
	return routes
}

//Arbor configurations
func ConfigArbor() {
	//Location for the Access Log
	security.AccessLogLocation = "/tmp/arbor_access.log"
	//Location for the Client Registry
	security.ClientRegistryLocation = "/tmp/arbor_clients.db"
	//Access Control for the Proxy
	proxy.AccessControlPolicy = "*"
	//Logging Verboseness
	logger.LogLevel = logger.DEBUG
}

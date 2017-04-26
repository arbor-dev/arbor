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

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/acm-uiuc/arbor/security"
	"github.com/acm-uiuc/arbor/server"
)

// Boot is a standard server CLI
//
// Provide a set of routes to serve and a port to serve on.
//
// 	executable [-r | --register-client client_name] [-c | --check-registration token] [-u | --unsecured]
//
// 	-r | --register-client client_name
// registers a client, generates a token
//
// 	-c | --check-registration token
// checks if a token is valid and returns name of client
//
// 	-u | --unsecured
// runs groot without the security layer
//
// 	without args
// runs groot with the security layer
//
// It will start the arbor instance, parsing the command arguments and execute the behavior.
func Boot(routes RouteCollection, port uint16) {
	if len(os.Args) == 3 && (os.Args[1] == "--register-client" || os.Args[1] == "-r") {
		RegisterClient(os.Args[2])
	} else if len(os.Args) == 3 && (os.Args[1] == "--check-registration" || os.Args[1] == "-c") {
		CheckRegistration(os.Args[2])
	} else if len(os.Args) == 2 && (os.Args[1] == "--unsecured" || os.Args[1] == "-u") {
		StartUnsecuredServer(routes, port)
	} else if len(os.Args) > 1 {
		fmt.Println("Invalid Command")
	} else {
		StartServer(routes, port)
	}
}

// RegisterClient will generate a access token for a client
//
// Currently uses a db of client names.
func RegisterClient(name string) {
	security.Init()
	token, err := security.AddClient(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Client " + name + " has been granted authorization token: " + token)

	defer security.Shutdown()
}

// CheckRegistration allows you to check what client was assigned to a particular token
func CheckRegistration(token string) {
	security.Init()
	fmt.Println(security.IsAuthorizedClient(token))
	defer security.Shutdown()
}

// StartServer starts a secured arbor server (Token required for access)
//
// Provide a set of routes to serve and a port to serve on.
func StartServer(routes RouteCollection, port uint16) {

	security.Init()
	router := server.NewRouter(routes.toServiceRoutes())

	log.Println("ROOTS BEING PLANTED [Server is listening on :" + fmt.Sprintf("%d", port) + "]")
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", port), router))

	defer security.Shutdown()
}

// StartUnsecuredServer starts an unsecured arbor server (Token required for access)
//
// Provide a set of routes to server and a port to serve on/
func StartUnsecuredServer(routes RouteCollection, port uint16) {
	router := server.NewRouter(routes.toServiceRoutes())

	log.Println("ROOTS BEING PLANTED [Server is listening on :" + fmt.Sprintf("%d", port) + "]")
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", port), router))
}

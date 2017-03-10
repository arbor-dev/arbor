/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/acm-uiuc/arbor/security"
	"github.com/acm-uiuc/arbor/services"
)

func Boot(routes services.RouteCollection) {
	if len(os.Args) == 3 && (os.Args[1] == "--register-client" || os.Args[1] == "-r") {
		RegisterClient(os.Args[2])
	} else if len(os.Args) == 3 && (os.Args[1] == "--check-registration" || os.Args[1] == "-c") {
		CheckRegistration(os.Args[2])
	} else if len(os.Args) == 2 && (os.Args[1] == "--unsecured" || os.Args[1] == "-u") {
		StartUnsecuredServer(routes)
	} else if len(os.Args) > 1 {
		fmt.Println("Invalid Command")
	} else {
		StartServer(routes)
	}
}

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

func CheckRegistration(token string) {
	security.Init()
	fmt.Println(security.IsAuthorizedClient(token))
	defer security.Shutdown()
}

func StartServer(routes services.RouteCollection) {

	security.Init()
	router := NewRouter(routes)

	log.Println("[INSERT FUNNY TREE JOKE HERE]! [Server is listening on :8000]")
	log.Fatal(http.ListenAndServe(":8000", router))

	defer security.Shutdown()
}

func StartUnsecuredServer(routes services.RouteCollection) {
	router := NewRouter(routes)

	log.Println("I AM GROOT! [Server is listening on :8000]")
	log.Fatal(http.ListenAndServe(":8000", router))
}

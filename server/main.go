/**
* Copyright © 2016, ACM@UIUC
*
* This file is part of the Groot Project.  
* 
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"github.com/acm-uiuc/groot/services"
	"github.com/acm-uiuc/groot/security"
)

func main() {
	if len(os.Args) == 3 && (os.Args[1] == "--register-client" || os.Args[1] == "-r") {
		RegisterClient(os.Args[2])
	} else if len(os.Args) == 3 && (os.Args[1] == "--check-registration" || os.Args[1] == "-c") {
		CheckRegistration(os.Args[2])
	} else if len(os.Args) == 2 && (os.Args[1] == "--unsecured" || os.Args[1] == "-u") {
		StartUnsecuredServer() 
	} else if len(os.Args) > 1 {
		fmt.Println("Invalid Command")
	} else {
		StartServer()
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

	defer security.Close()
}

func CheckRegistration(token string) {
	security.Init()
	fmt.Println(security.IsAuthorizedClient(token))
	defer security.Close()
}

func StartServer() {

	security.Init()
	services.RegisterAPIs()
	router := NewRouter()

	log.Println("I AM GROOT! [Server is listening on :8000]")
	log.Fatal(http.ListenAndServe(":8000", router))

	defer security.Close()
}

func StartUnsecuredServer() {
	services.RegisterAPIs()
	router := NewRouter()

	log.Println("I AM GROOT! [Server is listening on :8000]")
	log.Fatal(http.ListenAndServe(":8000", router))
}
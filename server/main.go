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
	"log"
	"net/http"
	"github.com/acm-uiuc/groot/services"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	services.RegisterAPIs()
	router := NewRouter()

	log.Println("I AM GROOT! [Server is listening on :8000]")
	log.Fatal(http.ListenAndServe(":8000", router))
}

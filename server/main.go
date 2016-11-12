package main

import (
	"log"
	"net/http"
	"github.com/acm-uiuc/groot/services"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	services.RegisterAPIs()

	log.Println(services.Routes)

	router := NewRouter()

	log.Println("I AM GROOT! [Server is listening on :8000]")
	log.Fatal(http.ListenAndServe(":8000", router))
}

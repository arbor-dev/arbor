package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	RegisterAPIs()

	router := NewRouter()

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

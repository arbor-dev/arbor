package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	//"strconv"

)

var JSONHeader string = "application/json; charset=UTF-8"


func Index(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!\n")
}

func GetHandler(w http.ResponseWriter, url string, r *http.Request) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("%+v\n", response)

	// w.Header().Set("Content-Type", JSONHeader)
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(todos); err != nil {
	// 	panic(err)
	// }
	}
}

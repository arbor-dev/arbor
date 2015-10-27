package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var JSONHeader string = "application/json; charset=UTF-8"
var TEXTHeader string = "text/plain; charset=utf-8"

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetHandler(w http.ResponseWriter, url string, r *http.Request) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		if contains(TEXTHeader, response.Header["Content-Type"]) {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "%s\n", string(contents))
		}

		if contains(JSONHeader, response.Header["Content-Type"]) {
			w.Header().Set("Content-Type", JSONHeader)
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response.Body); err != nil {
				panic(err)
			}
		}
	}
}

package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var JSONHeader string = "application/json; charset=UTF-8"
var TEXTHeader string = "text/plain; charset=utf-8"

func GETHandler(w http.ResponseWriter, url string, r *http.Request) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	switch {

	case contains(TEXTHeader, response.Header["Content-Type"]):
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s\n", string(contents))
		break

	case contains(JSONHeader, response.Header["Content-Type"]):
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		var data interface{}
		err = json.Unmarshal(contents, &data)
		if err != nil {
			panic(err)
		}

		fmt.Println(data)

		w.Header().Set("Content-Type", JSONHeader)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
		break
	}
}

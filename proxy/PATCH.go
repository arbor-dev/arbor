package proxy

import (
	//"bytes"
	"encoding/json"
    //"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func PATCH(w http.ResponseWriter, url string, r *http.Request) {
	content, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		InvalidPUT(w, err)
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		InvalidPUT(w, err)
		log.Println(err)
		return
	}

	log.Println(content)
}

func InvalidPATCH(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 400, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

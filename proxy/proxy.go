package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var JSONHeader string = "application/json; charset=UTF-8"
var TEXTHeader string = "text/plain; charset=utf-8"

func GETHandler(w http.ResponseWriter, url string, r *http.Request) {
	res, err := http.Get(url)
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	switch {

	case contains(TEXTHeader, res.Header["Content-Type"]):
		if err != nil {
			InvalidGET(w, err)
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "%s\n", string(contents))
		break

	case contains(JSONHeader, res.Header["Content-Type"]):
		var data interface{}
		err = json.Unmarshal(contents, &data)
		if err != nil {
			InvalidGET(w, err)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", JSONHeader)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			InvalidGET(w, err)
			log.Println(err)
			return
		}
		break
	}
}

func POSTHandler(w http.ResponseWriter, url string, r *http.Request) {

	content, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", JSONHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var data interface{}
	err = json.Unmarshal(contents, &data)
	if err != nil {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusCreated)
}

func InvalidPOST(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 422, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
func InvalidGET(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

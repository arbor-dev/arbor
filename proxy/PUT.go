package proxy

import (
	"bytes"
	"encoding/json"
	//"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func PUT(w http.ResponseWriter, url string, r *http.Request) {
	content, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		InvalidPUT(w, err)
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed Reception:%v", err)
		return
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", JSONHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed request: %v", err)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var data interface{}
	err = json.Unmarshal(contents, &data)
	log.Println(data)
	if err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed decode %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		InvalidPUT(w, err)
		log.Printf("Failed server %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed encode %v", err)
		return
	}
	w.Header().Set("Content-Type", JSONHeader)
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(http.StatusOK)
}

func InvalidPUT(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 400, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

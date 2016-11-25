package proxy

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func PUT(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	if format != "XML" && format != "JSON" { //TODO: Support Post form data
		err := errors.New("ERROR: unsupported data encoding")
		InvalidPUT(w, err)
		log.Println(err)
		return
	}
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

	var data interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}

	origin := r.Header.Get("Origin")

	//TODO: FIGURE OUT ORIGIN RULES
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	    w.Header().Set("Access-Control-Allow-Methods", "PUT")
	    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}


	switch format {
		case "XML":
			xmlPUT(w, url, data)
			return
		case "JSON":
			jsonPUT(w, url, data)
			return
		default:
			InvalidPUT(w, err)
			log.Println("Unsupported Data Encoding")
			return
	}
}

func InvalidPUT(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 400, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
func jsonPUT(w http.ResponseWriter, url string, data interface{}) {
	content, err := json.Marshal(data)
	if err != nil {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", JSONHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK  {
		InvalidPUT(w, err)
		log.Printf("Failed request: %v", err)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var serverData interface{}
	err = json.Unmarshal(contents, &serverData)
	if err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed decode %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed encode %v", err)
		return
	}
	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusOK)
}

func xmlPUT(w http.ResponseWriter, url string, data interface{}) {
	content, err := xml.Marshal(data)
	if err != nil {
		InvalidPOST(w, err)
		log.Println(err)
		return
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", XMLHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK  {
		InvalidPUT(w, err)
		log.Printf("Failed request: %v", err)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var serverData interface{}
	err = xml.Unmarshal(contents, &serverData)
	if err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed decode %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		InvalidPUT(w, err)
		log.Printf("Failed encode %v", err)
		return
	}
	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusOK)
}

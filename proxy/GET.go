/**
* Copyright © 2016, ACM@UIUC
*
* This file is part of the Groot Project.  
* 
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package proxy

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GET(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	sanitizeRequest(r)

	if !verifyAuthorization(r) {
		w.WriteHeader(403)
		log.Println("Unauthorized Access from " + r.RemoteAddr)
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if format == "JSON" {
		req.Header.Set("Content-Type", JSONHeader)
		req.Header.Set("Accept", "application/json")
	}

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil  || res.StatusCode != http.StatusOK {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	origin := r.Header.Get("Origin")

	//TODO: FIGURE OUT ORIGIN RULES
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	    w.Header().Set("Access-Control-Allow-Methods", "GET")
	    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	switch {
		case contains(JSONHeader, res.Header["Content-Type"]):
			jsonGET(w, url, contents)
			break
		case contains(TEXTHeader, res.Header["Content-Type"]):
			textGET(w, url, contents)
			break
		case contains(HTMLHeader, res.Header["Content-Type"]):
			textGET(w, url, contents)
			break
		case contains(XMLHeader, res.Header["Content-Type"]):
			xmlGET(w, url, contents)
			break
		default:
			if err != nil {
				InvalidGET(w, err)
				log.Println(err)
				return
			}
			textGET(w, url, contents)
			break
	}
}

func jsonGET(w http.ResponseWriter, url string, contents []byte) {
	var data interface{}
	err := json.Unmarshal(contents, &data)
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
}

func xmlGET(w http.ResponseWriter, url string, contents []byte) {
	var data interface{}
	err := xml.Unmarshal(contents, &data)
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", XMLHeader)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
}

func textGET(w http.ResponseWriter, url string, contents []byte) {
	var err error
	if err != nil {
		InvalidGET(w, err)
		log.Println(err)
		return
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}

func InvalidGET(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "server-err": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

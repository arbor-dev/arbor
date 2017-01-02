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
	"github.com/acm-uiuc/groot/config"
	"io/ioutil"
	"log"
	"net/http"
)

func DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	if !verifyAuthorization(r) {
		w.WriteHeader(403)
		log.Println("Unauthorized Access from " + r.RemoteAddr)
		return
	}

	req, err := http.NewRequest("DELETE", url, nil)
	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}

	if token != "" {
		req.Header.Set("Authorization", config.AuthPrefix + token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		InvalidDELETE(w, err)
		log.Printf("Failed server %v", err)
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
	
	origin := r.Header.Get("Origin")

	//TODO: FIGURE OUT ORIGIN RULES
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	    w.Header().Set("Access-Control-Allow-Methods", "GET")
	    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

func InvalidDELETE(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

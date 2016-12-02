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
	"log"
	"net/http"
)

func DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	req, err := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		InvalidDELETE(w, err)
		log.Printf("Failed server %v", err)
		return
	}

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

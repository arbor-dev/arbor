/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package proxy

import (
	//"bytes"
	"encoding/json"
	//"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/arbor-dev/arbor/logger"
)

// PATCH provides a proxy PATCH request allowing authorized clients to make PATHC
// requests of the microservices
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will call the service and return the result to the client.
func PATCH(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	preprocessing_err := requestPreprocessing(w, r)
	if preprocessing_err != nil {
		return
	}

	_, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxRequestSize))
	if err != nil {
		invalidPATCH(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		invalidPATCH(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	origin := r.Header.Get("Origin")

	//TODO: FIGURE OUT ORIGIN RULES
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

func invalidPATCH(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 400, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

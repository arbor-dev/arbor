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
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/arbor-dev/arbor/logger"
)

func POST(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	sanitizeRequest(r)

	if !verifyAuthorization(r) {
		w.WriteHeader(403)
		logger.Log(logger.WARN, "Attempted unauthorized Access from "+r.RemoteAddr)
		return
	}

	if format != "XML" && format != "JSON" { //TODO: Support Post form data
		err := errors.New("ERROR: unsupported data encoding")
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	content, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	var data interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	origin := r.Header.Get("Origin")

	//TODO: FIGURE OUT ORIGIN RULES
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	switch format {
	case "XML":
		xmlPOST(r, w, url, token, data)
		return
	case "JSON":
		jsonPOST(r, w, url, token, data)
		return
	default:
		invalidPOST(w, err)
		logger.Log(logger.ERR, "Unsupported Data Encoding")
		return
	}
}

func jsonPOST(r *http.Request, w http.ResponseWriter, url string, token string, data interface{}) {
	content, err := json.Marshal(data)
	if err != nil {
		logger.Log(logger.ERR, err.Error())
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", JSONHeader)
	req.Header.Set("Accept", "application/json")

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{Timeout: time.Duration(Timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		if resp != nil {
			log.Println(resp.StatusCode)
		}
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	} else if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		logger.Log(logger.WARN, "SERVICE FAILED - SERVICE RETURNED STATUS "+http.StatusText(resp.StatusCode))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(resp.StatusCode)
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, "Failed to read response")
		return
	}

	var serverData interface{}
	err = json.Unmarshal(contents, &serverData)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, "Failed to unmarshal json "+err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	w.Header().Set("Content-Type", JSONHeader)
	//NOTE: Apparently not needed but add back in if things break
	//w.WriteHeader(http.StatusCreated)
}

func xmlPOST(r *http.Request, w http.ResponseWriter, url string, token string, data interface{}) {
	content, err := xml.Marshal(data)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", XMLHeader)

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{Timeout: time.Duration(Timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusCreated {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var serverData interface{}
	err = xml.Unmarshal(contents, &serverData)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusCreated)
}

func invalidPOST(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 422, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

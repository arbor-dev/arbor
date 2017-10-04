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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/acm-uiuc/arbor/logger"
)

func PUT(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	if !verifyAuthorization(r) {
		w.WriteHeader(403)
		logger.Log(logger.WARN, "Attempted unauthorized Access Attempt from "+r.RemoteAddr)
		return
	}

	if format != "XML" && format != "JSON" { //TODO: Support Post form data
		err := errors.New("Unsupported data encoding")
		invalidPUT(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	content, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		invalidPUT(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		invalidPUT(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed Reception:%v", err))
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
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	switch format {
	case "XML":
		xmlPUT(r, w, url, token, data)
		return
	case "JSON":
		jsonPUT(r, w, url, token, data)
		return
	default:
		invalidPUT(w, err)
		logger.Log(logger.ERR, "Unsupported Data Encoding")
		return
	}
}

func invalidPUT(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	data := map[string]interface{}{"Code": 400, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Log(logger.ERR, err.Error())
		panic(err) //ASK: SHOULD THIS BE HERE?
	}
}

func jsonPUT(r *http.Request, w http.ResponseWriter, url string, token string, data interface{}) {
	content, err := json.Marshal(data)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(content))

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}
	req.Header.Set("Content-Type", JSONHeader)
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{Timeout: time.Duration(Timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		logger.Log(logger.ERR, "SERVICE FAILED - SERVICE RETURNED STATUS "+http.StatusText(resp.StatusCode))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var serverData interface{}
	err = json.Unmarshal(contents, &serverData)
	if err != nil {
		invalidPUT(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed to decode:%v", err))
		return
	}

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		invalidPUT(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed to encode:%v", err))
		return
	}
	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusOK)
}

func xmlPUT(r *http.Request, w http.ResponseWriter, url string, token string, data interface{}) {
	content, err := xml.Marshal(data)
	if err != nil {
		invalidPUT(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(content))
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
	if err != nil || resp.StatusCode != http.StatusOK {
		invalidPUT(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed request:%v", err))
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var serverData interface{}
	err = xml.Unmarshal(contents, &serverData)
	if err != nil {
		invalidPUT(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed decode:%v", err))
		return
	}

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		invalidPUT(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed encode:%v", err))
		return
	}
	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusOK)
}

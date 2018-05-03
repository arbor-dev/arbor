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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arbor-dev/arbor/logger"
)

func DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	preprocessing_err := requestPreprocessing(w, r)
	if preprocessing_err != nil {
		return
	}

	req, err := http.NewRequest("DELETE", url, nil)
	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}

	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{Timeout: time.Duration(Timeout) * time.Second}
	resp, err := client.Do(req)

	logger.LogResp(logger.DEBUG, resp)

	if err != nil {
		invalidDELETE(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Hit %s", err.Error()))
		return
	} else if resp.StatusCode != http.StatusOK {
		logger.Log(logger.WARN, "SERVER RETURNED STATUS " + http.StatusText(resp.StatusCode))
		w.WriteHeader(resp.StatusCode)
		return
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	var serverData interface{}
	err = json.Unmarshal(contents, &serverData)
	if err != nil {
		invalidDELETE(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed decode %v", err))
		return
	}

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		invalidDELETE(w, err)
		logger.Log(logger.ERR, fmt.Sprintf("Failed encode %v", err))
		return
	}

	w.Header().Set("Content-Type", JSONHeader)
	//IF THINGS START BREAKING UNCOMMENT
	//w.WriteHeader(http.StatusOK)
}

func invalidDELETE(w http.ResponseWriter, err error) {
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", AccessControlPolicy)
	w.WriteHeader(http.StatusNotFound)
	data := map[string]interface{}{"Code": http.StatusNotFound, "Text": "Not Found", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

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
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arbor-dev/arbor/logger"
)

// GET provides a proxy GET request allowing authorized clients to make GET
//
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
func GET(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	preprocessing_err := requestPreprocessing(w, r)
	if preprocessing_err != nil {
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

	client := &http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Do(req)

	logger.LogResp(logger.DEBUG, res)

	origin := r.Header.Get("Origin")

	if err != nil {
		// For an error in making the request
		// Log the error, but return the output from the service.
		logger.Log(logger.ERR, err.Error())
		notifyClientOfRequestError(w, http.StatusInternalServerError, "")
		return
	} else if res.StatusCode == http.StatusFound {
		// For redirects
		logger.Log(logger.DEBUG, "Service Returned Redirect")
		w.Header().Set("Location", res.Header.Get("Location"))
		w.WriteHeader(http.StatusFound)
		return
	} else if res.StatusCode != http.StatusOK {
		// For non-200 errors
		contents, readErr := ioutil.ReadAll(res.Body)

		if readErr != nil {
			// 503 Bad Gateway, indicating a proxy / gateway recieved an invalid response from upstream server.
			notifyClientOfRequestError(w, http.StatusBadGateway, "")
			return
		}

		logger.Log(logger.WARN, "SERVICE RETURNED STATUS "+http.StatusText(res.StatusCode))

		w.WriteHeader(res.StatusCode)
		switch {
		case contains(JSONHeader, res.Header["Content-Type"]):
			invalidJsonGET(w, contents, err)
			return
		case contains(XMLHeader, res.Header["Content-Type"]):
			invalidXmlGET(w, contents, err)
			return
		default:
			invalidTextGET(w, contents, err)
			return
		}
	}

	defer res.Body.Close()

	contents, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		// 503 Bad Gateway, indicating a proxy / gateway recieved an invalid response from upstream server.
		notifyClientOfRequestError(w, http.StatusBadGateway, "")
		return
	}

	//TODO: FIGURE OUT ORIGIN RULES
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	switch {
	case contains(JSONHeader, res.Header["Content-Type"]):
		jsonGET(w, url, contents)
		return
	case contains(TEXTHeader, res.Header["Content-Type"]):
		textGET(w, url, contents)
		return
	case contains(HTMLHeader, res.Header["Content-Type"]):
		textGET(w, url, contents)
		return
	case contains(XMLHeader, res.Header["Content-Type"]):
		xmlGET(w, url, contents)
		return
	default:
		if err != nil {
			invalidTextGET(w, contents, err)
			logger.Log(logger.ERR, err.Error())
			return
		}
		textGET(w, url, contents)
		return
	}
}

func jsonGET(w http.ResponseWriter, url string, contents []byte) {
	var data interface{}
	err := json.Unmarshal(contents, &data)
	if err != nil {
		invalidJsonGET(w, contents, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		invalidJsonGET(w, contents, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
}

func xmlGET(w http.ResponseWriter, url string, contents []byte) {
	var data interface{}
	err := xml.Unmarshal(contents, &data)
	if err != nil {
		invalidXmlGET(w, contents, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	w.Header().Set("Content-Type", XMLHeader)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		invalidXmlGET(w, contents, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
}

func textGET(w http.ResponseWriter, url string, contents []byte) {
	var err error
	if err != nil {
		invalidTextGET(w, contents, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}

func invalidJsonGET(w http.ResponseWriter, contents []byte, err error) {
	w.Header().Set("Content-Type", JSONHeader)

	var data interface{}
	err = json.Unmarshal(contents, &data)

	if err != nil {
		notifyClientOfRequestError(w, http.StatusBadGateway, "")
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		notifyClientOfRequestError(w, http.StatusBadGateway, "")
		return
	}
}

func invalidXmlGET(w http.ResponseWriter, contents []byte, err error) {
	w.Header().Set("Content-Type", XMLHeader)

	var data interface{}
	err = xml.Unmarshal(contents, &data)

	if err != nil {
		notifyClientOfRequestError(w, http.StatusBadGateway, "")
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		notifyClientOfRequestError(w, http.StatusBadGateway, "")
		return
	}
}

func invalidTextGET(w http.ResponseWriter, contents []byte, err error) {
	w.Write(contents)
}

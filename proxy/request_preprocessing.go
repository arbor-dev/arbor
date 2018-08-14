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
	"fmt"
	"net/http"

	"github.com/arbor-dev/arbor/logger"
	"github.com/arbor-dev/arbor/security"
)

func verifyAuthorization(authorization string, remoteAddr string) bool {
	//IsAuthorizedClient Handles empty token
	auth, err := security.IsAuthorizedClient(authorization)
	if err != nil {
		logger.Log(logger.WARN, "Attempted unauthorized access from "+remoteAddr)
		return false
	}
	return auth
}

func sanitizeRequest(r *http.Request) {
	security.SanitizeRequest(r)
}

type preprocessingError struct {
	arg  int
	prob string
}

func (e *preprocessingError) Error() string {
	return fmt.Sprintf("%s", e.prob)
}

func requestPreprocessing(w http.ResponseWriter, r *http.Request) error {
	logger.LogReq(logger.DEBUG, r)
	sanitizeRequest(r)
	if !verifyAuthorization(r.Header.Get(ClientAuthorizationHeaderField), r.RemoteAddr) {
		w.WriteHeader(http.StatusForbidden)
		return &preprocessingError{-1, "Client Not Authorized"}
	}
	return nil
}

/*
	Called if an error is caused while Marshalling / Unmarshalling data, or making proxy requests.
	If the message is left blank, and the error code is one of the supported ones, the error
	message will be pre-filled.
*/
func notifyClientOfRequestError(w http.ResponseWriter, httpStatusCode int, message string) {
	if message == "" {
		switch httpStatusCode {
		case http.StatusBadGateway:
			message = "The API Gateway received an invalid response."
		case http.StatusInternalServerError:
			message = "The API Gateway encountered an error while making the proxy request."
		default:
			message = "Please check the API Gateway logs for more details on the error."
		}
	}
	w.WriteHeader(httpStatusCode)
	fmt.Fprintf(w, "%s\n", message)
	logger.Log(logger.ERR, message)
}

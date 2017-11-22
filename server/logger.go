/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/acm-uiuc/arbor/logger"
)

type StatusResponseWriter struct {
    http.ResponseWriter
    status int
}

func (rec *StatusResponseWriter) WriteHeader(code int) {
    rec.status = code
    rec.ResponseWriter.WriteHeader(code)
}

func logRequest(method string, requestURI string, routeName string, responseStatus int, latency time.Duration) {
    logger.Log(logger.INFO, fmt.Sprintf("%s\t%s\t%s\t%d\t%s", method, requestURI, routeName, responseStatus, latency))
}

func httpLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
        s := &StatusResponseWriter{ResponseWriter: w, status: 200}
		inner.ServeHTTP(s, r)
        logRequest(r.Method, r.RequestURI, name, s.status, time.Since(start))
	})
}

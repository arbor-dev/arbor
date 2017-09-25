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

func httpLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Log(logger.INFO, fmt.Sprintf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start)))
		inner.ServeHTTP(w, r)
	})
}

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
	"net/http"
)

// Timeout is the default request timeout
var Timeout int64 = 10

// AccessControlPolicy is the default Access control policy
var AccessControlPolicy = "*"

// Client Authorization Token Field
var ClientAuthorizationHeaderField = "Authorization"

// Defines the maximum size for requests
const (
	MB = 1048576
	MaxRequestSize = 1 * MB
	MaxFileUploadSize = 16 * MB
)

var DefaultProxyRequestSettings = ProxyRequestSettings{
	ErrorHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}),
	RequestMiddlewares: nil,
	ResponseMiddlewares: nil,
}

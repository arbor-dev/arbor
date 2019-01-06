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

// AccessControlPolicy is the default Access control policy
var AccessControlPolicy = "*"

// ProxyMiddlewares is the default error handler and middlewares to use when proxying a request
var ProxyMiddlewares = RequestSettings{
	ErrorHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}),
	RequestMiddlewares: nil,
	ResponseMiddlewares: nil,
}

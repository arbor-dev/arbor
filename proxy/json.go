package proxy

import (
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"encoding/json"
)

// A handler for writing errors into the response sent to the caller
var JsonErrorHandler = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
	w.WriteHeader(500)
})

// A handler which validates the request body for valid json
var jsonValidator = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxRequestSize))

	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err != nil {
		JsonErrorHandler.ServeHTTP(w, r)
	}

	if len(body) == 0 {
		return
	}

	var jsonBody interface{}
	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		JsonErrorHandler.ServeHTTP(w, r)
	}
})

// A set of middlewares for validating json in the request to a service
var JsonRequestMiddlewares = []http.Handler{
	jsonValidator,
}

// A set of middlewares for validating json in the response from a service
var JsonResponseMiddlewares = []http.Handler{
	jsonValidator,
}

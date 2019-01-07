package middleware

import (
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"encoding/json"

	"github.com/arbor-dev/arbor/proxy/constants"
)

// JSONErrorHandler is the handler for writing errors into the response sent to the caller
var JSONErrorHandler = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
	w.WriteHeader(500)
})

// A handler which validates the request body for valid json
var jsonValidator = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, constants.MaxRequestSize))

	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err != nil {
		JSONErrorHandler.ServeHTTP(w, r)
	}

	if len(body) == 0 {
		return
	}

	var jsonBody interface{}
	err = json.Unmarshal(body, &jsonBody)

	if err != nil {
		JSONErrorHandler.ServeHTTP(w, r)
	}
})

// JSONRequestMiddlewares is the set of middlewares for validating json in the request to a service
var JSONRequestMiddlewares = []http.Handler{
	jsonValidator,
}

// JSONResponseMiddlewares is the set of middlewares for validating json in the response from a service
var JSONResponseMiddlewares = []http.Handler{
	jsonValidator,
}

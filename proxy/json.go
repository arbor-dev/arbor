package proxy

import (
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"encoding/json"
)

var JsonErrorHandler http.Handler
var JsonRequestMiddlewares []http.Handler
var JsonResponseMiddlewares []http.Handler

func init() {
	JsonErrorHandler = http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
		w.WriteHeader(500)
	})

	jsonValidator := http.HandlerFunc(func(w http.ResponseWriter, r* http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxRequestSize))

		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if err != nil {
			JsonErrorHandler.ServeHTTP(w, r)
		}

		var jsonBody interface{}
		err = json.Unmarshal(body, &jsonBody)

		if err != nil {
			JsonErrorHandler.ServeHTTP(w, r)
		}
	})

	JsonRequestMiddlewares = []http.Handler{
		jsonValidator,
	}

	JsonResponseMiddlewares = []http.Handler{
		jsonValidator,
	}
}

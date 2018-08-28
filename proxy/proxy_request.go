package proxy

import (
	"net/http"
	"io/ioutil"
	"io"
	"time"
	"bytes"
)

type ProxyRequestSettings struct {
	ErrorHandler        http.Handler
	RequestMiddlewares  []http.Handler
	ResponseMiddlewares []http.Handler
}

func proxyRequest(w http.ResponseWriter, r *http.Request, url string, settings ProxyRequestSettings) {
	requestBody, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxFileUploadSize))

	if err != nil {
		settings.ErrorHandler.ServeHTTP(w, r)
		return
	}

	err = r.Body.Close()

	if err != nil {
		settings.ErrorHandler.ServeHTTP(w, r)
		return
	}

	req, err := http.NewRequest(r.Method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		settings.ErrorHandler.ServeHTTP(w, r)
		return
	}

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}

	client := &http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)

	if err != nil {
		settings.ErrorHandler.ServeHTTP(w, r)
		return
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		settings.ErrorHandler.ServeHTTP(w, r)
		return
	}

	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = w.Write(responseBody)

	if err != nil {
		settings.ErrorHandler.ServeHTTP(w, r)
		return
	}
}

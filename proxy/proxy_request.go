package proxy

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/arbor-dev/arbor/proxy/constants"
	"github.com/koding/websocketproxy"
)

// MiddlewareSet contains the error handler and middlewares to use when proxying a request
type MiddlewareSet struct {
	ErrorHandler        http.Handler
	RequestMiddlewares  []http.Handler
	ResponseMiddlewares []http.Handler
}

// ProxyWebsocketWithMiddlewares proxies the provided request using the given middlewares and upgrades the request to a websocket
func ProxyWebsocketWithMiddlewares(w http.ResponseWriter, r *http.Request, url_str string, proxyMiddlewares MiddlewareSet) {
	for _, requestMiddleware := range proxyMiddlewares.RequestMiddlewares {
		requestMiddleware.ServeHTTP(w, r)
	}

	u, err := url.Parse(url_str)

	fmt.Println(url_str)

	if err != nil {
		proxyMiddlewares.ErrorHandler.ServeHTTP(w, r)
		return
	}

	websocketproxy.NewProxy(u).ServeHTTP(w, r)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func copyResponse(rw http.ResponseWriter, resp *http.Response) error {
	copyHeader(rw.Header(), resp.Header)
	rw.WriteHeader(resp.StatusCode)
	defer resp.Body.Close()

	_, err := io.Copy(rw, resp.Body)
	return err
}

// ProxyRequestWithMiddlewares proxies the provided request using the given middlewares
func ProxyRequestWithMiddlewares(w http.ResponseWriter, r *http.Request, url string, proxyMiddlewares MiddlewareSet) {
	for _, requestMiddleware := range proxyMiddlewares.RequestMiddlewares {
		requestMiddleware.ServeHTTP(w, r)
	}

	requestBody, err := ioutil.ReadAll(io.LimitReader(r.Body, constants.MaxFileUploadSize))

	if err != nil {
		proxyMiddlewares.ErrorHandler.ServeHTTP(w, r)
		return
	}

	err = r.Body.Close()

	if err != nil {
		proxyMiddlewares.ErrorHandler.ServeHTTP(w, r)
		return
	}

	req, err := http.NewRequest(r.Method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		proxyMiddlewares.ErrorHandler.ServeHTTP(w, r)
		return
	}

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}

	client := &http.Client{
		Timeout: time.Duration(constants.Timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)

	if err != nil {
		proxyMiddlewares.ErrorHandler.ServeHTTP(w, r)
		return
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		proxyMiddlewares.ErrorHandler.ServeHTTP(w, r)
		return
	}

	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	for _, responseMiddleware := range proxyMiddlewares.ResponseMiddlewares {
		responseMiddleware.ServeHTTP(w, r)
	}

	w.WriteHeader(resp.StatusCode)

	_, err = w.Write(responseBody)

	if err != nil {
		proxyMiddlewares.ErrorHandler.ServeHTTP(w, r)
		return
	}
}

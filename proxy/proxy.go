package proxy

import (
	"net/http"
	"github.com/arbor-dev/arbor/proxy/middleware"
)

// GET proxies a GET request
func GET(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, format, token)
}

// POST proxies a POST request
func POST(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, format, token)
}

// PUT proxies a PUT request
func PUT(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, format, token)
}

// DELETE proxies a DELETE request
func DELETE(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, format, token)
}

// PATCH proxies a PATCH request
func PATCH(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, format, token)
}

// ProxyRequest proxies the caller's request based on the url, format, and token
func ProxyRequest(w http.ResponseWriter, r* http.Request, url string, format string, token string) {
	middlewares := ProxyMiddlewaresFactory(format, token)

	ProxyRequestWithMiddlewares(w, r, url, middlewares)
}

// ProxyMiddlewaresFactory a set of middlewares based on the provided format and token
func ProxyMiddlewaresFactory(format string, token string) MiddlewareSet {
	middlewares := ProxyMiddlewares

	middlewares.RequestMiddlewares = append(middlewares.RequestMiddlewares, middleware.PreprocessingMiddleware)
	middlewares.RequestMiddlewares = append(middlewares.RequestMiddlewares, middleware.TokenMiddlewareFactory(token))

	middlewares.ResponseMiddlewares = append(middlewares.ResponseMiddlewares, middleware.CORSMiddleware)

	switch format {
	case "JSON":
		middlewares.ErrorHandler = middleware.JSONErrorHandler
		middlewares.RequestMiddlewares = append(middlewares.RequestMiddlewares, middleware.JSONRequestMiddlewares...)
		middlewares.ResponseMiddlewares = append(middlewares.ResponseMiddlewares, middleware.JSONResponseMiddlewares...)
	case "RAW":
		fallthrough
	default:
	}

	return middlewares
}

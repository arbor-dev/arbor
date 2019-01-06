package proxy

import (
	"net/http"
	"github.com/arbor-dev/arbor/proxy/middleware"
)

// GET proxies a GET request
func GET(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

// POST proxies a POST request
func POST(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

// PUT proxies a PUT request
func PUT(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

// DELETE proxies a DELETE request
func DELETE(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

// PATCH proxies a PATCH request
func PATCH(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

// Proxy the caller's request to the correct service with proxy request settings
// Settings contain the error handler, request middlewares, and response middlewares
func proxyRequestWithSettings(w http.ResponseWriter, r* http.Request, url string, format string, token string) {
	settings := ProxyMiddlewares

	settings.RequestMiddlewares = append(settings.RequestMiddlewares, middleware.PreprocessingMiddleware)
	settings.RequestMiddlewares = append(settings.RequestMiddlewares, middleware.TokenMiddlewareFactory(token))

	settings.ResponseMiddlewares = append(settings.ResponseMiddlewares, middleware.CORSMiddleware)

	switch format {
	case "JSON":
		settings.ErrorHandler = middleware.JsonErrorHandler
		settings.RequestMiddlewares = append(settings.RequestMiddlewares, middleware.JsonRequestMiddlewares...)
		settings.ResponseMiddlewares = append(settings.ResponseMiddlewares, middleware.JsonResponseMiddlewares...)
	case "RAW":
		fallthrough
	default:
	}

	proxyRequest(w, r, url, settings)
}

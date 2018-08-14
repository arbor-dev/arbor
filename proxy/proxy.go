package proxy

import (
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, DefaultProxyRequestSettings)
}

func POST(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, DefaultProxyRequestSettings)
}

func PUT(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, DefaultProxyRequestSettings)
}

func DELETE(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, DefaultProxyRequestSettings)
}

func PATCH(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	ProxyRequest(w, r, url, DefaultProxyRequestSettings)
}

package arbor

import (
	"net/http"

	"github.com/acm-uiuc/arbor/proxy"
)

func DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.DELETE(w, url, format, token, r)
}

func GET(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.GET(w, url, format, token, r)
}

func PATCH(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.PATCH(w, url, format, token, r)
}

func POST(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.POST(w, url, format, token, r)
}

func PUT(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.PUT(w, url, format, token, r)
}

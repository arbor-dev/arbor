package arbor

import (
	"net/http"

	"github.com/acm-uiuc/arbor/proxy"
)

// DELETE provides a proxy DELETE request allowing authorized clients to make DELETE requests of the microservices
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will call the service and return the result to the client.
func DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.DELETE(w, url, format, token, r)
}

// GET provides a proxy GET request allowing authorized clients to make GET requests of the microservices
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will call the service and return the result to the client.
func GET(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.GET(w, url, format, token, r)
}

// PATCH provides a proxy PATCH request allowing authorized clients to make PATHC requests of the microservices
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will call the service and return the result to the client.
func PATCH(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.PATCH(w, url, format, token, r)
}

// POST provides a proxy POST request allowing authorized clients to make POST requests of the microservices
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will call the service and return the result to the client.
func POST(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.POST(w, url, format, token, r)
}

// PUT provides a proxy PUT request allowing authorized clients to make PUT requests of the microservices
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will call the service and return the result to the client.
func PUT(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.PUT(w, url, format, token, r)
}

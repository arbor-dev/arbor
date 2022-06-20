package arbor

import (
	"net/http"

	"github.com/arbor-dev/arbor/proxy"
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
	proxy.DELETE(w, r, url, format, token)
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
	proxy.GET(w, r, url, format, token)
}

// PATCH provides a proxy PATCH request allowing authorized clients to make PATCH requests of the microservices
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
	proxy.PATCH(w, r, url, format, token)
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
	proxy.POST(w, r, url, format, token)
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
	proxy.PUT(w, r, url, format, token)
}

// Establishes a websocket proxy between a client and a microservice
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will first attempt to establish a websocket connection on microservice, then upgrade the client connection.
// Returns once either websocket connection closes.
func ProxyWebsocket(w http.ResponseWriter, url string, format string, token string, r *http.Request) {
	proxy.ProxyWebsocket(w, r, url, format, token)
}

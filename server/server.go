package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/arbor-dev/arbor/logger"
	"github.com/arbor-dev/arbor/security"
	"github.com/arbor-dev/arbor/services"
	"github.com/gorilla/mux"
)

//ArborServer is a struct that manages the proxy server
type ArborServer struct {
	addr   string
	router *mux.Router
	server *http.Server
}

//NewServer creates a new Arbor Server
func NewArborServer(routes services.RouteCollection, addr string, port uint16) *ArborServer {
	a := new(ArborServer)
	a.addr = fmt.Sprintf("%s:%d", addr, port)
	a.router = NewRouter(routes)
	a.server = &http.Server{Addr: a.addr, Handler: a.router}
	return a
}

//StartServer starts the http server in a goroutine to start listening
func (a *ArborServer) StartServer() {
	logger.Log(logger.SPEC, "Roots being planted [Server is listening on "+a.addr+"]")

	err := a.server.ListenAndServe()
	if err != nil {
		if err.Error() == "http: Server closed" {
			return
		}
		logger.Log(logger.FATAL, err.Error())
	}
}

//KillServer ends the http server
func (a *ArborServer) KillServer() {
	logger.Log(logger.SPEC, "Pulling up the roots [Shutting down the server...]")
	a.server.Shutdown(context.Background())
	if security.IsEnabled() {
		security.Shutdown()
	}
}

// StartSecuredServer starts a secured arbor server (Token required for access)
//
// Provide a set of routes to serve and a port to serve on.
func StartSecuredServer(routes services.RouteCollection, addr string, port uint16) *ArborServer {
	srv := NewArborServer(routes, addr, port)
	security.Init()
	srv.StartServer()
	return srv
}

// StartUnsecuredServer starts an unsecured arbor server (Token required for access)
//
// Provide a set of routes to server and a port to serve on/
func StartUnsecuredServer(routes services.RouteCollection, addr string, port uint16) *ArborServer {
	srv := NewArborServer(routes, addr, port)
	srv.StartServer()
	return srv
}

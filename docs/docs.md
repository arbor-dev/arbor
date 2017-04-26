# arbor
--
    import "github.com/acm-uiuc/arbor"


## Usage

#### func  Boot

```go
func Boot(routes RouteCollection, port uint16)
```
Boot is a standard server CLI

Provide a set of routes to serve and a port to serve on.

Usage: executable [-r | --register-client client_name] [-c |
--check-registration token] [-u | --unsecured]

    -r | --register-client client_name

registers a client, generates a token

    -c | --check-registration token

checks if a token is valid and returns name of client

    -u | --unsecured

runs groot without the security layer

    without args

runs groot with the security layer

It will start the arbor instance, parsing the command arguments and execute the
behavior.

#### func  CheckRegistration

```go
func CheckRegistration(token string)
```
CheckRegistration allows you to check what client was assigned to a particular
token

#### func  DELETE

```go
func DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
DELETE provides a proxy DELETE request allowing authorized clients to make
DELETE requests of the microservices

Pass the http Request from the client and the ResponseWriter it expects.

Pass the target url of the backend service (not the url the client called).

Pass the format of the service.

Pass a authorization token (optional).

Will call the service and return the result to the client.

#### func  GET

```go
func GET(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
GET provides a proxy GET request allowing authorized clients to make GET
requests of the microservices

Pass the http Request from the client and the ResponseWriter it expects.

Pass the target url of the backend service (not the url the client called).

Pass the format of the service.

Pass a authorization token (optional).

Will call the service and return the result to the client.

#### func  PATCH

```go
func PATCH(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
PATCH provides a proxy PATCH request allowing authorized clients to make PATHC
requests of the microservices

Pass the http Request from the client and the ResponseWriter it expects.

Pass the target url of the backend service (not the url the client called).

Pass the format of the service.

Pass a authorization token (optional).

Will call the service and return the result to the client.

#### func  POST

```go
func POST(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
POST provides a proxy POST request allowing authorized clients to make POST
requests of the microservices

Pass the http Request from the client and the ResponseWriter it expects.

Pass the target url of the backend service (not the url the client called).

Pass the format of the service.

Pass a authorization token (optional).

Will call the service and return the result to the client.

#### func  PUT

```go
func PUT(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
PUT provides a proxy PUT request allowing authorized clients to make PUT
requests of the microservices

Pass the http Request from the client and the ResponseWriter it expects.

Pass the target url of the backend service (not the url the client called).

Pass the format of the service.

Pass a authorization token (optional).

Will call the service and return the result to the client.

#### func  RegisterClient

```go
func RegisterClient(name string)
```
RegisterClient will generate a access token for a client

Currently uses a db of client names.

#### func  StartServer

```go
func StartServer(routes RouteCollection, port uint16)
```
StartServer starts a secured arbor server (Token required for access)

Provide a set of routes to serve and a port to serve on.

#### func  StartUnsecuredServer

```go
func StartUnsecuredServer(routes RouteCollection, port uint16)
```
StartUnsecuredServer starts an unsecured arbor server (Token required for
access)

Provide a set of routes to server and a port to serve on/

#### type Route

```go
type Route struct {
	Name        string           `json:"Name"`
	Method      string           `json:"Method"`
	Pattern     string           `json:"Pattern"`
	HandlerFunc http.HandlerFunc `json:"Handler"`
}
```

Route is a struct that defines a route for a microservice

Name: Name of the route.

Method: The type of request (GET, POST, DELETE, etc.).

Pattern: The exposed url pattern for clients to hit, allows for url encoded
variables to be specified with {VARIABLE}.

HandlerFunc: The function to handle the request, this basicically should just be
the proxy call, but it allows you to specify more specific things.

#### type RouteCollection

```go
type RouteCollection []Route
```

RouteCollection is a slice of routes that is used to represent a service (may
change name here)

Usage: The recomendation is to create a RouteCollection variable for all of you
services and for each service create a specific one then in a registration
function append all the service collections into the single master collection.

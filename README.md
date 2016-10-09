# GROOT

Groot is the next generation web application serving the UIUC Chapter of ACM.
It is the replacement for liquid which goes defunct 1/1/2016.
The groot repo itself is an API Gateway written in Fall of 2015
It provides the following capabilities:
  * Easy registration of services
  * Universal Authentication for the entire application
  * Proxying API calls
  * Managing inter-service communication

Groot provides a JSON face to any service. When registering as service specify the data encoding and when requesting a resource though groot make the request using json.




## RUNNING GROOT AS OF 10/28/15

Add the API spec in a new file (ex. todo.go) in the services package

There is a set of proxy api calls defined in the proxy package that will route call to the backend services

AS OF 10/28/15
```go
/**
 *  Pass the http Request from the client and the ResponseWriter it expects
 *  Pass the target url of the backend service (not the url the client called)
 *  Will call the service and return the result to the client.
 **/
 func GET(w http.ResponseWriter, url string, format string, r *http.Request)
```
```go
 /**
  *  Pass the http Request from the client and the ResponseWriter it expects
  *  Pass the target url of the backend service (not the url the client called)
  *  Passes the encoded json(only format currently supported) to the service.
  *  Will call the service and return the result to the client.
  **/
  func POST(w http.ResponseWriter, url string, format string, r *http.Request)
```
```go
 /**
  *  Pass the http Request from the client and the ResponseWriter it expects
  *  Pass the target url of the backend service (not the url the client called)
  *  Passes the encoded json(only format currently supported) to the service.
  *  Will call the service and return the result to the client.
  **/
  func PUT(w http.ResponseWriter, url string, format string, r *http.Request)
```
```go
/**
 *  Pass the http Request from the client and the ResponseWriter it expects
 *  Pass the target url of the backend service (not the url the client called)
 *  Will call the service and return the result to the client.
 **/
 func DELETE(w http.ResponseWriter, url string, format string, r *http.Request)
```

All secret data should be kept in a file called secrets.go in the secrets directory 

install the proxy and services packages

```sh
go install github.com/acm-uiuc/groot/proxy

go install github.com/acm-uiuc/groot/secrets

go install github.com/acm-uiuc/groot/services
```

run the server
(make sure you're on IllinoisNet or on the university VPN before running)

```sh
go run ./server/*.go
```

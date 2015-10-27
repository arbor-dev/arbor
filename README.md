# GROOT

Groot is the next generation web application serving the UIUC Chapter of ACM.
It is the replacement for liquid which goes defunct 1/1/2016.
The groot repo itself is an API Gateway written in Fall of 2015
It provides the following capabilities:
	- Easy registration of services
	- Universal Authentication for the entire application
	- Proxying API calls
	- Managing inter-service communication



## RUNNING GROOT AS OF 10/27/15

Add the API spec in a new file (ex. todo.go) in the services package

There is a set of proxy api calls defined in the proxy package that will route call to the backend services

AS OF 10/27/15
```go
/**
 *  Pass the http Request from the client and the ResponseWriter it expects
 *  Pass the target url of the backend service (not the url the client called)
 *  Will call the service and return the result to the client.
 **/
 func GETHandler(w http.ResponseWriter, url string, r *http.Request)
```

install the proxy and services packages

```sh
cd proxy
go install

cd ../services
go install
```

run the server

```sh
go run ./server/*.go
```

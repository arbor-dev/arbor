
![alt tag](https://github.com/acm-uiuc/arbor/blob/master/docs/readme.png)

Arbor is a statically configured framework for creating API Gateways. 
It provides the following capabilities:
  * Easy registration of services
  * Basic application authorization and request sanitization
  * Proxying API calls
  * Managing inter-service communication

Arbor currently provides a JSON face to any service. When registering as service specify the data encoding and when requesting a resource though groot make the request using json.




## Registering Services

Add the API spec in a new file (ex. todo.go) in the services package

There is a set of proxy api calls defined in the proxy package that will route call to the backend services

AS OF 10/28/15
```go
/**
 *  Pass the http Request from the client and the ResponseWriter it expects
 *  Pass the target url of the backend service (not the url the client called)
 *  Pass the format of the service
 *  Pass a authorization token (optional)
 *  Will call the service and return the result to the client.
 **/
 func GET(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
```go
 /**
  *  Pass the http Request from the client and the ResponseWriter it expects
  *  Pass the target url of the backend service (not the url the client called)
  *  Passes the encoded json(only format currently supported) to the service.
  *  Pass a authorization token (optional)
  *  Will call the service and return the result to the client.
  **/
  func POST(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
```go
 /**
  *  Pass the http Request from the client and the ResponseWriter it expects
  *  Pass the target url of the backend service (not the url the client called)
  *  Passes the encoded json(only format currently supported) to the service.
  *  Pass a authorization token (optional)
  *  Will call the service and return the result to the client.
  **/
  func PUT(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```
```go
/**
 *  Pass the http Request from the client and the ResponseWriter it expects
 *  Pass the target url of the backend service (not the url the client called)
 *  Pass a authorization token (optional)
 *  Will call the service and return the result to the client.
 **/
 func DELETE(w http.ResponseWriter, url string, format string, token string, r *http.Request)
```

All secret data should be kept in a file called config.go in the config directory

Install Dependencies [First time setup]

```sh
go get github.com/gorilla/mux

go get github.com/boltdb/bolt

go get github.com/kennygrant/sanitize

go install github.com/gorilla/mux

go install github.com/boltdb/bolt

go install github.com/kennygrant/sanitize
```

install packages

```sh
go install github.com/acm-uiuc/arbor/proxy

go install github.com/acm-uiuc/arbor/services

go install github.com/acm-uiuc/arbor/security

go install github.com/acm-uiuc/arbor/server
```

run the server

```sh
go run ./server/*.go
```
compile the service 

```sh
go build -o groot [PATH TO GROOT]/server
```

## CLI 
```sh
groot [-r | --register-client client_name] [-c | --check-registration token] [-u | --unsecured]
```

-r | --register-client *client_name*
> registers a client, generates a token

-c | --check-registration *token*
> checks if a token is valid and returns name of client

-u | --unsecured
> runs groot without the security layer 

*without args* 
> runs groot with the security layer


## License

This project is licensed under the University of Illinois/NCSA Open Source License. For a full copy of this license take a look at the LICENSE file. 

When contributing new files to this project, preappend the following header to the file as a comment: 

```
Copyright © 2017, ACM@UIUC

This file is part of the Groot Project.  
 
The Groot Project is open source software, released under the University of Illinois/NCSA Open Source License. 
You should have received a copy of this license in a file with the distribution.
```
